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
	"github.com/secrethook/backend/pkg/utils"
	"github.com/secrethook/backend/platform/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewChannel(t *testing.T) {
	err := godotenv.Load(repository.TestEnvFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var dataString []string = []string{
		`{"encryption": false}`,
		`{"encryption": true}`,
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
			description:   "Create new channel without any payload",
			route:         "/api/v1/channel/new",
			method:        "POST",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Create new channel with payload",
			route:         "/api/v1/channel/new",
			method:        "POST",
			body:          strings.NewReader(dataString[0]),
			expectedError: false,
			expectedCode:  201,
		},
		{
			description:   "Created new encrypted channel but no key was sent",
			route:         "/api/v1/channel/new",
			method:        "POST",
			body:          strings.NewReader(dataString[1]),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Create new encrypted channel with key",
			route:         "/api/v1/channel/new",
			method:        "POST",
			body:          strings.NewReader(dataString[2]),
			expectedError: false,
			expectedCode:  201,
		},
	}
	database.ConnectToMongodb()
	if err := utils.InitSnowflakeNode(); err != nil {
		panic(err)
	}
	app := fiber.New()
	routes.ChannelRoutes(app.Group("/api/v1/channel"))

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

func TestSendMessage(t *testing.T) {
	err := godotenv.Load(repository.TestEnvFile)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	channelId := "72101158402920448"
	var dataString []string = []string{
		`{"test": true}`,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`,
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
			description:   "Send new message wiht wrong channel",
			route:         "/api/v1/channel/send/" + "123",
			method:        "POST",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Send new message wiht wrong channel and vaild data",
			route:         "/api/v1/channel/send/" + "123",
			method:        "POST",
			body:          strings.NewReader(dataString[0]),
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "Send new message without body",
			route:         "/api/v1/channel/send/" + channelId,
			method:        "POST",
			body:          nil,
			expectedError: false,
			expectedCode:  201,
		},
		{
			description:   "Send new message with paylout",
			route:         "/api/v1/channel/send/" + channelId,
			method:        "POST",
			body:          strings.NewReader(dataString[0]),
			expectedError: false,
			expectedCode:  201,
		},
		{
			description:   "Send new message with paylout",
			route:         "/api/v1/channel/send/" + channelId,
			method:        "POST",
			body:          strings.NewReader(dataString[1]),
			expectedError: false,
			expectedCode:  201,
		},
	}
	database.ConnectToMongodb()
	if err := utils.InitSnowflakeNode(); err != nil {
		panic(err)
	}
	app := fiber.New()
	routes.ChannelRoutes(app.Group("/api/v1/channel"))

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
