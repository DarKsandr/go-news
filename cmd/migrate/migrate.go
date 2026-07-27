package main

import (
	"main/pkg"
	"math/rand"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	pkg.Init()

	db := pkg.OpenDB()

	tables := []any{
		&pkg.User{},
		&pkg.News{},
	}

	db.Migrator().DropTable(tables...)

	db.AutoMigrate(tables...)

	image := "https://upload.wikimedia.org/wikipedia/commons/1/11/Test-Logo.svg"

	for range 100 {
		user := pkg.User{
			Name:      randomdata.FullName(randomdata.RandomGender),
			About:     randomdata.Paragraph(),
			AvatarUrl: image,
		}
		db.Create(&user)

		news_count := rand.Intn(5)
		news := []*pkg.News{}
		for range news_count {
			news = append(news, &pkg.News{
				Title:      randomdata.SillyName(),
				Content:    randomdata.Paragraph(),
				ImageUrl:   image,
				ReadTime:   rand.Intn(10) + 1,
				Views:      rand.Intn(1000),
				ShareCount: rand.Intn(100),
				UserID:     user.ID,
			})
		}
		if len(news) > 0 {
			db.Create(&news)
		}
	}
}
