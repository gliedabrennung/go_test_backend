package entity

type RawUser struct {
	Username string `validate:"required,min=4,max=14,alphanum"`
	Password string `validate:"required,min=8,max=16"`
}
