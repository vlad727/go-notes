package main

import (
	"fmt"
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}
// global var 
var (
	invalidRequest = "invalid request"
)

func Validate(req UserCreateRequest) string {

	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest + "Incorrect FirstName"
	} else if req.Age <= 0 || req.Age > 150 {
		return invalidRequest + "Incorrect Age"
	} else {
		return "Valid string"
	}
}
func main() {
	responseFromValidate := Validate(UserCreateRequest{"John", 25})
	fmt.Println(responseFromValidate)
}
