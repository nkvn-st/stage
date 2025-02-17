package userservice

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var Users []User
	err := r.db.Find(&Users).Error
	return Users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	err := r.db.Model(&User{}).Where("id = ?", id).Update("email", user.Email).Error
	if err != nil {
		return User{}, err
	}

	var result User
	err = r.db.First(&result, id).Error
	return result, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	err := r.db.Unscoped().Delete(&User{}, id).Error
	return err
}
