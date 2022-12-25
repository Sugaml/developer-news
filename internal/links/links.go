package links

import (
	"fmt"

	database "github.com/Sugaml/developer-news/internal/pkg/db/postgres"
	"github.com/Sugaml/developer-news/internal/users"
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	Title   string      `json:"title"`
	Address string      `json:"address"`
	User    *users.User `json:"user"`
}

func (link Link) Save() uint {
	fmt.Println(link)
	err := database.DB.Model(&Link{}).Save(link).Error
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return link.ID
}
