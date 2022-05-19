package user

import "time"

type User struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"date_created"`
	Generated int       `json:"number_of_generations"`
}
