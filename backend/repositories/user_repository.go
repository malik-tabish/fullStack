package repositories

import (
	"fmt"
	"gin-example/db"
	"gin-example/models"
)

func CreateUserRepository(user *models.User) error {
	_, err := db.Conn.Exec("INSERT INTO users (name, username, password) VALUES (?,?,?)", user.Name, user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("error while inserting into the table")
	}
	return nil
}

func CheckUsernameExists(username string) bool {
	var count int
	err := db.Conn.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}

func GetAllUsersRepository() (*[]models.User, error) {
	var users []models.User
	rows, err := db.Conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error occured while executing the query")
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.UserName, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("error while adding data into array")
		}
		users = append(users, user)
	}
	return &users, nil
}

func CheckUsers() bool {
	var count int
	err := db.Conn.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}

func GetUserByIDRepository(userID string) (*models.User, error) {
	var user models.User
	err := db.Conn.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&user.Id, &user.Name, &user.UserName, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("error while getting user by ID")
	}
	return &user, nil
}

func DeleteUserByIDRepository(userID string) error {
	_, err := db.Conn.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return fmt.Errorf("error while deleting user by ID")
	}
	return nil
}

func UpdateUserRepository(userID string, updatedUser *models.User) error {
	_, err := db.Conn.Exec("UPDATE users SET name = ?, username = ?, password = ? WHERE id = ?", updatedUser.Name, updatedUser.UserName, updatedUser.Password, userID)
	if err != nil {
		return fmt.Errorf("error while updating the user")
	}
	return nil
}
