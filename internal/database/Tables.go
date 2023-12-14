package database

import (
	"github.com/airbornharsh/stackoverflow-backend/pkg/models"
)

func TableCreate() {
	if err := DB.AutoMigrate(&models.User{}, &models.PostType{}, &models.Post{}, &models.VoteType{}, &models.Vote{}, &models.Comment{}, &models.Tag{}, &models.PostTag{}); err != nil {
		panic(err)
	}

	// DB.Migrator().DropTable(&models.User{}, &models.Post{}, &models.PostType{}, &models.VoteType{}, &models.Vote{}, &models.Comment{}, &models.Tag{}, &models.PostTag{})
}
