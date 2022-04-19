package handlers

import (
	"encoding/json"
	"fmt"
	"lgmontenegro/collantz/internal/service/collatz"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Steps struct {
	TotalSteps int `json:"totalSteps"`
}

func CollatzHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	factor, err := strconv.ParseInt(vars["factor"], 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "wrong input"}`)
		return
	}

	steps := collatz.Collatz(int(factor))
	stepResponse := Steps { TotalSteps: steps}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stepResponse)
	return
}
