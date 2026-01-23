package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/middleware"
	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/models"
	"github.com/sujeet-crossml/GoLang_Backend_Project/internal/utils"
)

// Register
func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	hashedPwd, _ := utils.HashPassword(u.Password)
	u.Password = hashedPwd

	if err := models.CreateUser(&u); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error creating user(Email might represent)")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user registered successfully"})

}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	token, _ := utils.GenerateToken(user.ID)
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

}

// get profile (protected)
func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)

	user, err := models.GetUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIDKey).(int)

	orders, err := models.GetOrdersByUserID(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error fetching orders")
		return
	}

	if orders == nil {
		orders = []models.Order{}
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}
