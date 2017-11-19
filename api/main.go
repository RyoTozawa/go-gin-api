package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"time"

)

type User struct {
	Id int `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

type UserForm struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		users := getUsers()
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		user := getUser(c.Param("id"))
		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var user UserForm
		c.Bind(&user)
		postedUser := createUser(user)
		c.JSON(200, postedUser)
	})
	r.Run()
}

func getUsers() ([]User) {
	db, _:= gorm.Open("postgres", "host=127.0.0.1 user=user dbname=user sslmode=disable password=password")
	defer db.Close()

	usersModel := []User{}
	db.Find(&usersModel)

	return usersModel
}

func getUser(id string) (User) {
	db, _:= gorm.Open("postgres", "host=127.0.0.1 user=user dbname=user sslmode=disable password=password")
	defer db.Close()
	userModel := User{}
	db.First(&userModel, id)

	return userModel
}

func createUser(userInfo UserForm) (User) {
	db, _:= gorm.Open("postgres", "host=127.0.0.1 user=user dbname=user sslmode=disable password=password")
	defer db.Close()
	userModel := User{}
	userModel.Name = userInfo.Name
	userModel.Email = userInfo.Email
	db.Create(&userModel)
	return userModel
}