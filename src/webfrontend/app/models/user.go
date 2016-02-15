package models

import (
	"github.com/mrjones/oauth"
	"math/rand"
)

type User struct {
	Username     string
	Uid          int
	RequestToken *oauth.RequestToken
	AccessToken  *oauth.AccessToken
}

func FindOrCreate(username string) *User {
	if user, ok := db[username]; ok {
		return user
	}
	user := &User{Username: username}
	db[username] = user
	return user
}

func GetUser(username string) *User {
	return db[username]
}

func NewUser(username string) *User {
	user := &User{Username: username, Uid: rand.Intn(100000)}
	db[username] = user
	return user
}

var db = make(map[string]*User)
