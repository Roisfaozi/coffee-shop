package models

import "time"

type User struct {
	ID        string    `db:"id" form:"id" json:"id"`
	Username  string    `db:"username" form:"username" json:"username"`
	Email     string    `db:"email" form:"email" json:"email"`
	Password  string    `db:"password" form:"password" json:"password"`
	Role      string    `db:"role" form:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}

type UserResponse struct {
	ID       string `db:"id" form:"id" json:"id"`
	Username string `db:"username" form:"username" json:"username"`
	Email    string `db:"email" form:"email" json:"email"`
	Role     string `db:"role" form:"role" json:"role"`
}
