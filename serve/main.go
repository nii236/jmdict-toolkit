package serve

import (
	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"

	goalogrus "github.com/goadesign/logging/logrus"
	"github.com/goadesign/middleware"
	"github.com/nii236/jmdict-toolkit/serve/app"
	"github.com/nii236/jmdict-toolkit/serve/swagger"
)

func main() {
	// Create service
	service := goa.New("API")
	logger := logrus.New()
	goa.Log = goalogrus.New(logger)

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

	// Mount "Word" controller
	c := NewWordController(service)
	app.MountWordController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
