package types

type SignupCredentials struct {
  Fullname string `json:"fullname" binding:"required"`
  Email    string `json:"email" binding:"required"` 
  Password string `json:"password" binding:"required"`
}
