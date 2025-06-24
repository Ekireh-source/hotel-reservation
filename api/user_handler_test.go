package api

import (
	"bytes"
	"context"
	"encoding/json"

	
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Ekireh-source/hotel-reservation/db"
	"github.com/Ekireh-source/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ( 
	testdburi = "mongodb://localhost:27017"
	dbname = "hotel-reservation-test"
	)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}


func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "doe@gmail.ccom",
		Password:  "password123",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error("Error making request:", err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	if len(user.ID) == 0 {
		t.Error("Expected user ID to be set, got empty ID")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("Expected encrypted password to be empty, got:", user.EncryptedPassword)
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected first name %s, got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected last name %s, got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected email %s, got %s", params.Email, user.Email)
	}
}