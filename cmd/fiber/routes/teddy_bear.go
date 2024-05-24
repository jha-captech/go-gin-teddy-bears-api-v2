package routes

import (
	"log/slog"
	"net/http"

	"teddy_bears_api_v2/logic"

	"github.com/gofiber/fiber/v2"
)

const MaxLimit = 15

type responseOneTeddyBear struct {
	TeddyBear logic.TeddyBearReturn `json:"location"`
}

type responseAllTeddyBear struct {
	TeddyBears []logic.TeddyBearReturn `json:"locations"`
}

func (h Handler) teddyBear(r fiber.Router) {
	r.Get("/", h.listAllTeddyBears)
	r.Get("/paginated", h.listPaginatedTeddyBears)
	r.Get("/:name", h.fetchTeddyBearByName)
	r.Put("/:name", h.updateTeddyBearByName)
	r.Post("/", h.createTeddyBear)
	r.Delete("/:name", h.deleteTeddyBearByName)
}

// @Summary		List all teddy bears
// @Description	List all teddy bears
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllTeddyBear
// @Failure		500			{object}	routes.responseError
// @Router		/teddy-bear [GET]
func (h Handler) listAllTeddyBears(c *fiber.Ctx) error {
	// get values from db
	bears, err := h.Logic.ListTeddyBears()
	if err != nil {
		slog.Error("error getting a object from db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error retrieving data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseAllTeddyBear{TeddyBears: bears})
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
func (h Handler) listPaginatedTeddyBears(c *fiber.Ctx) error {
	// Parse query parameters
	page, _ := c.ParamsInt("page")
	limit, _ := c.ParamsInt("limit")

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
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error retrieving data"})
	}

	// Return paginated response
	return c.
		Status(http.StatusOK).
		JSON(responseAllTeddyBear{TeddyBears: bears})
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
func (h Handler) fetchTeddyBearByName(c *fiber.Ctx) error {
	// get and validate name
	name := c.Params("name")
	if name == "" {
		slog.Error("error not a valid name")
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid name"})
	}

	// get value from db
	bear, err := h.Logic.FetchTeddyBearByName(name)
	if err != nil {
		slog.Error("error getting a objects from db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error retrieving data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseOneTeddyBear{TeddyBear: bear})
}

// @Summary		Update a teddy bear by name
// @Description	Update a teddy bear by name
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		name				path		string	true	"Teddy Bear Name"
// @Param		teddyBear			body		logic.TeddyBearInput	true	"Teddy Bear Object"
// @Success		200					{object}	routes.responseOneTeddyBear
// @Failure		500					{object}	routes.responseError
// @Failure		422					{object}	routes.responseError
// @Router		/teddy-bear/{name}	[PUT]
func (h Handler) updateTeddyBearByName(c *fiber.Ctx) error {
	// get and validate name
	name := c.Params("name")
	if name == "" {
		slog.Error("error not a valid name")
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid name"})
	}

	// get and validate body as object
	var inputBear logic.TeddyBearInput
	if err := c.BodyParser(&inputBear); err != nil {
		slog.Error("BodyParser error", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "missing values or malformed body"})
	}

	// get value from db
	bear, err := h.Logic.UpdateTeddyBearByName(name, inputBear)
	if err != nil {
		slog.Error("error updating object in db", "error", err)
		return c.
			Status(http.StatusInternalServerError).
			JSON(responseError{Error: "Error updating data"})
	}

	// return response
	return c.
		Status(http.StatusOK).
		JSON(responseOneTeddyBear{TeddyBear: bear})
}

// @Summary		Create a teddy bear
// @Description	Create a teddy bear
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		teddyBear	body		logic.TeddyBearInput	true	"Teddy Bear Object"
// @Success		201			{object}	routes.responseID
// @Failure		422			{object}	routes.responseError
// @Failure		500			{object}	routes.responseError
// @Failure		409			{object}	routes.responseError
// @Router		/teddy-bear	[POST]
func (h Handler) createTeddyBear(c *fiber.Ctx) error {
	// get and validate body as object
	var inputBear logic.TeddyBearInput
	if err := c.BodyParser(&inputBear); err != nil {
		slog.Error("BodyParser error", "error", err)
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "missing values or malformed body"})
	}

	// add to db
	id, err := h.Logic.CreateTeddyBear(inputBear)
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
func (h Handler) deleteTeddyBearByName(c *fiber.Ctx) error {
	// get and validate name
	name := c.Params("name")
	if name == "" {
		slog.Error("error not a valid name")
		return c.
			Status(http.StatusBadRequest).
			JSON(responseError{Error: "Not a valid name"})
	}

	// remove from db
	if err := h.Logic.DeleteTeddyBearByName(name); err != nil {
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
