package routes

import (
	"calculator-backend/internal/controller"
	"net/http"
)

func SetupRoutes() {

	http.HandleFunc("/add", controller.AddHandler)
	http.HandleFunc("/subtract", controller.SubtractHandler)
	http.HandleFunc("/multiply", controller.MultiplyHandler)
	http.HandleFunc("/divide", controller.DivideHandler)

}