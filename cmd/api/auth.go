package main

import (
	"SocialMedia/internal/models"
	"SocialMedia/internal/store"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// RegisterUserHandler godoc
//
//	@Summary		Register a user
//	@Description	Registers a new user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User credentials"
//	@Success		201		{object}	models.User			"User registered"
//	@Failure		400		{object}	error				"Invalid request"
//	@Failure		409		{object}	error				"Username or email already exist"
//	@Failure		500		{object}	error
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	plainToken := uuid.New().String()

	//? Store the hashed invitation token for added security
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.exp); err != nil {
		switch err {
		case store.ErrDuplicateUsername:
			app.conflictResponse(w, r, err)
		case store.ErrDuplicateEmail:
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	// Mail Integration
	app.logger.Infow("user created", "token", plainToken)

	if err := app.jsonResponse(w, http.StatusCreated, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
