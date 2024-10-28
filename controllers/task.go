package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-manager/config"
	"task-manager/middlewares"
	"task-manager/models"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	task.UserID = userID
	if err := config.DB.Create(&task).Error; err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task created successfully"})
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	var tasks []models.Task
	if err := config.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		http.Error(w, "Invalid Id", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	taskID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		http.Error(w, "Error finding task", http.StatusInternalServerError)
		return
	}

	if updatedTask.Title != nil {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != nil {
		task.Description = updatedTask.Description
	}
	if updatedTask.Status != nil {
		task.Status = updatedTask.Status
	}
	if updatedTask.DueDate != nil {
		task.DueDate = updatedTask.DueDate
	}

	if err := config.DB.Save(&task).Error; err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(uint)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	taskID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if err := config.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error; err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}
