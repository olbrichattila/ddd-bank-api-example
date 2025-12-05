package user

type User interface {
	Create(entity UserEntity) (string, error)
	Get(userId string) (UserEntity, error)
	GetByEmail(email string) (UserEntity, error)
	Update(entity UserEntity) error
	Delete(userID string) (int64, error)
}
