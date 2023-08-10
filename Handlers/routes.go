// handlers/routes.go

package handlers

import (
    "github.com/gorilla/mux"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(r *mux.Router) {
    setupAuthRoutes(r)
    setupKYCRoutes(r) // Add this line to integrate KYC routes
    // ...
}
