package controllers

import (
	"gin-example/models"
	"gin-example/services"
	"gin-example/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CREATE NEW USER....
func CreateUserController(context *gin.Context) {
	context.Header("Content-Type", "json")
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "Invalid json data"))
		return
	}
	// CALLING SERVICE LAYER....
	err = services.CreateUserService(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "User creation failed"))
		return
	}
	context.JSON(http.StatusCreated, utils.NewMessage(http.StatusCreated, err, "Created", "User created successfully"))
}

// GET ALL USERS....
func GetAllUsersController(context *gin.Context) {
	context.Header("Content-Type", "Json")
	allUsers, err := services.GetAllUsersService()

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "Error while getting all users"))
		return
	}
	context.JSON(http.StatusOK, allUsers)
}

// GET USER BY ID....
func GetByIdController(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	userID := context.Param("id")

	user, err := services.GetUserByIDService(userID)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "User not found"))
		return
	}
	context.JSON(http.StatusOK, user)
}

// DELETE USER....
func DeleteByIdController(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	userID := context.Param("id")

	//CHECK IF USER EXISTS....
	_, err := services.GetUserByIDService(userID)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.NewMessage(http.StatusNotFound, err, "Not Found", "User not found"))
		return
	}

	//IF USER EXISTS PROCEED WITH DELETION.......
	err = services.DeleteUserByIDService(userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "Error while deleting User by ID"))
		return
	}
	context.JSON(http.StatusOK, utils.NewMessage(http.StatusOK, nil, "OK", "User deleted successfully"))
}

// UPDATE USER....
func UpdateUserController(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	userID := context.Param("id")

	// CHECK IF USER EXISTS.....
	_, err := services.GetUserByIDService(userID)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.NewMessage(http.StatusNotFound, err, "Not Found", "User not found"))
		return
	}

	// IF USER EXISTS PROCEDE WITH UPDATE.....
	var updatedUser models.User
	if err := context.BindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "Invalid JSON data"))
		return
	}

	err = services.UpdateUserService(userID, &updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.NewMessage(http.StatusBadRequest, err, "Bad Request", "Error while updating the user"))
		return
	}
	context.JSON(http.StatusOK, utils.NewMessage(http.StatusOK, nil, "OK", "User updated sucessfully"))
}
