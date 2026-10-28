package controller

import (
	"encoding/json"
	"net/http"

	"calculator-backend/internal/model"
	"calculator-backend/internal/service"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {

	var request model.AddRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := service.Add(request.A, request.B)

	response := model.AddResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
func SubtractHandler(w http.ResponseWriter, r *http.Request) {

	var request model.SubtractRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := service.Subtract(request.A, request.B)

	response := model.SubtractResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {

	var request model.MultiplyRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := service.Multiply(request.A, request.B)

	response := model.MultiplyResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
func DivideHandler(w http.ResponseWriter, r *http.Request) {

	var request model.DivideRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := service.Divide(request.A, request.B)

	response := model.DivideResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}