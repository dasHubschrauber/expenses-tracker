package main

import (
	"github.com/dasHubschrauber/expenses-tracker/pkg/db"
	"github.com/dasHubschrauber/expenses-tracker/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/expenses", h.GetAllExpenses).Methods(http.MethodGet)
	router.HandleFunc("/expense", h.AddExpense).Methods(http.MethodPost)
	router.HandleFunc("/expense/{id}", h.GetExpense).Methods(http.MethodGet)
	router.HandleFunc("/expense/{id}", h.UpdateExpense).Methods(http.MethodPut)
	router.HandleFunc("/expense/{id}", h.DeleteExpense).Methods(http.MethodDelete)

	log.Println("Expenses Tracker is running!")
	http.ListenAndServe(":4000", router)
}
