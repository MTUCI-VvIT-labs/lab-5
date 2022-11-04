package handlers

import (
	"MTUCI-VvIT-labs/lab-4/pkg/pg"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

// RegistrationPage обрабатывает запрос на страницу регистрации и отдает готовую-html страницу
func RegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", gin.H{})
}

func Registration(c *gin.Context) {
	fullName, login, password := parseForm(c)            // вызов функции парсинга формы
	if fullName == "" || login == "" || password == "" { // проверка на пустые поля
		c.HTML(http.StatusOK, "account.html", gin.H{"content": "One of the fields is empty"})

	} else if isUserExist(login) { // проверка на существование пользователя с таким логином
		c.HTML(http.StatusOK, "account.html", gin.H{"content": "User with this login already exist"})

	} else { // если все поля заполнены и пользователь с таким логином не существует, то добавляем пользователя в бд
		_, err := pg.DB.Exec(context.Background(), "INSERT INTO service.users (full_name, login, password) VALUES ($1, $2, $3)", fullName, login, password)
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "account.html", gin.H{"content": "Registration was successful"}) // отдаем страницу с сообщением об успешной регистрации
	}
}

// ParseForm парсит html-форму и возвращает значения полей
func parseForm(c *gin.Context) (string, string, string) {
	fullName := c.PostForm("full_name")
	login := c.PostForm("username")
	password := c.PostForm("password")
	return fullName, login, password
}

// isUserExist проверяет существование пользователя с таким логином
func isUserExist(login string) (isExist bool) {
	row := pg.DB.QueryRow(context.Background(), "SELECT * FROM service.users WHERE login=$1", login)
	err := row.Scan()
	if err == pgx.ErrNoRows {
		return false
	} else if err == nil {
		return true
	} else {
		panic(err)
	}

}
