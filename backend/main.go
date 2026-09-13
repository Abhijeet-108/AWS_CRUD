package main

import (
	"fmt"
	"log"
	"net/http"

	"employee-management-backend/auth"
	"employee-management-backend/config"
	"employee-management-backend/handlers"
	"employee-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORS())
	cfg := config.Load()

	issuer := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s",
		cfg.CognitoRegion,
		cfg.CognitoUserPoolID,
	)

	jwksURL := issuer + "/.well-known/jwks.json"

	jwtValidator, err := auth.NewJWTValidator(
		jwksURL,
		issuer,
		cfg.CognitoClientID,
	)

	if err != nil {
		log.Fatalf("Failed to create JWT validator: %v", err)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "employee management backend is running",
		})
	})

	router.GET("/api/profile", middleware.AuthMiddleware(jwtValidator), handlers.ProfileHandler)

	router.GET("/api/auth/login", handlers.LoginHandler(cfg))
	router.GET("/api/auth/callback", handlers.CallbackHandler(cfg))
	router.GET("/api/auth/logout", handlers.LogoutHandler(cfg))
	router.GET("/api/auth/logout-complete", handlers.LogoutCompleteHandler)

	router.POST("/api/employees", middleware.AuthMiddleware(jwtValidator), handlers.CreateEmployeeHandler(cfg.EmployeeAPIURL))
	router.GET("/api/employees", middleware.AuthMiddleware(jwtValidator), handlers.GetEmployeesHandler(cfg.EmployeeAPIURL))
	router.GET("/api/employees/search", middleware.AuthMiddleware(jwtValidator), handlers.SearchEmployeesHandler(cfg.EmployeeAPIURL))
	router.GET("/api/employees/:id", middleware.AuthMiddleware(jwtValidator), handlers.GetEmployeeByIDHandler(cfg.EmployeeAPIURL))
	router.PUT("/api/employees/:id", middleware.AuthMiddleware(jwtValidator), handlers.UpdateEmployeeHandler(cfg.EmployeeAPIURL))
	router.DELETE("/api/employees/:id", middleware.AuthMiddleware(jwtValidator), handlers.DeleteEmployeeHandler(cfg.EmployeeAPIURL))
	router.Run(":8080")
}
