package serve

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
	"github.com/nii236/jmdict/serve/app"
	"github.com/nii236/jmdict/serve/swagger"
)

func main() {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount "Translate" controller
	c := NewTranslateController(service)
	app.MountTranslateController(service, c)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
