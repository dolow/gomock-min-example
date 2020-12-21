package service

import (
	"gomock_example/dao"
)

func Signup(user dao.User) error {
	return user.Create()
}
