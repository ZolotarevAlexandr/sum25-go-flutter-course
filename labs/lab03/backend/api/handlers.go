package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"lab03-backend/models"
	"lab03-backend/storage"

	"github.com/gorilla/mux"
)

var (
	httpCatURL   = "https://http.cat/"
	ErrInvalidID = errors.New("invalid ID")
)

// Handler holds the storage instance
type Handler struct {
	storage *storage.MemoryStorage
}

// NewHandler creates a new handler instance
func NewHandler(storage *storage.MemoryStorage) *Handler {
	return &Handler{storage: storage}
}

// SetupRoutes configures all API routes
func (h *Handler) SetupRoutes() *mux.Router {
	// GET /messages -> h.GetMessages
	// POST /messages -> h.CreateMessage
	// PUT /messages/{id} -> h.UpdateMessage
	// DELETE /messages/{id} -> h.DeleteMessage
	// GET /status/{code} -> h.GetHTTPStatus
	// GET /health -> h.HealthCheck
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/messages", h.GetMessages).Methods("GET")
	api.HandleFunc("/messages", h.CreateMessage).Methods("POST")
	api.HandleFunc("/messages/{id}", h.UpdateMessage).Methods("PUT")
	api.HandleFunc("/messages/{id}", h.DeleteMessage).Methods("DELETE")
	api.HandleFunc("/status/{code}", h.GetHTTPStatus).Methods("GET")
	api.HandleFunc("/health", h.HealthCheck).Methods("GET")
	return router
}

// GetMessages handles GET /api/messages
func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// Get all messages from storage
	// Create successful API response
	// Write JSON response with status 200
	// Handle any errors appropriately

	messages := h.storage.GetAll()
	response := models.APIResponse{
		Success: true,
		Data:    messages,
	}
	writeJSON(w, http.StatusOK, &response)
}

// CreateMessage handles POST /api/messages
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body into CreateMessageRequest
	// Validate the request
	// Create message in storage
	// Create successful API response
	// Write JSON response with status 201
	// Handle validation and storage errors appropriately

	request := models.CreateMessageRequest{}
	if err := parseJSON(r, &request); err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}
	if err := request.Validate(); err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}
	msg, err := h.storage.Create(request.Username, request.Content)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusInternalServerError, &response)
		return
	}
	response := models.APIResponse{
		Success: true,
		Data:    msg,
	}
	writeJSON(w, http.StatusCreated, &response)
}

// UpdateMessage handles PUT /api/messages/{id}
func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path variables
	// Parse JSON request body into UpdateMessageRequest
	// Validate the request
	// Update message in storage
	// Create successful API response
	// Write JSON response with status 200
	// Handle validation, parsing, and storage errors appropriately

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	request := models.UpdateMessageRequest{}
	if err = parseJSON(r, &request); err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	if err = request.Validate(); err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	msg, err := h.storage.Update(id, request.Content)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		if errors.Is(err, storage.ErrMessageNotFound) {
			writeJSON(w, http.StatusNotFound, response)
		} else {
			writeJSON(w, http.StatusInternalServerError, &response)
		}
		return
	}
	response := models.APIResponse{
		Success: true,
		Data:    msg,
	}
	writeJSON(w, http.StatusOK, &response)
}

// DeleteMessage handles DELETE /api/messages/{id}
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path variables
	// Delete message from storage
	// Write response with status 204 (No Content)
	// Handle parsing and storage errors appropriately

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	err = h.storage.Delete(id)
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		if errors.Is(err, storage.ErrMessageNotFound) {
			writeJSON(w, http.StatusNotFound, response)
		} else {
			writeJSON(w, http.StatusInternalServerError, &response)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetHTTPStatus handles GET /api/status/{code}
func (h *Handler) GetHTTPStatus(w http.ResponseWriter, r *http.Request) {
	// Extract status code from URL path variables
	// Validate status code (must be between 100-599)
	// Create HTTPStatusResponse with:
	//   - StatusCode: parsed code
	//   - ImageURL: "https://http.cat/{code}"
	//   - Description: HTTP status description
	// Create successful API response
	// Write JSON response with status 200
	// Handle parsing and validation errors appropriately

	vars := mux.Vars(r)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		response := models.APIResponse{
			Success: false,
			Error:   err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	if code < 100 || code > 599 {
		response := models.APIResponse{
			Success: false,
			Error:   ErrInvalidID.Error(),
		}
		writeJSON(w, http.StatusBadRequest, &response)
		return
	}

	response := models.APIResponse{
		Success: true,
		Data: models.HTTPStatusResponse{
			StatusCode:  code,
			ImageURL:    fmt.Sprint(httpCatURL, code),
			Description: getHTTPStatusDescription(code),
		},
	}
	writeJSON(w, http.StatusOK, &response)
}

// HealthCheck handles GET /api/health
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Create a simple health check response with:
	//   - status: "ok"
	//   - message: "API is running"
	//   - timestamp: current time
	//   - total_messages: count from storage
	// Write JSON response with status 200

	response := map[string]any{
		"status":         "ok",
		"message":        "API is running",
		"timestamp":      time.Now(),
		"total_messages": len(h.storage.GetAll()),
	}
	writeJSON(w, http.StatusOK, response)
}

// Helper function to write error responses
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error writing JSON: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Helper function to parse JSON request body
func parseJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(dst)
}

// Helper function to get HTTP status description
func getHTTPStatusDescription(code int) string {
	// Return appropriate description for common HTTP status codes
	// 200: "OK", 201: "Created", 204: "No Content"
	// 400: "Bad Request", 401: "Unauthorized", 404: "Not Found"
	// 500: "Internal Server Error", etc.
	// Return "Unknown Status" for unrecognized codes

	desc := http.StatusText(code)
	if desc == "" {
		return "Unknown Status"
	}
	return desc
}
