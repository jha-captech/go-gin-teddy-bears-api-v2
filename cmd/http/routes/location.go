package routes

import (
	"log/slog"
	"net/http"
	"strconv"

	"teddy_bears_api_v2/models"

	"github.com/gin-gonic/gin"
)

type responseOneLocation struct {
	Location *models.PicnicLocation `json:"location"`
}

type responseAllLocation struct {
	Locations []models.PicnicLocation `json:"locations"`
}

func (router Router) location(r *gin.RouterGroup) {
	r.GET("/", router.listAllLocations)
	r.GET("/:id", router.fetchLocationById)
	r.PUT("/:id", router.updateLocationById)
	r.POST("/", router.createLocation)
	r.DELETE("/:id", router.deleteLocationById)
}

// @Summary		List all picnic locations
// @Description	List all picnic locations
// @Tags		location
// @Accept		json
// @Produce		json
// @Success		200			{object}	routes.responseAllLocation
// @Failure		500			{object}	routes.error
// @Router		/location 	[GET]
func (router Router) listAllLocations(c *gin.Context) {
	// get values from db
	locations, err := router.Logic.ListLocations()
	if err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseAllLocation{Locations: locations},
	)
}

// @Summary		Fetch a picnic location by id
// @Description	Fetch a picnic location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int	true	"Location ID"
// @Success		200				{object}	routes.responseOneLocation
// @Failure		400				{object}	routes.error
// @Failure		500				{object}	routes.error
// @Router		/location/{id}	[GET]
func (router Router) fetchLocationById(c *gin.Context) {
	// get and validate id
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		slog.Error("error getting id", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid id"},
		)
		return
	}

	// get value from db
	location, err := router.Logic.FetchLocationByID(id)
	if err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error retrieving data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseOneLocation{Location: location},
	)
}

// @Summary		Update a picnic location by id
// @Description	Update a picnic location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int							true	"Location ID"
// @Param		location		body		models.PicnicLocationInput	true	"Location Object"
// @Success		200				{object}	routes.responseOneLocation
// @Failure		400				{object}	routes.error
// @Failure		500				{object}	routes.error
// @Router		/location/{id}	[PUT]
func (router Router) updateLocationById(c *gin.Context) {
	// get and validate id
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		slog.Error("error getting id", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid id"},
		)
		return
	}

	// get and validate body as object
	var inputLocation models.PicnicLocationInput
	if err := c.ShouldBindJSON(&inputLocation); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "missing values or malformed body"},
		)
		return
	}

	// update db
	location, err := router.Logic.UpdateLocationByID(id, inputLocation)
	if err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error updating data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseOneLocation{Location: location},
	)
}

// @Summary		Create a picnic location
// @Description	Create a picnic location
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		location	body		models.PicnicLocationInput	true	"Location Object"
// @Success		201			{object}	routes.responseID
// @Failure		400			{object}	routes.error
// @Failure		500			{object}	routes.error
// @Router		/location	[POST]
func (router Router) createLocation(c *gin.Context) {
	// get and validate body as object
	var inputLocation models.PicnicLocationInput
	if err := c.ShouldBindJSON(&inputLocation); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "missing values or malformed body"},
		)
		return
	}

	// add to db
	id, err := router.Logic.CreateLocation(inputLocation)
	if err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error updating data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseID{ObjectID: id},
	)
}

// @Summary		Delete a location by id
// @Description	Delete a location by id
// @Tags		location
// @Accept		json
// @Produce		json
// @Param		id				path		int	true	"Location ID"
// @Success		202				{object}	routes.responseMessage
// @Failure		500				{object}	routes.error
// @Failure		400				{object}	routes.error
// @Router		/location/{id} 	[DELETE]
func (router Router) deleteLocationById(c *gin.Context) {
	// get and validate id
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		slog.Error("error getting id", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid id"},
		)
		return
	}

	// remove from db
	if err := router.Logic.DeleteLocationByID(id); err != nil {
		slog.Error("error getting all locations", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error adding data"},
		)
		return
	}

	// return response
	c.JSON(
		http.StatusOK,
		responseMessage{Message: "object successful deleted"},
	)
}
