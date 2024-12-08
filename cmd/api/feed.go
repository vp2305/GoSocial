package main

import (
	"SocialMedia/internal/store"
	"net/http"
)

// GetUserFeed godoc
//
//	@Summary		Get user feed
//	@Description	Get user feed respective to the pagination, filters and sort
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int		false	"Limit post per request"
//	@Param			offset	path		int		false	"Offset by the previous post"
//	@Param			sort	path		int		false	"Sort post by asc or desc"
//	@Param			search	path		string	false	"Search by title or content"
//	@Param			tags	path		string	false	"Filter by relative tags"
//	@Success		200		{object}	[]models.PostWithMetadata
//	@Failure		404		{object}	error	"Post not found"
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	// pagination, filters and sort
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, user.ID, fq)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
