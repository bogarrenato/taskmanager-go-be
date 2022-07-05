package controllers

import (
	"fmt"
	"taskmanagerapp/database"
	"taskmanagerapp/models"

	"github.com/gofiber/fiber/v2"
)

func DownloadAttachment(c *fiber.Ctx) error {

	fmt.Println("lefutott a get")

	id := c.Params("id")
	id = "54"
    var attachment models.Attachment

    result := database.DB.Find(&attachment, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

	fmt.Println("lefutott a get")
	fmt.Println(attachment.Path)
	fmt.Println(&attachment)
	

   
	return c.Download("./uploads/pic.png")
}
