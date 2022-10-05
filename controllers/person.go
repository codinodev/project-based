package controllers

import (
	
	"net/http"
	"../controllers/structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson( ctx *gin.Context){
  var(
    person structs.Person
    result gin.H
  )
  id := ctx.Param("id")
  err := idb.DB.Where("id = ?", id).First(&person).Error
  if err != nil{
    result = gin.H{
      "result": err.Error(),
      "count": 0,
    }
  }else{
    result = gin.H{
      "result": person,
      "count": 1,
    }
  }
  ctx.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPerson(ctx *gin.Context){
  var(
    persons []structs.Person
    result gin.H
  )
  idb.DB.Find(&persons)
  if len(persons) <= 0{
    result = gin.H{
      "result": nil,
      "count": 0,
    }
  }else{
    result = gin.H{
      "result": persons,
      "count": len(persons),
    }
  }
  ctx.JSON(http.StatusOK, result)
}
//create new data to the database
func (idb *InDB) CreatePerson(ctx *gin.Context){
  var(
    person structs.Person
    result gin.H
  )
  first_name := ctx.PostForm("first_name")
  last_name := ctx.PostForm("last_name")
  person.First_Name = first_name
  person.Last_Name = last_name
  idb.DB.Create(&person)
  result = gin.H{
  	"result": person,
  }
  ctx.JSON(http.StatusOK, result)
}

//update data with {id} as query
func (idb *InDB) UpdatePerson(ctx *gin.Context){
	id := ctx.Query("id")
	first_name := ctx.PostForm("first_name")
	last_name := ctx.PostForm("last_name")
	var (
		person structs.Person
		newPerson structs.Person
		result gin.H
	)
	err := idb.DB.First(&person, id).Error
	if err != nil{
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person). Updates(newPerson).Error
	if err != nil{
		result = gin.H{
			"result": "update failed",
		}
	}else{
		result = gin.H{
			"result": "successfully update data",
		}
	}
	ctx.JSON(http.StatusOK, result)
}

//delete data with {id}
func (idb *InDB) DeletePerson(ctx *gin.Context){
	var(
		person structs.Person
		result gin.H
	)
	id := ctx.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil{
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil{
		result = gin.H{
			"result": "delete failed",
		}
	}else{
		result = gin.H{
			"result": "Data delete successfully",
		}
	}
	ctx.JSON(http.StatusOK, result)
}
