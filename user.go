package users

import "gopkg.in/dealancer/validate.v2"

// User struct
type User struct {
	Username string `validate:"gte=3 & lte=25 & format=alnum_unicode"`
	Email    string `validate:"empty=true | format=email"`
	Password string
	IsActive *bool `validate:"nil=true"`
}

// CreateUser does ...
func CreateUser(username string, email string, pass string) (User, error) {
	var err error

	user := User{
		Username: username,
		Email:    email,
		Password: pass,
	}

	if err = validate.Validate(&user); err != nil {
		return user, err
	}
	return user, nil
}
