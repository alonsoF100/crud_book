package handlers

import (
	"crud_book/internal/dto"
	"crud_book/internal/services"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func newUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) createUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserToResponse(user)
	c.JSON(201, response)
}

func (h *userHandler) getUser(c *gin.Context) {
	var req dto.GetUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	user, err := h.userService.GetUser(req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserToResponse(user)
	c.JSON(200, response)
}

func (h *userHandler) deleteUser(c *gin.Context) {
	var req dto.DeleteUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err := h.userService.DeleteUser(req)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func (h *userHandler) getAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get users"})
		return
	}

	response := dto.UsersToResponse(users)
	c.JSON(200, response)
}
