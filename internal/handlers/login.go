package handlers

import (
	"MTUCI-VvIT-labs/lab-4/internal/entities"
	"MTUCI-VvIT-labs/lab-4/pkg/pg"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

// LoginPage отдаем готовую html-страницу с формой авторизации
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Authorization(c *gin.Context) {
	var user entities.User             // создаем структуру для хранения данных пользователя
	login := c.PostForm("username")    // получаем логин из формы
	password := c.PostForm("password") // получаем пароль из формы

	// проверяем на пустоту логин и папароль, если пустые -> возвращаем страницу с сообщением об ошибке
	if login == "" || password == "" {
		c.HTML(http.StatusOK, "account.html", gin.H{"content": "Login or password is empty"})
		return
	}

	err := getUser(&user, login, password) // вызываем функцию для получения данных пользователя из БД
	if err == pgx.ErrNoRows {              // если пользователь не найден -> возвращаем страницу с сообщением об ошибке
		c.HTML(http.StatusOK, "account.html", gin.H{"content": "User not found"})
		return
	} else if err != nil { // обрабатываем остальные ошибки
		log.Println(pgx.ErrNoRows)
		log.Fatal(err)
	}

	content := makeContent(&user) // вызываем функцию для формирования строки с данными пользователя

	c.HTML(http.StatusOK, "account.html", gin.H{"content": content}) // возвращаем страницу с данными пользователя
}

// getUser получаем данные пользователя из БД
func getUser(user *entities.User, login string, password string) (err error) {
	row := pg.DB.QueryRow(context.Background(), "SELECT * FROM service.users WHERE login=$1 AND password=$2", login, password)
	return row.Scan(&user.Id, &user.FullName, &user.Login, &user.Password)

}

// makeContent формируем строку с данными пользователя
func makeContent(user *entities.User) string {
	return "Hello, " + user.FullName + "! Your login: " + user.Login + ". Your password: " + user.Password + "."
}
