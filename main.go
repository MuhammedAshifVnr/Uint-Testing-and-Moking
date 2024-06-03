package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Name string
	Id   int
}

var users = map[int]Users{
	1: {Name: "Ashif", Id: 1},
	2: {Name: "Nuhman", Id: 2},
	3: {Name: "Abdin", Id: 3},
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(400,"invalid user id")
		return
	}
	user, exicit := users[id]
	if !exicit {
		c.JSON(404, "user not found")
		return
	}
	c.JSON(200, user)
}

func main() {
	r := gin.Default()
	r.GET("/user/:id", GetUser)
	r.Run(":9090")
}
