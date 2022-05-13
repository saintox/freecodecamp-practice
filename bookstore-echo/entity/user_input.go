package entity

type (
	UserInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}

	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
