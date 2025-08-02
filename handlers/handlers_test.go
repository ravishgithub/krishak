package handlers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	// Set JWT secret for token generation
	os.Setenv("JWT_SECRET", "ebee1a4380a9ab9a0a84b091c1f7abcf30c3428608f122dbd91e13db134b16bc")

	// Set CONFIG_PATH as an absolute path
	cwd, _ := os.Getwd()
	configPath := filepath.Join(cwd, "../configs/config.json")
	os.Setenv("CONFIG_PATH", configPath)

	// Optional debug
	println("CONFIG_PATH set to:", configPath)

	// Run the tests
	os.Exit(m.Run())
}
