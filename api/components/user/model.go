package user

import (
	"account_system/api/common"
)

type User struct {
	Id   common.IdType `json:"id"`
	Name string `json:"name"`
}

func New(name string) *User {
	u := &User{}
	u.Id = common.NewId()
	u.Name = name
	return u
}
