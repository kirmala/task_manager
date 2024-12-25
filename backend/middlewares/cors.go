package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func CORSHandler(next http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow React frontend
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),          // Allowed headers
	)(next)
}
