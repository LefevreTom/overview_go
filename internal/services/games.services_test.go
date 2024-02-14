package services_test

import (
	"os"
	"testing"

	. "github.com/user/project/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/user/project/internal/database"
)

func TestGetGamesByPage(t *testing.T) {
	// Arrange
	store, err := database.NewStore("test.db")

	if err != nil {
		t.Fatalf("🔥 failed to connect to the database: %s", err)
	}

	t.Setenv("API_KEY", "yourapikey")

	gameService := NewGamesServices(Game{}, store, os.Getenv("API_KEY"))

	// Act
	result, err := gameService.GetGamesByPage(0)

	if err != nil {
		t.Fatalf("🔥 failed to get games: %s", err)
	}

	// Assert
	assert.NotNil(t, result)
}
