package service

import (
	"math/rand"
	"time"
)

type Role int

const (
	OPERATOR Role = 0x1
	ADMIN    Role = 0x1 << 1
)

var usersDb = map[string]userInfo{
	"john": {
		username: "john",
		password: "123456",
		role:     OPERATOR,
		id:       1234567890,
		token:    "JG3LDSkEzbQgnGcIU7o1P8p2FxuHUMg8", //generate by RandStringBytes
	},
	"alice": {
		username: "alice",
		password: "123456",
		role:     ADMIN,
		id:       9987654321,
		token:    "Tex1h8VF75cB3Y6WTRPVF6hkUUFrK9lj",
	},
}

type LoginService interface {
	LoginUser(username string, password string) (Role, int64, bool)
}

type server struct {
	db map[string]userInfo
}

type userInfo struct {
	username string
	password string
	role     Role
	id       int64
	token    string
}

func StaticLoginService() LoginService {
	return &server{db: usersDb}
}

func (s *server) LoginUser(username string, password string) (Role, int64, bool) {
	if user, ok := s.db[username]; ok {
		if user.password == password {
			return user.role, user.id, true
		}
	}
	return -1, -1, false
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
