package authsystem

import (
	"fmt"
)

type UserNotFoundError struct {
	msg string
}

func (unfErr UserNotFoundError) Error() string {
	return unfErr.msg
}

type InvalidPasswordError struct {
	msg string
}

func (ipErr InvalidPasswordError) Error() string {
	return ipErr.msg
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func authenticate(username, password string) error {
	userNotFoundErr := UserNotFoundError{msg: username}
	invalidPassErr := InvalidPasswordError{msg: username}
	if passValue, ok := users[username]; ok {
		if passValue == password {
			return nil
		} else {
			return fmt.Errorf("invalid password for user: \"%w\"", invalidPassErr)
		}
	} else {
		return fmt.Errorf("user not found: \"%w\"", userNotFoundErr)
	}
}

func RunTask5() {
	err := authenticate("user1", "password1")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %w\n", err))
	} else {
		fmt.Println("Login successful!")
	}

}
