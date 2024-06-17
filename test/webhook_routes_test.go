package routes

import (
	"io"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/secrethook/backend/pkg/repository"
	"github.com/secrethook/backend/pkg/routes"
	"github.com/secrethook/backend/platform/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewWebhook(t *testing.T) {
    err := godotenv.Load(repository.TestEnvFile)
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	var dataString []string = []string{
		`{"encryption": false}`,
		`{"encryption": true,}`,
		`{"encryption": true, "publicKey": "test",  "privateKey": "test"}`,
	}

	tests := []struct {
		description   string
		route         string
		method        string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "Create new webhook without any payload",
			route:         "/api/v1/webhook/new",
			method:        "POST",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Create new webhook with payload",
			route:         "/api/v1/webhook/new",
			method:        "POST",
			body:          strings.NewReader(dataString[0]),
			expectedError: false,
			expectedCode:  201,
		},
		{
			description:   "Created new encrypted webhook but no key was sent",
			route:         "/api/v1/webhook/new",
			method:        "POST",
			body:          strings.NewReader(dataString[1]),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Create new encrypted webhook with key",
			route:         "/api/v1/webhook/new",
			method:        "POST",
			body:          strings.NewReader(dataString[2]),
			expectedError: false,
			expectedCode:  201,
		},
	}
	database.ConnectToMongodb()
	app := fiber.New()
	routes.WebhookRoutes(app.Group("/api/v1/webhook"))

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
