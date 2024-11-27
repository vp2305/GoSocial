package main

import (
	"net/http"
)

// GetHealth godoc
//
//	@Summary		Get health
//	@Description	Get configuration information of the service
//	@Tags			ops
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
