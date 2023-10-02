package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/expenses-tracker/pkg/models"
	"io"
	"log"
	"net/http"
)

func (h handler) AddExpense(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var expense models.Expense
	json.Unmarshal(body, &expense)

	if result := h.DB.Create(&expense); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added new expense")
}
