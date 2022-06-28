package controllers

import (
	"fmt"
	"taskmanagerapp/database"
	"taskmanagerapp/models"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	fmt.Println(c)

	form, err := c.MultipartForm()
	fmt.Println("Bejött a kérés")

	if err != nil {
		fmt.Println("error")
		return err
	}
	
	files := form.File["image"]
	
	filename := ""

	fmt.Println("jön a files")
	fmt.Println(files)
	for _, file := range files {
		fmt.Println("Benne vagyok a forban")
		filename = file.Filename
		//taskid := form.File["fileid"]
		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			fmt.Println("Megy a mentés")
			return err
		}
	}

	var task models.Attachment

	task.Path ="http://localhost:8000/api/uploads/" + filename

	database.DB.Create(&task)


	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
