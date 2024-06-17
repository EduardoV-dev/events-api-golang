package auth

// Credentials used for login method 
type loginCredentials struct {
  Email    string `json:"email" binding:"required"` 
  Password string `json:"password" binding:"required"`
}

