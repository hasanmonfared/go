package new_in_memory

import "interface_test/user"

type Store struct {
	users map[uint]user.User
}

func (s *Store) GetUserByID(id uint) user.User {

	return s.users[id]
}
