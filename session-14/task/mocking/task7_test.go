package mocking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUser(id int) (User, error) {
	correctUser := User{
		ID:   1,
		Name: "Koroglu",
		Age:  7777,
	}
	wrongUser := User{
		ID:   2,
		Name: "Koroglu",
		Age:  7777,
	}

	if id == 1 {
		return correctUser, nil
	} else if id == 2 {
		return wrongUser, fmt.Errorf("User data is wrong")
	} else {
		return User{}, fmt.Errorf("User not found")
	}
}

func TestGetUser(t *testing.T) {
	testUser := User{
		ID:   1,
		Name: "Koroglu",
		Age:  7777,
	}
	mockObject := new(MockDatabase)

	mockObject.On("GetUserByID", 1).Return(testUser, nil)
	mockObject.On("GetUserByID", 2).Return(User{}, fmt.Errorf("User data is wrong"))
	mockObject.On("GetUserByID", 3).Return(User{}, fmt.Errorf("User not found"))

	res1, err1 := mockObject.GetUser(1)
	res2, err2 := mockObject.GetUser(2)
	assert.Equal(t, testUser, res1, "ID 1 should return the testUser")
	assert.Nil(t, err1)
	assert.Equal(t, User{}, res2, "ID 2 should return empty User{}")
	assert.NotNil(t, err2)

}
