package handlers

import (
	"backend/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// Connect ke Supabase
	config.ConnectDB()

	code := m.Run()

	os.Exit(code)
}
