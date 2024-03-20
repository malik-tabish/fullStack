package services

import (
	"fmt"
	"gin-example/models"
	"gin-example/repositories"
)

func CreateUserService(user *models.User) error {
	if !repositories.CheckUsernameExists(user.UserName) {
		return fmt.Errorf("username already exists")
	}
	return repositories.CreateUserRepository(user)
}

func GetAllUsersService() (*[]models.User, error) {
	if repositories.CheckUsers() {
		return nil, fmt.Errorf("users table is empty")
	}
	return repositories.GetAllUsersRepository()
}

func GetUserByIDService(userID string) (*models.User, error) {
	user, err := repositories.GetUserByIDRepository(userID)

	if err != nil {
		return nil, fmt.Errorf("error while getting user by id")
	}
	return user, nil
}

func DeleteUserByIDService(userID string) error {
	err := repositories.DeleteUserByIDRepository(userID)

	if err != nil {
		return fmt.Errorf("error while deleting user by ID")
	}
	return nil
}

func UpdateUserService(userID string, updatedUser *models.User) error {
	err := repositories.UpdateUserRepository(userID, updatedUser)

	if err != nil {
		return fmt.Errorf("error while updating the user")
	}
	return nil
}
