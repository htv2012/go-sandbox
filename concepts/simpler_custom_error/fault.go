package main

type CustomError struct {
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}

// Sentinel errors
var (
	ErrEmptyName     = CustomError{"name must be non-empty"}
	ErrEmptyPassword = CustomError{"password must be non-empty"}
	ErrLoginFailed   = CustomError{"login failed"}
)

func login(user, password string) error {
	if user == "foo" && password == "bar" {
		return nil
	} else if user == "" {
		return ErrEmptyName
	} else if password == "" {
		return ErrEmptyPassword
	}
	return ErrLoginFailed
}
