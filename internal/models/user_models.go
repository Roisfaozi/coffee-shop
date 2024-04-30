package models

import "time"

type User struct {
	ID        string    `db:"id" form:"id" json:"id,omitempty" valid:"-"`
	Username  string    `db:"username" form:"username" json:"username" valid:"type(string),required"`
	Email     string    `db:"email" form:"email" json:"email,omitempty" valid:"email,required"`
	Password  string    `db:"password" form:"password" json:"password,omitempty" valid:"stringlength(6|100)~Password minimal 6,required"`
	Role      string    `db:"role" form:"role" json:"role,omitempty" valid:"-"`
	CreatedAt time.Time `db:"created_at" form:"created_at" json:"created_at" valid:"-"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at" json:"updated_at" valid:"-"`
}

type UserResponse struct {
	ID       string `db:"id" form:"id" json:"id"`
	Username string `db:"username" form:"username" json:"username"`
	Email    string `db:"email" form:"email" json:"email"`
}
type Users []User
