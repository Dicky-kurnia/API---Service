package listerr

import "fmt"

var (
	BAD_REQUEST            = "BAD_REQUEST"
	NOT_BLANK              = "NOT_BLANK"
	NOT_VALID              = "NOT_VALID"
	NOT_MATCH              = "NOT_MATCH"
	NOT_FOUND              = "NOT_FOUND"
	AUTHENTICATION_FAILURE = "AUTHENTICATION_FAILURE"
	INTERNAL_ERROR         = "INTERNAL_ERROR"
	UNPROCESSABLE_ENTITY   = "UNPROCESSABLE_ENTITY"
)

func Min(x interface{}) string {
	return fmt.Sprintf("MIN_%v", x)
}

func Max(x interface{}) string {
	return fmt.Sprintf("MAX_%v", x)
}

type CustomErr struct {
	status  int
	code    string
	message string
}

func (c *CustomErr) Status() int {
	return c.status
}

func (c *CustomErr) Code() string {
	return c.code
}

func (c *CustomErr) Error() string {
	return c.message
}

func NewError(status int, code, message string) error {
	return &CustomErr{
		status:  status,
		code:    code,
		message: message,
	}
}
