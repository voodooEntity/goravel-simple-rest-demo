package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	fileController := controllers.NewFileController()
	facades.Route().Get("/api/file/{ident}", fileController.Get)
	facades.Route().Post("/api/file", fileController.Upload)
}
