package models

type User struct {
	ID       int
	Name     string
	Login    string `json:"login"`
	Password string `json:"password"`
	Roles    []*Role
	Groups   []*Group
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
