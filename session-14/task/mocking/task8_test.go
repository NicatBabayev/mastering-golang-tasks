package mocking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotifier struct {
	mock.Mock
}

func (m *MockNotifier) NotifyUser(message string) error {
	if message == "success" {
		return nil
	} else {
		return fmt.Errorf("failed to do notify user")
	}
}

func TestNotifyUser(t *testing.T) {
	mockObject := new(MockNotifier)
	mockObject.On("NotifyUser", "success").Return(nil)
	mockObject.On("NotifyUser", "fail").Return(fmt.Errorf("failed to do notify user"))

	res1 := mockObject.NotifyUser("success")
	res2 := mockObject.NotifyUser("fail")

	assert.Equal(t, nil, res1, "This test should return the nil as message is \"success\"")
	assert.Equal(t, fmt.Errorf("failed to do notify user"), res2, "This test should return error as message is not \"sucess\"")
}
