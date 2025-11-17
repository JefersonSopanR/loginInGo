package main

import (
	"fmt"
)

type user struct {
	username string
	password string
	gmail string
	displayName string
}

type db struct {
	users []*user
}

func goDataBase(f form, db *db) *user {
	for _, u := range db.users {
		if f.username == u.username {
			if f.password == u.password {
				return u
			}
		}
	}
	fmt.Println("")
	return nil
}

func (d *db) deleteUser(u *user) {
	for i, uu := range d.users {
		if uu == u {
			d.users = append(d.users[:i], d.users[i+1:]...)
			return
		}
	}
}

