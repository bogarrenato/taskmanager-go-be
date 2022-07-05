package routes

import (
	"taskmanagerapp/controllers"
	//"taskmanagerapp/middlewares"
	//"taskmanagerapp/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//app.Use(middlewares.IsAuthenticated)

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)

	app.Get("/api/products", controllers.AllProducts)
	//EZ A MÁSIK AMI LEFUT
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	//LEFUT file feltöltéskor
	app.Post("/api/upload/:id", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/orders", controllers.AllOrders)
	app.Post("/api/export", controllers.Export)
	app.Get("/api/chart", controllers.Chart)

	//Renato part
	app.Get("/api/tasks/paginated",controllers.AllTasksPaginated)
	app.Get("/api/tasks",controllers.AllTasks)
	app.Get("/api/task/:id", controllers.GetTask)
	app.Post("/api/tasks/add", controllers.CreateTask)

	app.Post("/api/tasks/export", controllers.ExportTasks)
	app.Post("/api/attachment/:id", controllers.DownloadAttachment)
}
