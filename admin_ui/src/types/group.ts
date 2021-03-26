export type Group = {
	id?: string,
	realm_id?: string,
	usable?: boolean,
	name?: string,
	slug?: string,
	parent_id?: string,
	archived?: boolean,
	archived_at?: Date,
	active_at?: Date,
	subgroups?: Array<Group>,
	created_at?: Date,
	updated_at?: Date,
};