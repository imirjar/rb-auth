package models

type User struct {
	ID     string `json:"id"`
	Name   string
	Roles  []*Role
	Groups []*Group
}

func (u *User) HasRole(rname string) bool {
	for _, i := range u.Roles {
		if i.Name == rname {
			return true
		}
	}
	return false
}

type Group struct {
	ID    string
	Name  string
	Roles []*Role
}

func (g *Group) HasRole(rname string) bool {
	for _, i := range g.Roles {
		if i.Name == rname {
			return true
		}
	}
	return false
}

type Role struct {
	ID   string
	Name string
}
