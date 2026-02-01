package handlers

import (
	"encoding/json"
	"net/http"

	"forgedeploy/internal/store"
)

type PipelineHandler struct {
	repo store.PipelineRepository
}

func NewPipelineHandler(repo store.PipelineRepository) *PipelineHandler {
	return &PipelineHandler{repo: repo}
}

func (h *PipelineHandler) GetPipelines(w http.ResponseWriter, r *http.Request) {
	pipelines, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(pipelines)
}
