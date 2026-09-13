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

		response, err := http.Post(
			apiURL+"/employees",
			"application/json",
			bytes.NewBuffer(requestBody),
		)

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

func GetEmployeesHandler(apiURL string) gin.HandlerFunc {
	return func(c *gin.Context) {

		response, err := http.Get(
			apiURL + "/employees",
		)

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

		response, err := http.Get(
			apiURL + "/employees/" + employeeID,
		)

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

		response, err := http.NewRequest(
			http.MethodPut,
			apiURL+"/employees/"+employeeID,
			bytes.NewBuffer(requestBody),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create request",
			})
			return
		}

		response.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		result, err := client.Do(response)

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

		request, err := http.NewRequest(
			http.MethodDelete,
			apiURL+"/employees/"+employeeID,
			nil,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create request",
			})
			return
		}

		client := &http.Client{}

		result, err := client.Do(request)

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

		response, err := http.Get(searchURL)

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
