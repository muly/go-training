package user

import (
	"testing"

	gomock "github.com/golang/mock"

	"github.com/muly/go-training/20200504/gomock/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}


/*
mulys-mbp:user s$ go test
# github.com/muly/go-training/20200504/gomock/user
package github.com/muly/go-training/20200504/gomock/user (test)
        imports github.com/golang/mock: build constraints exclude all Go files in /Users/s/go/src/github.com/golang/mock
FAIL    github.com/muly/go-training/20200504/gomock/user [setup failed]
*/