package models

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
    UsernameOrEmail string `json:"username_or_email" binding:"required"`
    Password        string `json:"password" binding:"required"`
}