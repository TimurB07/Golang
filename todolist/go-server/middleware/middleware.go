package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-server/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const connectionString = "host=localhost user=postgres dbname=ToDoList password=12345 port=5432 sslmode=disable"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	fmt.Println("Connected to PostgreSQL!")
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.ToDoList
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if task.Task == "" {
		http.Error(w, "Task cannot be empty", http.StatusBadRequest)
		return
	}

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTask()
	json.NewEncoder(w).Encode(count)
}

func getAllTask() []models.ToDoList {
	var tasks []models.ToDoList
	db.Find(&tasks)
	return tasks
}

func insertOneTask(task models.ToDoList) {
	result := db.Create(&task)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Inserted a Single Record ", task.ID)
}

func taskComplete(taskID string) {
	var task models.ToDoList
	if err := db.First(&task, taskID).Error; err != nil {
		log.Fatal(err)
	}
	task.Status = true
	db.Save(&task)
	fmt.Println("Task marked as complete:", task.ID)
}

func undoTask(taskID string) {
	var task models.ToDoList
	if err := db.First(&task, taskID).Error; err != nil {
		log.Fatal(err)
	}
	task.Status = false
	db.Save(&task)
	fmt.Println("Task marked as incomplete:", task.ID)
}

func deleteOneTask(taskID string) {
	if err := db.Delete(&models.ToDoList{}, taskID).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Task with ID:", taskID)
}

func deleteAllTask() int64 {
	result := db.Delete(&models.ToDoList{}, "1 = 1")
	return result.RowsAffected
}
