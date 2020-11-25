// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Aurèle Oulès",
            "url": "https://www.aureleoules.com",
            "email": "contact@epitaf.fr"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/classes": {
            "get": {
                "description": "Get list of all registered classes",
                "tags": [
                    "classes"
                ],
                "summary": "Get classes",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Get tasks",
                "tags": [
                    "tasks"
                ],
                "summary": "Get tasks",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error\" \"Server error"
                    }
                }
            },
            "post": {
                "description": "Create a new task",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create task",
                "parameters": [
                    {
                        "description": "Task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        },
        "/tasks/{short_id}": {
            "get": {
                "description": "Get a specific task",
                "tags": [
                    "tasks"
                ],
                "summary": "Get task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short_id",
                        "name": "short_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error\" \"Server error"
                    }
                }
            },
            "put": {
                "description": "Edit a specific task",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short_id",
                        "name": "short_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            },
            "delete": {
                "description": "Delete a specific task",
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short_id",
                        "name": "short_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        },
        "/tasks/{short_id}/complete": {
            "post": {
                "description": "Mark a specific task as completed",
                "tags": [
                    "tasks"
                ],
                "summary": "Complete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short_id",
                        "name": "short_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            },
            "delete": {
                "description": "Mark a specific task as uncompleted",
                "tags": [
                    "tasks"
                ],
                "summary": "Uncomplete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short_id",
                        "name": "short_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        },
        "/users/authenticate": {
            "post": {
                "description": "Build Microsoft oauth url",
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate URL",
                "parameters": [
                    {
                        "default": "https://www.epitaf.fr/callback",
                        "description": "redirect_uri",
                        "name": "redirect_uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "406": {
                        "description": "Not acceptable"
                    }
                }
            }
        },
        "/users/calendar": {
            "get": {
                "description": "Get user calendar",
                "tags": [
                    "users"
                ],
                "summary": "Get calendar",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error\" \"Server error"
                    }
                }
            }
        },
        "/users/callback": {
            "post": {
                "description": "Authenticate user and return JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "OAuth Callback",
                "parameters": [
                    {
                        "description": "code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "redirect_uri",
                        "name": "redirect_uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "description": "Retrieve data about current user",
                "tags": [
                    "users"
                ],
                "summary": "Get self",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not found"
                    },
                    "406": {
                        "description": "Not acceptable"
                    },
                    "500": {
                        "description": "Server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Task": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Class",
                    "type": "string"
                },
                "completed": {
                    "type": "boolean"
                },
                "completed_at": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_by_login": {
                    "description": "Meta",
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "members": {
                    "description": "Students",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "promotion": {
                    "description": "Promotion",
                    "type": "integer"
                },
                "region": {
                    "type": "string"
                },
                "semester": {
                    "type": "string"
                },
                "short_id": {
                    "description": "Meta",
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "title": {
                    "description": "Body",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                },
                "updated_by_login": {
                    "type": "string"
                },
                "visibility": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "https://api.epitaf.fr",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Epitaf API v1 Docs",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
