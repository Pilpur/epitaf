package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/aureleoules/epitaf/lib/cri"
	"github.com/aureleoules/epitaf/models"
	"github.com/gin-gonic/gin"
)

const (
	// Endpoint api url
	Endpoint = "https://login.microsoftonline.com/3534b3d7-316c-4bc9-9ede-605c860f49d2/oauth2/v2.0"
)

func handleAuth() {
	users := api.Group("/users")

	users.POST("/authenticate", authenticateHandler)
	users.POST("/callback", auth.LoginHandler)
}

func authMiddleware() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "epitaf",
		Key:        []byte(os.Getenv("JWT_SECRET")),
		Timeout:    time.Hour * 48,
		MaxRefresh: time.Hour * 48,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			u := data.(*models.User)
			return jwt.MapClaims{
				"uuid":      u.UUID.String(),
				"email":     u.Email,
				"name":      u.Name,
				"promotion": u.Promotion,
				"class":     u.Class,
				"region":    u.Region,
				"semester":  u.Semester,
			}
		},
		Authenticator: callbackHandler,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.AbortWithError(code, errors.New(message))
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		panic(err)
	}

	return authMiddleware
}

func authenticateHandler(c *gin.Context) {
	req, _ := http.NewRequest("GET", Endpoint+"/authorize", nil)
	q := req.URL.Query()

	q.Add("client_id", os.Getenv("CLIENT_ID"))
	q.Add("response_type", "code")
	q.Add("response_mode", "query")
	q.Add("state", "0000")
	q.Add("scope", "https://graph.microsoft.com/User.Read")

	if os.Getenv("DEV") == "true" {
		q.Add("redirect_uri", "http://localhost:3000/callback")
	} else {
		q.Add("redirect_uri", "https://epitaf.aureleoules.com/callback")
	}

	req.URL.RawQuery = q.Encode()
	c.JSON(http.StatusOK, req.URL.String())
}

func callbackHandler(c *gin.Context) (interface{}, error) {
	var m map[string]string
	c.Bind(&m)
	var uri string
	if os.Getenv("DEV") == "true" {
		uri = "http://localhost:3000/callback"
	} else {
		uri = "https://epitaf.aureleoules.com/callback"
	}

	resp, err := http.PostForm(Endpoint+"/token", url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {m["code"]},
		"client_id":     {os.Getenv("CLIENT_ID")},
		"client_secret": {os.Getenv("CLIENT_SECRET")},
		"redirect_uri":  {uri},
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, jwt.ErrFailedAuthentication
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	var result map[string]string
	json.Unmarshal([]byte(body), &result)

	token := result["access_token"]
	if token == "" {
		return nil, jwt.ErrFailedAuthentication
	}

	profile, err := getProfile(token)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	u, err := models.GetUserByEmail(profile.Mail)
	if err != nil {
		user := models.User{
			Name:  profile.DisplayName,
			Email: profile.Mail,
		}

		// CRI req
		client := cri.NewClient(os.Getenv("CRI_USERNAME"), os.Getenv("CRI_PASSWORD"), nil)
		r, err := client.SearchUser(user.Email)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		var slug string
		for _, g := range r.GroupsHistory {
			if g.IsCurrent {
				slug = g.Group.Slug
				user.Promotion = g.GraduationYear
				break
			}
		}

		group, err := client.GetGroup(slug)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		g := strings.Split(group.Name, " ")
		user.Semester = g[0]
		user.Region = g[1]
		user.Class = g[2]

		err = user.Insert()
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return &user, nil
	}

	return u, nil
}

func getProfile(token string) (models.MicrosoftProfile, error) {
	endpoint := "https://graph.microsoft.com/v1.0/me"
	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.MicrosoftProfile{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.MicrosoftProfile{}, err
	}

	var result models.MicrosoftProfile
	fmt.Println(string(body))
	json.Unmarshal([]byte(body), &result)

	if result.Mail == "" {
		return models.MicrosoftProfile{}, errors.New("invalid token")
	}

	return result, nil
}