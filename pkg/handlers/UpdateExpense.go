package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dasHubschrauber/expenses-tracker/pkg/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var updatedExpense models.Expense
	json.Unmarshal(body, &updatedExpense)

	var expense models.Expense

	if result := h.DB.Find(&expense, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	expense.Title = updatedExpense.Title
	expense.Amount = updatedExpense.Amount
	expense.Desc = updatedExpense.Desc

	h.DB.Save(&expense)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated expense")

}
