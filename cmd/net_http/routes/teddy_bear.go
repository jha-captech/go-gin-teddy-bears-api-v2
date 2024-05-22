package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"teddy_bears_api_v2/models"

	"github.com/go-chi/chi/v5"
)

const MaxLimit = 15

type responseOneTeddyBear struct {
	TeddyBear *models.TeddyBearReturn `json:"location"`
}

type responseAllTeddyBear struct {
	TeddyBears []models.TeddyBearReturn `json:"locations"`
}

func (h Handler) teddyBear(r Router) {
	r.get("/", h.listAllTeddyBears)
	r.get("/paginated", h.listPaginatedTeddyBears)
	r.get("/:name", h.fetchTeddyBearByName)
	r.put("/:name", h.updateTeddyBearByName)
	r.post("/", h.createTeddyBear)
	r.delete("/:name", h.deleteTeddyBearByName)
}

// @Summary		List all teddy bears
// @Description	List all teddy bears
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllTeddyBear
// @Failure		500			{object}	routes.responseError
// @Router		/teddy-bear [GET]
func (h Handler) listAllTeddyBears(w http.ResponseWriter, r *http.Request) {
	// get values from db
	bears, err := h.Logic.ListTeddyBears()
	if err != nil {
		slog.Error("error getting a object from db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseAllTeddyBear{TeddyBears: bears})
}

// @Summary		List all teddy bears with page and limit
// @Description	List all teddy bears with page and limit
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		page		query		int	false	"Page number (default 1)"
// @Param		limit		query		int	false	"Items per page (default 10, max 15)"
// @Success		200			{object}	routes.responseAllTeddyBear
// @Failure		500			{object}	routes.responseError
// @Router		/teddy-bear/paginated 	[GET]
func (h Handler) listPaginatedTeddyBears(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	page, _ := strconv.Atoi(chi.URLParam(r, "page"))
	limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))

	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	} else if limit > MaxLimit {
		limit = MaxLimit
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get paginated teddy bears from the database
	bears, err := h.Logic.ListPaginatedTeddyBears(offset, limit)
	if err != nil {
		slog.Error("error getting all objects from db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// Return paginated response
	encode(w, http.StatusOK, responseAllTeddyBear{TeddyBears: bears})
}

// @Summary		Fetch a teddy bear by name
// @Description	Fetch a teddy bear by name
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		name				path		string	true	"Teddy Bear Name"
// @Success		200					{object}	routes.responseOneTeddyBear
// @Failure		400					{object}	routes.responseError
// @Failure		500					{object}	routes.responseError
// @Router		/teddy-bear/{name}	[GET]
func (h Handler) fetchTeddyBearByName(w http.ResponseWriter, r *http.Request) {
	// get and validate name
	name := chi.URLParam(r, "name")
	if name == "" {
		slog.Error("error not a valid name")
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// get value from db
	bear, err := h.Logic.FetchTeddyBearByName(name)
	if err != nil {
		slog.Error("error getting a objects from db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseOneTeddyBear{TeddyBear: bear})
}

// @Summary		Update a teddy bear by name
// @Description	Update a teddy bear by name
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		name				path		string	true	"Teddy Bear Name"
// @Param		teddyBear			body		models.TeddyBearInput	true	"Teddy Bear Object"
// @Success		200					{object}	routes.responseOneTeddyBear
// @Failure		500					{object}	routes.responseError
// @Failure		422					{object}	routes.responseError
// @Router		/teddy-bear/{name}	[PUT]
func (h Handler) updateTeddyBearByName(w http.ResponseWriter, r *http.Request) {
	// get and validate name
	name := chi.URLParam(r, "name")
	if name == "" {
		slog.Error("error not a valid name")
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// get and validate body as object
	inputBear, err := decode[models.TeddyBearInput](r)
	if err != nil {
		slog.Error("BodyParser error", "error", err)
		encode(
			w,
			http.StatusBadRequest,
			responseError{Error: "missing values or malformed body"},
		)
		return
	}

	// get value from db
	bear, err := h.Logic.UpdateTeddyBearByName(name, inputBear)
	if err != nil {
		slog.Error("error updating object in db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error updating data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseOneTeddyBear{TeddyBear: bear})
}

// @Summary		Create a teddy bear
// @Description	Create a teddy bear
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		teddyBear	body		models.TeddyBearInput	true	"Teddy Bear Object"
// @Success		201			{object}	routes.responseID
// @Failure		422			{object}	routes.responseError
// @Failure		500			{object}	routes.responseError
// @Failure		409			{object}	routes.responseError
// @Router		/teddy-bear	[POST]
func (h Handler) createTeddyBear(w http.ResponseWriter, r *http.Request) {
	// get and validate body as object
	inputBear, err := decode[models.TeddyBearInput](r)
	if err != nil {
		slog.Error("BodyParser error", "error", err)
		encode(
			w,
			http.StatusBadRequest,
			responseError{Error: "missing values or malformed body"},
		)
		return
	}

	// add to db
	id, err := h.Logic.CreateTeddyBear(inputBear)
	if err != nil {
		slog.Error("error adding object to db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error updating data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseID{ObjectID: id})
}

// @Summary		Delete a teddy bear by name
// @Description	Delete a teddy bear by name
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		name				path		string	true	"Teddy Bear Name"
// @Success		202					{object}	routes.responseMessage
// @Failure		500					{object}	routes.responseError
// @Failure		404					{object}	routes.responseError
// @Router		/teddy-bear/{name}	[DELETE]
func (h Handler) deleteTeddyBearByName(w http.ResponseWriter, r *http.Request) {
	// get and validate name
	name := chi.URLParam(r, "name")
	if name == "" {
		slog.Error("error not a valid name")
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// remove from db
	if err := h.Logic.DeleteTeddyBearByName(name); err != nil {
		slog.Error("error deleting object from db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error adding data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseMessage{Message: "object successful deleted"})
}
