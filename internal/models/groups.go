package models

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
