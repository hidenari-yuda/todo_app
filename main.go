package main

import (
	"fmt"

	"github.com/hidenari-yuda/todo_app/app/controllers"
	"github.com/hidenari-yuda/todo_app/app/models"
)

func main() {

	fmt.Println(models.Db)

	/*u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "test1234"
	fmt.Println(u)

	u.CreateUser()*/

	/*user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, _ := user.CreateSession()

	fmt.Println(session)

	valid, _ := session.CheckSession()
	fmt.Println(valid)*/

	/*fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)*/
	/*user, _ := models.GetUser(2)

	user.CreateTodo("First Todo")*/

	/*user, _ := models.GetUser(3)

	user.CreateTodo("Third Todo")

	/*todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}*/

	/*user2, _ := models.GetUser(2)
	todos, _ := user2.()
	for _, v := range todos {
		fmt.Println(v)
	}*/

	/*t, _ := models.GetTodo(3)
	t.DeleteTodo()*/

	controllers.StartMainServer()

}
