package main

import (
	"MTUCI-VvIT-labs/lab-4/internal/handlers"
	"MTUCI-VvIT-labs/lab-4/pkg/pg"
	"github.com/gin-gonic/gin"
	"log"
)

const PostgresUrl = "postgres://postgres:123456@localhost/service_db?sslmode=disable" // адресс бд

func main() {
	err := pg.ConnectToDB(PostgresUrl) // подключение к бд
	if err != nil {                    // обработка ошибок при подключении к бд
		log.Fatalln(err)
	}

	r := gin.Default()                                                   // создание роутера
	r.LoadHTMLGlob("/root/go/src/MTUCI-VvIT-labs/lab-4/web/templates/*") // загрузка шаблонов

	r.GET("/login", handlers.LoginPage)       // обработка запроса на страницу логина
	r.POST("/login/", handlers.Authorization) // обработка запроса на авторизацию

	err = r.Run()   // запуск сервера
	if err != nil { // обработка ошибок при запуске сервера
		log.Fatal(err)
	}
}
