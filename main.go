package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Sergio",
		Content: "Programming",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("An error has occurred getting request body", err)
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println("An error has ocurred getting the task id", err)
		return
	}

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "not valid id")
	}

	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Task removed succesfully")
		}
	}
}

func updatedTask(w http.ResponseWriter, r *http.Request) {
	var change task
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "not valid id")
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Body not valid")
	}

	json.Unmarshal(reqBody, &change)

	for _, task := range tasks {
		if task.ID == taskId {
			task.Content = change.Content
			task.Name = change.Name
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
		}
	}
}

func main() {
	// declare router
	router := mux.NewRouter().StrictSlash(true)

	// index route
	router.HandleFunc("/", indexRoute)

	// get tasks
	router.HandleFunc("/tasks", getTasks).Methods("GET")

	// get task by id
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")

	// create task
	router.HandleFunc("/tasks", createTask).Methods("POST")

	// delete task
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// updated task
	router.HandleFunc("/tasks/{id}", updatedTask).Methods("PUT")

	// start server
	log.Fatal(http.ListenAndServe(":3000", router))
}
