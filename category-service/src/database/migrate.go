package database

import (
	"github.com/adiet95/book-store/category-service/src/database/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
)

var migUp bool
var migDown bool

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "start migration",
	RunE:  dbMigrate,
}

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration rollback")

}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "004",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Category{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.Category{})
			},
		},
	})
	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("Migration up done")
		return nil
	}
	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("Migration rollback done")
		return nil
	}
	log.Println("init schema database done")
	return nil
}
