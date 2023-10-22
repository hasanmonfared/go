package storage

import "interface_test/user"

type Memory struct {
	users []user.User
}

func (m *Memory) CreateUser(u user.User) {
	m.users = append(m.users, u)
}
func (m *Memory) ListUsers() []user.User {
	return m.users
}
func (m *Memory) GetUserByID(id uint) user.User {
	for _, u := range m.users {
		if id == u.ID {
			return u
		}
	}
	return user.User{}
}
