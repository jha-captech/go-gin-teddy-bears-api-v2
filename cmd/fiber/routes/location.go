package routes

import (
	"log/slog"
	"net/http"

	"teddy_bears_api_v2/logic"

	"github.com/gofiber/fiber/v2"
)

type responseOneLocation struct {
	Location logic.PicnicLocationReturn `json:"location"`
}

type responseAllLocation struct {
	Locations []logic.PicnicLocationReturn `json:"locations"`
}

func (h Handler) location(r fiber.Router) {
	r.Get("/", h.listAllLocations)
	r.Get("/:id", h.fetchLocationById)
	r.Put("/:id", h.updateLocationById)
	r.Post("/", h.createLocation)
	r.Delete("/:id", h.deleteLocationById)
}

// @Summary		List all picnic locations
// @Description	List all picnic locations
// @Tags		location
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllLocation
// @Failure		500			{object}	routes.responseError
// @Router		/location 	[GET]
func (h Handler) listAllLocations(c *fiber.Ctx) error {
	// get values from db
	locations, err := h.Logic.ListLocations()
	if err != nil {
		slog.Error("error getting all objects from db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error retrieving data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseAllLocation{Locations: locations})
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
func (h Handler) fetchLocationById(c *fiber.Ctx) error {
	// get and validate id
	id, err := c.ParamsInt("id")
	if err != nil {
		slog.Error("error getting id", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid id"})
	}

	// get value from db
	location, err := h.Logic.FetchLocationByID(id)
	if err != nil {
		slog.Error("error getting a objects from db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error retrieving data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseOneLocation{Location: location})
}

// @Summary		Update a picnic location by id
// @Description	Update a picnic location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int							true	"Location ID"
// @Param		location		body		logic.PicnicLocationInput	true	"Location Object"
// @Success		200				{object}	routes.responseOneLocation
// @Failure		400				{object}	routes.responseError
// @Failure		500				{object}	routes.responseError
// @Router		/location/{id}	[PUT]
func (router Handler) updateLocationById(c *fiber.Ctx) error {
	// get and validate id
	id, err := c.ParamsInt("id")
	if err != nil {
		slog.Error("error getting id", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid id"})
	}

	// get and validate body as object
	var inputLocation logic.PicnicLocationInput
	if err := c.BodyParser(&inputLocation); err != nil {
		slog.Error("BodyParser error", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "missing values or malformed body"})
	}

	// update db
	location, err := router.Logic.UpdateLocationByID(id, inputLocation)
	if err != nil {
		slog.Error("error updating object in db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error updating data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseOneLocation{Location: location})
}

// @Summary		Create a picnic location
// @Description	Create a picnic location
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		location	body		logic.PicnicLocationInput	true	"Location Object"
// @Success		201			{object}	routes.responseID
// @Failure		400			{object}	routes.responseError
// @Failure		500			{object}	routes.responseError
// @Router		/location	[POST]
func (router Handler) createLocation(c *fiber.Ctx) error {
	// get and validate body as object
	var inputLocation logic.PicnicLocationInput
	if err := c.BodyParser(&inputLocation); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "missing values or malformed body"})
	}

	// add to db
	id, err := router.Logic.CreateLocation(inputLocation)
	if err != nil {
		slog.Error("error adding object to db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error updating data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseID{ObjectID: id})
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
func (router Handler) deleteLocationById(c *fiber.Ctx) error {
	// get and validate id
	id, err := c.ParamsInt("id")
	if err != nil {
		slog.Error("error getting id", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid id"})
	}

	// remove from db
	if err := router.Logic.DeleteLocationByID(id); err != nil {
		slog.Error("error deleting object from db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error adding data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseMessage{Message: "object successful deleted"})
}
