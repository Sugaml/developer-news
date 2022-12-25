package migrations

import (
	"github.com/Sugaml/developer-news/internal/links"
	"github.com/jinzhu/gorm"
)

func MigarateUp(db *gorm.DB) {
	db.AutoMigrate(
		links.Link{},
	)
}
