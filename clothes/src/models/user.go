package models

import (
	"sort"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Name       string `orm:"size(64)"`
	Department string `orm:size(64)`
	Office     string `orm:size(64)`
	Size       int
}

type Users []*User

func (s Users) Len() int { return len(s) }

func (s Users) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Users) Less(i, j int) bool { return s[i].Id < s[j].Id }

func (u *User) Insert() (err error) {
	_, err = orm.NewOrm().Insert(u)
	return
}

func (u *User) Update(fields ...string) (err error) {
	_, err = orm.NewOrm().Update(u, fields...)
	return
}

func (u *User) Read(fields ...string) (err error) {
	err = orm.NewOrm().Read(u, fields...)
	return
}

func ReadUsers() ([]*User, error) {
	var users Users
	qs := orm.NewOrm().QueryTable("user")
	_, err := qs.All(&users)

	sort.Sort(users)

	return users, err
}

func init() {
	orm.RegisterModel(new(User))
}
