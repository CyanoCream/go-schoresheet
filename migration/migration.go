package migration

import (
	"fmt"
	master "go-scoresheet/master/models"
	"go-scoresheet/middleware"
	workflow "go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	fmt.Println("Running Migrations...")

	db.AutoMigrate(
		// Master
		&master.User{},
		&master.UserRole{},
		&master.Role{},
		&master.PermissionRole{},
		&master.Permission{},
		&middleware.Session{},
		// Workflow

		&workflow.Tbl_club{},
		&workflow.Tbl_match{},
		&workflow.Turnament{},
		&workflow.Tbl_player{},
		&workflow.Tbl_score{},
	)

	fmt.Println("Migrations Completed.")
}
