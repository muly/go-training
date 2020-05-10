package user

import "github.com/muly/go-training/20200504/gomock/doer"

type User struct {
    Doer doer.Doer
}

func (u *User) Use() error {
    return u.Doer.DoSomething(123, "Hello GoMock")
}

// // https://blog.codecentric.de/en/2017/08/gomock-tutorial/
// go get github.com/golang/mock/gomock
// go get github.com/golang/mock/mockgen
// mockgen -destination=mocks/mock_doer.go -package=mocks github.com/muly/go-training/20200504/gomock/doer Doer
