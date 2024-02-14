// Filename: main.go
package main

// Filename: main.go
import (
	"flag"
	"os"

	"github.com/user/project/internal/database"
	"github.com/user/project/internal/handlers"
	"github.com/user/project/internal/services"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	PORT := flag.String("port", ":"+os.Getenv("PORT"), "port to run the server on")

	store, err := database.NewStore(os.Getenv("DB_NAME"))

	if err != nil {
		e.Logger.Fatal(err)
	}

	gameServices := services.NewGamesServices(services.Game{}, store, os.Getenv("API_KEY"))
	gameHandler := handlers.NewGamesHandlers(gameServices)

	handlers.SetupRoutes(e, gameHandler)

	// Start the server
	e.Logger.Fatal(e.Start(*PORT))
}
