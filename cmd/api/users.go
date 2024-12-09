package main

import (
	"SocialMedia/internal/models"
	"SocialMedia/internal/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "users"

// Getuser godoc
//
//	@Summary		User profile
//	@Description	Fetches profile of the current user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		404	{object}	error	"Invalid request"
//	@Security		ApiKeyAuth
//	@Router			/user [get]
func (app *application) getUserProfile(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetUserById godoc
//
//	@Summary		Profile by ID
//	@Description	Fetches user profile by given ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int	true	"Target User ID"
//	@Success		200		{object}	models.User
//	@Failure		404		{object}	error	"Invalid request"
//	@Failure		400		{object}	error	"Malformed param"
//	@Security		ApiKeyAuth
//	@Router			/user/{userID} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.store.Users.GetByID(r.Context(), userId)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// FollowUser godoc
//
//	@Summary		Follow a user
//	@Description	Follow a user by providing the target user's ID in the path and the current user's ID in the request body.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int		true	"Target User ID"
//	@Success		204		{string}	string	"User followed successfully"
//	@Failure		409		{object}	error	"Already following the user"
//	@Failure		404		{object}	error	"Target user not found"
//	@Failure		400		{object}	error	"Invalid request"
//	@Security		ApiKeyAuth
//	@Router			/user/{userID}/follow [put]
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)
	followedID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, user.ID, followedID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// UnfollowUser godoc
//
//	@Summary		Unfollow a user
//	@Description	Unfollow a user by providing the target user's ID in the path and the current user's ID in the request body.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int		true	"Target User ID"
//	@Success		204		{string}	string	"User unfollowed successfully"
//	@Failure		404		{object}	error	"Target user not found"
//	@Failure		400		{object}	error	"Invalid request"
//	@Security		ApiKeyAuth
//	@Router			/user/{userID}/unfollow [put]
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)
	unfollowedID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.store.Followers.UnFollow(r.Context(), user.ID, unfollowedID); err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// ActivateUser godoc
//
//	@Summary		Activates/Register a user
//	@Description	Activates/Register a user by invitation token
//	@Tags			authentication
//	@Produce		json
//	@Param			token	path		string	true	"Invitation token"
//	@Success		204		{string}	string	"User activated"
//	@Failure		404		{object}	error	"No users found associated to the token provided"
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/user/activate/{token} [put]
func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	if err := app.store.Users.Activate(r.Context(), token); err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, ""); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func getUserFromCtx(r *http.Request) *models.User {
	user, _ := r.Context().Value(userCtx).(*models.User)

	return user
}
