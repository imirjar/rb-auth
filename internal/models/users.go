package models

type User struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Roles    []*Role  `json:"roles,omitempty"`
	Groups   []*Group `json:"groups,omitempty"`
}

func (u *User) IsValid() bool {
	if u.Login == "" || u.Password == "" {
		return false
	}
	return true
}

func (u *User) HasRole(rname string) bool {
	for _, i := range u.Roles {
		if i.Name == rname {
			return true
		}
	}
	return false
}
