package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"teddy_bears_api_v2/models"

	"github.com/go-chi/chi/v5"
)

type responseOneLocation struct {
	Location *models.PicnicLocation `json:"location"`
}

type responseAllLocation struct {
	Locations []models.PicnicLocation `json:"locations"`
}

func (h Handler) location(r Router) {
	r.get("/", h.listAllLocations)
	r.get("/{id}", h.fetchLocationById)
	r.put("/{id}", h.updateLocationById)
	r.post("/", h.createLocation)
	r.delete("/{id}", h.deleteLocationById)
}

// @Summary		List all picnic locations
// @Description	List all picnic locations
// @Tags		location
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllLocation
// @Failure		500			{object}	routes.responseError
// @Router		/location 	[GET]
func (h Handler) listAllLocations(w http.ResponseWriter, r *http.Request) {
	// get values from db
	locations, err := h.Logic.ListLocations()
	if err != nil {
		slog.Error("error getting all objects from db", "error", err)
		encode(
			w,
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	encode(w, http.StatusOK, responseAllLocation{Locations: locations})
}

// @Summary		Fetch a picnic location by id
// @Description	Fetch a picnic location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int	true	"Location ID"
// @Success		200				{object}	routes.responseOneLocation
// @Failure		400				{object}	routes.responseError
// @Failure		500				{object}	routes.responseError
// @Router		/location/{id}	[GET]
func (h Handler) fetchLocationById(w http.ResponseWriter, r *http.Request) {
	// get and validate id
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 {
		slog.Error("error getting id")
		encode(w, http.StatusBadRequest, responseError{Error: "Not a valid id"})
		return
	}

	// get value from db
	location, err := h.Logic.FetchLocationByID(id)
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
	encode(w, http.StatusOK, responseOneLocation{Location: location})
}

// @Summary		Update a picnic location by id
// @Description	Update a picnic location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int							true	"Location ID"
// @Param		location		body		models.PicnicLocationInput	true	"Location Object"
// @Success		200				{object}	routes.responseOneLocation
// @Failure		400				{object}	routes.responseError
// @Failure		500				{object}	routes.responseError
// @Router		/location/{id}	[PUT]
func (h Handler) updateLocationById(w http.ResponseWriter, r *http.Request) {
	// get and validate id
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 {
		slog.Error("error getting id")
		encode(w, http.StatusBadRequest, responseError{Error: "Not a valid id"})
		return
	}

	// get and validate body as object
	inputLocation, err := decode[models.PicnicLocationInput](r)
	if err != nil {
		slog.Error("BodyParser error", "error", err)
		encode(
			w,
			http.StatusBadRequest,
			responseError{Error: "missing values or malformed body"},
		)
		return
	}

	// update db
	location, err := h.Logic.UpdateLocationByID(id, inputLocation)
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
	encode(w, http.StatusOK, responseOneLocation{Location: location})
}

// @Summary		Create a picnic location
// @Description	Create a picnic location
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		location	body		models.PicnicLocationInput	true	"Location Object"
// @Success		201			{object}	routes.responseID
// @Failure		400			{object}	routes.responseError
// @Failure		500			{object}	routes.responseError
// @Router		/location	[POST]
func (h Handler) createLocation(w http.ResponseWriter, r *http.Request) {
	// get and validate body as object
	inputLocation, err := decode[models.PicnicLocationInput](r)
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
	id, err := h.Logic.CreateLocation(inputLocation)
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

// @Summary		Delete a location by id
// @Description	Delete a location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int	true	"Location ID"
// @Success		202				{object}	routes.responseMessage
// @Failure		500				{object}	routes.responseError
// @Failure		400				{object}	routes.responseError
// @Router		/location/{id} 	[DELETE]
func (h Handler) deleteLocationById(w http.ResponseWriter, r *http.Request) {
	// get and validate id
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if id == 0 {
		slog.Error("error getting id")
		encode(w, http.StatusBadRequest, responseError{Error: "Not a valid id"})
		return
	}

	// remove from db
	if err := h.Logic.DeleteLocationByID(id); err != nil {
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
