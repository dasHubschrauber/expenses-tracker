package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/expenses-tracker/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) GetExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var expense models.Expense

	if result := h.DB.Find(&expense, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(expense)
}
