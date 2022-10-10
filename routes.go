package main

import "github.com/gofiber/fiber/v2"

func setupApp(app *fiber.App) {
	app.Post("/image/upload", uploadImage)
	app.Static("/image/images", "./images")
	app.Delete("/image/:imageName", deleteImage)
}
