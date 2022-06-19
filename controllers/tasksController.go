package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"taskmanagerapp/database"
	"taskmanagerapp/models"

	"github.com/gofiber/fiber/v2"
)

func AllTasksPaginated(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Task{}, page))
}

func AllTasks(c *fiber.Ctx) error {

	var tasks []models.Task

    database.DB.Find(&tasks)
    return c.Status(200).JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
    id := c.Params("id")
    var task models.Task

    result := database.DB.Find(&task, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&task)
}

func ExportTasks(c *fiber.Ctx) error {
	fmt.Println("lefut az export csv")
	filePath := "./csv/tasks.csv"

	if err := CreateTaskFile(filePath); err != nil {
		fmt.Println("lefut az error csv")
		return err
	}
	fmt.Println("itt vagyok a c.download előtt")
	return c.Download(filePath)
}

func CreateTaskFile(filePath string) error {
	fmt.Println("lefut  a csv task")
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var tasks []models.Task

	database.DB.Preload("TaskItems").Find(&tasks)

	writer.Write([]string{
		"ID", "Name", "Priority", "Description", "Created at", "Updated at", "Due date", "Is in progress",
	})

	fmt.Println("jön a tasks")
	fmt.Println(tasks)

	//This package provides an Itoa() function which is equivalent to FormatInt(int64(x), 10). Or in other words, Itoa() function returns the string representation of x when the base is 10
	for _, task := range tasks {
		data := []string{
			strconv.Itoa(int(task.Id)),
			task.Name ,
			task.DueDate,
			strconv.Itoa(int(task.Priority)),
			task.Description,
			task.CreatedAt,
			task.UpdatedAt,
			task.DueDate,
		}

		if err := writer.Write(data); err != nil {
			return err
		}
		fmt.Println("Benne vagyok az írás ciklusban")

	}

	return nil
}


