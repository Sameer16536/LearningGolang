package domain

// UserRepository defines behavior required by the business layer.
// Any database that can store users must implement this.
type UserRepository interface {
	Create(user User) error
	GetByID(id int) (User, error)
}
