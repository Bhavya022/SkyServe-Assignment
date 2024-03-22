//models: This folder contains data models or structs used in your application.

//user.go: Defines the structure of a user object.

// user.go

package models

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // NOTE: In practice, use a more secure way to handle passwords
}
