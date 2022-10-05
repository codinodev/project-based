package configs

import (
	"project-based/controllers"

	"github.com/jinzhu/gorm"

)

func DBInit() *gorm.DB{
  db, err := gorm.Open("mysql", "root:dicoding@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
  if err != nil{
    panic("feiled to connect to database")
  }
  db.AutoMigrate(controllers.Person{})
  return db
}
