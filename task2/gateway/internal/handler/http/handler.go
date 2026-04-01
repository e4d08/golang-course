package http

import (
	"encoding/json"
	"errors"
	"gateway/internal/domain"
	"gateway/internal/usecase"
	"log"
	"net/http"
)

type Handler struct {
	service *usecase.GatewayService
}

func NewHandler(service *usecase.GatewayService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to write json: %s", err)
	}
}

func (h *Handler) WriteError(w http.ResponseWriter, status int, message string) {
	resp := ErrorResponse{Error: message}
	h.WriteJson(w, status, resp)
}

// GetRepository godoc
//
//	@Summary		Get repository by owner and name
//	@Description	get GitHub repository information by its owner and name
//	@Accept			json
//	@Produce		json
//	@Param			owner	query		string	true	"Repository owner"
//	@Param			name	query		string	true	"Repository name"
//	@Success		200		{object}	GetRepositoryResponse
//	@Failure		400		{object}	ErrorResponse	"Bad parameters"
//	@Failure		404		{object}	ErrorResponse	"Repository not found"
//	@Failure		500		{object}	ErrorResponse	"Internal server error"
//	@Router			/repository/get [get]
func (h *Handler) GetRepository(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := r.URL.Query()
	owner := query.Get("owner")
	name := query.Get("name")

	if owner == "" {
		h.WriteError(w, http.StatusBadRequest, "owner is required")
		return
	}
	if name == "" {
		h.WriteError(w, http.StatusBadRequest, "name is required")
		return
	}

	repo, err := h.service.GetRepository(ctx, owner, name)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRepositoryNotFound):
			h.WriteError(w, http.StatusNotFound, "repository not found")
		case errors.Is(err, domain.ErrInternal):
			h.WriteError(w, http.StatusInternalServerError, "internal server error")
		default:
			h.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	resp := GetRepositoryResponse{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       repo.Stars,
		Forks:       repo.Forks,
		CreatedAt:   repo.CreatedAt,
		License:     repo.License,
	}
	h.WriteJson(w, http.StatusOK, resp)
}
