package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeRequest struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Department string  `json:"department" binding:"required"`
	Position   string  `json:"position" binding:"required"`
	Salary     float64 `json:"salary" binding:"required"`
}

func CreateEmployeeHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request CreateEmployeeRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			return
		}

		requestBody, err := json.Marshal(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create request",
			})
			return
		}

		req, err := createEmployeeAPIRequest(
			c,
			http.MethodPost,
			apiURL+"/employees",
			bytes.NewBuffer(requestBody),
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to contact employee service",
			})
			return
		}

		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			response.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func GetEmployeesHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Cognito access token from cookie
		tokenString, err := c.Cookie("ACCESS_TOKEN")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		// Create request to API Gateway
		req, err := http.NewRequest(
			http.MethodGet,
			apiURL+"/employees",
			nil,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create employee service request",
			})
			return
		}

		// Forward Cognito access token
		req.Header.Set(
			"Authorization",
			"Bearer "+tokenString,
		)

		// Send request
		response, err := http.DefaultClient.Do(req)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to communicate with employee service",
			})
			return
		}

		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			response.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func GetEmployeeByIDHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		employeeID := c.Param("id")

		req, err := createEmployeeAPIRequest(
			c,
			http.MethodGet,
			apiURL+"/employees/"+employeeID,
			nil,
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		response, err := http.DefaultClient.Do(req)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to communicate with employee service",
			})
			return
		}

		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			response.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func UpdateEmployeeHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		employeeID := c.Param("id")

		var request struct {
			Name       string  `json:"name" binding:"required"`
			Email      string  `json:"email" binding:"required,email"`
			Department string  `json:"department" binding:"required"`
			Position   string  `json:"position" binding:"required"`
			Salary     float64 `json:"salary" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			return
		}

		requestBody, err := json.Marshal(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create request",
			})
			return
		}

		req, err := createEmployeeAPIRequest(
			c,
			http.MethodPut,
			apiURL+"/employees/"+employeeID,
			bytes.NewBuffer(requestBody),
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		result, err := http.DefaultClient.Do(req)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to communicate with employee service",
			})
			return
		}

		defer result.Body.Close()

		responseBody, err := io.ReadAll(result.Body)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			result.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func DeleteEmployeeHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		employeeID := c.Param("id")

		req, err := createEmployeeAPIRequest(
			c,
			http.MethodDelete,
			apiURL+"/employees/"+employeeID,
			nil,
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		result, err := http.DefaultClient.Do(req)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to communicate with employee service",
			})
			return
		}

		defer result.Body.Close()

		responseBody, err := io.ReadAll(result.Body)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			result.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func SearchEmployeesHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		query := c.Query("q")

		if strings.TrimSpace(query) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "search query is required",
			})
			return
		}

		searchURL := apiURL + "/employees/search?q=" +
			url.QueryEscape(query)

		req, err := createEmployeeAPIRequest(
			c,
			http.MethodGet,
			searchURL,
			nil,
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		response, err := http.DefaultClient.Do(req)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "failed to contact employee service",
			})
			return
		}

		defer response.Body.Close()

		responseBody, err := io.ReadAll(response.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to read employee service response",
			})
			return
		}

		c.Data(
			response.StatusCode,
			"application/json",
			responseBody,
		)
	}
}

func createEmployeeAPIRequest(
	c *gin.Context,
	method string,
	requestURL string,
	body io.Reader,
) (*http.Request, error) {

	tokenString, err := c.Cookie("ACCESS_TOKEN")

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		method,
		requestURL,
		body,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+tokenString,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	return req, nil
}
