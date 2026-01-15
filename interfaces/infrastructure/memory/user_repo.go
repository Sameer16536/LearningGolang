// infrastructure/memory/user_repo.go
package memory

import domain "learninggolang/interfaces/domain"

type MemoryUserRepo struct {
	data map[int]domain.User
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{
		data: make(map[int]domain.User),
	}
}

func (m *MemoryUserRepo) Create(user domain.User) error {
	user.ID = len(m.data) + 1
	m.data[user.ID] = user
	return nil
}

func (m *MemoryUserRepo) GetByID(id int) (domain.User, error) {
	return m.data[id], nil
}
