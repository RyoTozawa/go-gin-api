package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	gorm.Model
	Id int `gorm:"primary_key"`
	Name string
	Email string
	CreatedAt time.Time `json:"string"`
	UpdatedAt time.Time `json:"string"`
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
	r.Run()
}

func getUsers() ([]User) {
	db, _:= gorm.Open("postgres", "host=127.0.0.1 user=user dbname=user sslmode=disable password=password")
	defer db.Close()

	usersModel := []User{}
	db.Find(&usersModel)

	return usersModel
}