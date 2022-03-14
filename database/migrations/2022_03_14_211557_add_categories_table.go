package migrations

import (
    "database/sql"
    "gohub/app/models"
    "gohub/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type User struct {
        models.BaseModel

        Name        string `gorm:"type:varchar(255);not null;index"`
        Description string `gorm:"type:varchar(255);default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&User{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&User{})
    }

    migrate.Add("2022_03_14_211557_add_categories_table", up, down)
}