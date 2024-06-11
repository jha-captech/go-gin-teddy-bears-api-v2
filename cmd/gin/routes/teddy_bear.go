package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"teddy_bears_api_v2/logic"

	"github.com/gin-gonic/gin"
)

const MaxLimit = 15

type responseOneTeddyBear struct {
	TeddyBear logic.TeddyBearReturn `json:"location"`
}

type responseAllTeddyBear struct {
	TeddyBears []logic.TeddyBearReturn `json:"locations"`
}

func (router Handler) teddyBear(r *gin.RouterGroup) {
	r.GET("/", router.listAllTeddyBears)
	r.GET("/paginated", router.listPaginatedTeddyBears)
	r.GET("/:name", router.fetchTeddyBearByName)
	r.PUT("/:name", router.updateTeddyBearByName)
	r.POST("/", router.createTeddyBear)
	r.DELETE("/:name", router.deleteTeddyBearByName)
}

// @Summary		List all teddy bears
// @Description	List all teddy bears
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllTeddyBear
// @Failure		500			{object}	routes.responseError
// @Router		/teddy-bear [GET]
func (router Handler) listAllTeddyBears(c *gin.Context) {
	// get values from db
	bears, err := router.Logic.ListTeddyBears()
	if err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseAllTeddyBear{TeddyBears: bears},
	)
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
func (router Handler) listPaginatedTeddyBears(c *gin.Context) {
	// Parse query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

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
	bears, err := router.Logic.ListPaginatedTeddyBears(offset, limit)
	if err != nil {
		slog.Error("error getting paginated teddy bears", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// Return paginated response
	c.JSON(
		http.StatusOK,
		responseAllTeddyBear{TeddyBears: bears},
	)
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
func (router Handler) fetchTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// get value from db
	bear, err := router.Logic.FetchTeddyBearByName(name)
	if err != nil {
		slog.Error("error getting bear by name", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseOneTeddyBear{TeddyBear: bear},
	)
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
func (router Handler) updateTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// get and validate body as object
	var inputBear logic.TeddyBearInput
	if err := c.ShouldBindJSON(&inputBear); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			responseError{Error: "missing values or malformed body"},
		)
		return
	}

	// get value from db
	bear, err := router.Logic.UpdateTeddyBearByName(name, inputBear)
	if err != nil {
		slog.Error("error getting bear by name", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseOneTeddyBear{TeddyBear: bear},
	)
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
func (router Handler) createTeddyBear(c *gin.Context) {
	// get and validate body as object
	var inputBear logic.TeddyBearInput
	if err := c.ShouldBindJSON(&inputBear); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			responseError{Error: "missing values or malformed body"},
		)
		return
	}

	// add to db
	id, err := router.Logic.CreateTeddyBear(inputBear)
	if err != nil {
		slog.Error("error adding teddy bear", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error updating data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseID{ObjectID: id},
	)
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
func (router Handler) deleteTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			responseError{Error: "Not a valid name"},
		)
		return
	}

	// remove from db
	if err := router.Logic.DeleteTeddyBearByName(name); err != nil {
		slog.Error("error deleting resource", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			responseError{Error: "Error adding data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseMessage{Message: "object successful deleted"},
	)
}
