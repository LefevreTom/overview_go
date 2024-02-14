package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/user/project/internal/services"
)

// The contract for the implementation
type GamesServices interface {
	GetGamesByPage(page int) ([]services.Game, error)
}

// Any implementation as paramaters that validate the contract
func NewGamesHandlers(gs GamesServices) *GamesHandler {

	return &GamesHandler{
		GamesServices: gs,
	}
}

type GamesHandler struct {
	GamesServices GamesServices
}

func (gh *GamesHandler) GetGamesByPage(c echo.Context) error {

	// strconv use to transform string to int
	// because everything that come from web is string typed
	page, err := strconv.Atoi(c.QueryParam("page"))

	// if error when converting we return Bad request
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid page")
	}

	// We use the service we created before
	games, err := gh.GamesServices.GetGamesByPage(page)

	// if error appears in services means it is a server error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, games)
}
