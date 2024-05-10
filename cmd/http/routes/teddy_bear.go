package routes

import (
	"log/slog"
	"net/http"

	"teddy_bears_api_v2/models"

	"github.com/gin-gonic/gin"
)

type responseOneTeddyBear struct {
	TeddyBear *models.TeddyBearReturn `json:"location"`
}

type responseAllTeddyBear struct {
	TeddyBears []models.TeddyBearReturn `json:"locations"`
}

func (router Router) teddyBear(r *gin.RouterGroup) {
	r.GET("/", router.listAllTeddyBears)
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
// @Failure		500			{object}	routes.error
// @Router		/teddy-bear [GET]
func (router Router) listAllTeddyBears(c *gin.Context) {
	// get values from db
	bears, err := router.Logic.ListTeddyBears()
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
// @Failure		400					{object}	routes.error
// @Failure		500					{object}	routes.error
// @Router		/teddy-bear/{name}	[GET]
func (router Router) fetchTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid name"},
		)
		return
	}

	// get value from db
	bear, err := router.Logic.FetchTeddyBearByName(name)
	if err != nil {
		slog.Error("error getting bear by name", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error retrieving data"},
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
// @Param		teddyBear			body		models.TeddyBearInput	true	"Teddy Bear Object"
// @Success		200					{object}	routes.responseOneTeddyBear
// @Failure		500					{object}	routes.error
// @Failure		422					{object}	routes.error
// @Router		/teddy-bear/{name}	[PUT]
func (router Router) updateTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid name"},
		)
		return
	}

	// get and validate body as object
	var inputBear models.TeddyBearInput
	if err := c.ShouldBindJSON(&inputBear); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "missing values or malformed body"},
		)
		return
	}

	// get value from db
	bear, err := router.Logic.UpdateTeddyBearByName(name, inputBear)
	if err != nil {
		slog.Error("error getting bear by name", "error", err)
		c.JSON(
			http.StatusInternalServerError,
			error{Error: "Error retrieving data"},
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
// @Param		teddyBear	body		models.TeddyBearInput	true	"Teddy Bear Object"
// @Success		201			{object}	routes.responseID
// @Failure		422			{object}	routes.error
// @Failure		500			{object}	routes.error
// @Failure		409			{object}	routes.error
// @Router		/teddy-bear	[POST]
func (router Router) createTeddyBear(c *gin.Context) {
	// get and validate body as object
	var inputBear models.TeddyBearInput
	if err := c.ShouldBindJSON(&inputBear); err != nil {
		slog.Error("ShouldBindJSON error", "error", err)
		c.JSON(
			http.StatusBadRequest,
			error{Error: "missing values or malformed body"},
		)
		return
	}

	// add to db
	id, err := router.Logic.CreateTeddyBear(inputBear)
	if err != nil {
		slog.Error("error adding teddy bear", "error", err)
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

// @Summary		Delete a teddy bear by name
// @Description	Delete a teddy bear by name
// @Tags		teddy-bear
// @Accept		json
// @Produce		json
// @Param		name				path		string	true	"Teddy Bear Name"
// @Success		202					{object}	routes.responseMessage
// @Failure		500					{object}	routes.error
// @Failure		404					{object}	routes.error
// @Router		/teddy-bear/{name}	[DELETE]
func (router Router) deleteTeddyBearByName(c *gin.Context) {
	// get and validate name
	name := c.Param("name")
	if name == "" {
		slog.Error("error not a valid name")
		c.JSON(
			http.StatusBadRequest,
			error{Error: "Not a valid name"},
		)
		return
	}

	// remove from db
	if err := router.Logic.DeleteTeddyBearByName(name); err != nil {
		slog.Error("error deleting resource", "error", err)
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
