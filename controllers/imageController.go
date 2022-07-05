package controllers

import (
	"fmt"
	"taskmanagerapp/database"
	"taskmanagerapp/models"
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {

	form, err := c.MultipartForm()

	if err != nil {
		return err
	}
	
	files := form.File["image"]
	
	filename := ""

	
	var attachment models.Attachment
	fmt.Println("jön a files")
	fmt.Println(files)
	for _, file := range files {
		filename = file.Filename
		

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}

		attachment.Path ="http://localhost:8000/api/uploads/" + filename

		id := c.Params("id")
    	var task models.Task

    	result := database.DB.Find(&task, id)

		if result.RowsAffected == 0 {
			return c.SendStatus(404)
		}

		fmt.Println("task attachment")
		fmt.Println(&task)
		fmt.Println(&attachment)

		database.DB.Model(&task).Association("Attachments").Append(&attachment)
		
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
