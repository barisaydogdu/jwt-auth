package domain

import "time"

type User struct {
	ID        uint      `db:"id" json:"id"`
	FirstName string    `db:"firstname" json:"firstname"`
	LastName  string    `db:"lastname" json:"lastname"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Role      string    `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
