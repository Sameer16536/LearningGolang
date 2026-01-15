// infrastructure/postgres/user_repo.go
package postgres

import domain "learninggolang/interfaces/domain"

// PostgresUserRepo is a concrete DB implementation
type PostgresUserRepo struct {
	// db *sql.DB (real world)
}

// Ensure compile-time interface compliance
var _ domain.UserRepository = (*PostgresUserRepo)(nil)

func NewPostgresUserRepo() *PostgresUserRepo {
	return &PostgresUserRepo{}
}

func (r *PostgresUserRepo) Create(user domain.User) error {
	// INSERT INTO users ...
	return nil
}

func (r *PostgresUserRepo) GetByID(id int) (domain.User, error) {
	return domain.User{ID: id, Name: "Sameer"}, nil
}
