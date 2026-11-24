package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestPostGres(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Logf("no env found")
	}

	dbUrl := os.Getenv("POSTGRES_URL")

	ctx := t.Context()
	db := New(dbUrl)
	err = db.Connect(ctx)
	if err != nil {
		t.Fatalf("failed to connect %v", err)
	}
	t.Log("connection established")
	defer db.Close()
}