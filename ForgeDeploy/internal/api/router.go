package api

import (
	"net/http"

	"forgedeploy/internal/api/handlers"
)

func NewRouter(ph *handlers.PipelineHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/pipelines", ph.GetPipelines)

	return mux
}
