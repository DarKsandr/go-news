package main

import (
	"log"
	"main/pkg"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	log.Println("Старт миграции")
	pkg.Init()

	db := pkg.OpenDB()

	tables := []any{
		&pkg.User{},
		&pkg.News{},
	}

	db.Migrator().DropTable(tables...)

	db.AutoMigrate(tables...)

	users := []*pkg.User{}
	news := []*pkg.News{}
	categories := []*pkg.Category{}

	userAvatar := []string{
		"/static/img/features-fashion.webp",
		"/static/img/features-life-style.webp",
		"/static/img/features-sports-1.webp",
		"/static/img/lifestyle-1.webp",
		"/static/img/lifestyle-2.webp",
	}

	newsImage := []string{
		"/static/img/news-1.webp",
		"/static/img/news-2.webp",
		"/static/img/news-3.webp",
		"/static/img/news-4.webp",
		"/static/img/news-5.webp",
		"/static/img/news-6.webp",
		"/static/img/news-7.webp",
		"/static/img/news-8.webp",
	}

	log.Println("Заполняем пользователей")
	for range 100 {
		users = append(users, &pkg.User{
			Name:      randomdata.FullName(randomdata.RandomGender),
			About:     randomdata.Paragraph(),
			AvatarUrl: userAvatar[rand.Intn(len(userAvatar)-1)],
		})
	}
	db.Create(&users)

	log.Println("Заполняем категории")
	categoriesName := []string{
		"Sports",
		"Magazine",
		"Lifestyle",
		"Politician",
		"Technology",
		"Intertainment",
	}
	for _, name := range categoriesName {
		categories = append(categories, &pkg.Category{
			Name: name,
		})
	}
	db.Create(&categories)

	log.Println("Заполняем новости")
	for range 1000 {
		t, err := time.Parse(randomdata.DateOutputLayout, randomdata.FullDate())
		if err != nil {
			panic(err)
		}
		news = append(news, &pkg.News{
			Title:      randomdata.SillyName(),
			Content:    randomdata.Paragraph(),
			ImageUrl:   newsImage[rand.Intn(len(newsImage)-1)],
			ReadTime:   rand.Intn(10) + 1,
			Views:      rand.Intn(1000),
			ShareCount: rand.Intn(100),
			UserID:     users[rand.Intn(len(users)-1)].ID,
			CategoryID: categories[rand.Intn(len(categories)-1)].ID,
			CreatedAt:  t,
			UpdatedAt:  t,
		})
	}
	db.Create(&news)

	log.Println("Конец миграции")
}
