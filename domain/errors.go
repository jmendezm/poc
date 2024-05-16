package domain

type CustomError struct {
	Code    int
	Err     string
	Message string
}

func (ce *CustomError) Error() string {
	return ce.Message
}

var ErrUserNotFound = CustomError{
	Code:    404,
	Err:     "ERR_USER_NOT_FOUND",
	Message: "User not found",
}

var ErrConnectionNotFound = CustomError{
	Code:    404,
	Err:     "ERR_CONNECTION_NOT_FOUND",
	Message: "Connection not found",
}

var ErrInternalServerError = CustomError{
	Code:    500,
	Err:     "ERR_INTERNAL_SERVER_ERROR",
	Message: "Internal server error",
}

var ErrWrongCredentials = CustomError{
	Code:    403,
	Err:     "ERR_WRONG_CREDENTIALS",
	Message: "Not valid email or password",
}

var ErrAccountUnAvailable = CustomError{
	Code:    403,
	Err:     "ERR_ACCOUNT_UNAVAILABLE",
	Message: "Your account is deleted/inactive",
}

var ErrNotLoggedIn = CustomError{
	Code:    403,
	Err:     "ERR_NOT_LOGGED_IN",
	Message: "You must login",
}

var ErrSiteNotFound = CustomError{
	Code:    404,
	Err:     "ERR_SITE_NOT_FOUND",
	Message: "Site does not exists",
}
