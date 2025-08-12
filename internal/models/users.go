package models

type User struct {
	ID       int      `json:"id,omitempty" swaggerignore:"true"`
	Name     string   `json:"name,omitempty" swaggerignore:"true"`
	Login    string   `json:"login" example:"user"`
	Password string   `json:"password" example:"password"`
	Roles    []*Role  `json:"roles,omitempty" swaggerignore:"true"`
	Groups   []*Group `json:"groups,omitempty" swaggerignore:"true"`
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
