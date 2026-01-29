package models

// UserProfile - DTO для отдачи профиля пользователя
// генерирует токены
type UserProfile struct {
	ID       string   `json:"id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
