package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Task struct {
	Id                    uint        `json:"id"`
	Name                  string      `json:"name"`
	ImageUrl              string      `json:"imageUrl"                     gorm:"image_url"`
	Priority              uint32      `json:"priority"`
	Description           string      `json:"description"`
	AttachmentNumber      uint32      `json:"numberOfAttachments"           gorm:"number_of_attachments"`
	DueDate               string      `json:"dueDate"                       gorm:"due_date"`
	UpdatedAt             string      `json:"updatedAt"                     gorm:"updated_at"`
	CreatedAt             string      `json:"createdAt"                     gorm:"created_at"`
	Total                 float32     `json:"total"                         gorm:"-"`
	Attachments           []Attachment  `json:"attachments" gorm:"many2many:related_attachments"`
}

func (task *Task) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(task).Count(&total)

	return total
}

func (task *Task) Take(db *gorm.DB, limit int, offset int) interface{} {
	var tasks []Task

	db.Preload("TaskItems").Offset(offset).Limit(limit).Find(&tasks)

	fmt.Println(tasks)

	for i, _ := range tasks {
		var total float32 = 0

		tasks[i].Name = tasks[i].Name 
		tasks[i].Total = total
	}

	return tasks
}