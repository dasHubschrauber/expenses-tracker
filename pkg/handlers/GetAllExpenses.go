package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/expenses-tracker/pkg/models"
	"net/http"
)

func (h handler) GetAllExpenses(w http.ResponseWriter, r *http.Request) {

	var expenses []models.Expense

	if result := h.DB.Find(&expenses); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expenses)
}
