package main
import (
	"fmt"
	"net/http"
	"calculator-backend/internal/routes"
)
func main() {
	routes.SetupRoutes()
	fmt.Println("Server is running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
