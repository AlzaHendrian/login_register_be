package authdto

type AuthRequest struct {
	FirstName  string `json:"first_name" form:"first_name" validate:"required"`
	LastName   string `json:"last_name" form:"last_name" validate:"required"`
	DateOfBird string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Gender     string `json:"gender" form:"gender" validate:"required"`
	Email      string `json:"email" form:"email" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	Role       string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
