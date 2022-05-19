package user

import "time"

type User struct {
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"date_created"`
	Generated int       `json:"number_of_generations"`
	password  string
}
