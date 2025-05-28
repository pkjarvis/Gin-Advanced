package main

import (
    // "net/http"
    "github.com/gin-gonic/gin"
	"go-tutorial/controllers"
)

func main(){
    r:=gin.Default()

    // Note: Go mod init creates go.mod file which contains the packages 

    // Fetching and Updating Remote Packages The go get command is primarily used for retrieving remote packages from version control repositories and making them available for use in your projects

    // r.GET("/ping",func(c*gin.Context){
    //     c.JSON(http.StatusOK,gin.H{
    //         "message":"pong",
    //         "context":"hello from gin",
    //     })
    // })

    // Fetching params from link
    // r.GET("/me/:id/:newId",func(c*gin.Context){
	// 	var id=c.Param("id")
	// 	var newId=c.Param("newId")
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"user":id,
	// 		"user_newId":newId,
	// 	})
	// })



	// email,password route
	// r.POST("/me",func(c*gin.Context){
	// 	// Email , Password
	// 	type MeRequest struct{
	// 		Email string    `json:"email" binding:"required"`   // backticks are used for binding with frontend
	// 		Password string `json:"password"`
	// 	}

	// 	var merequest MeRequest

	// 	if err:=c.BindJSON(&merequest);err!=nil{
	// 		c.JSON(http.StatusBadRequest,gin.H{
	// 			"error":err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.BindJSON(&merequest)

	// 	c.JSON(http.StatusOK,gin.H{
	// 		"email":merequest.Email,
	// 		"password":merequest.Password,
	// 	})

	// })


    // Getting mail,password from frontend

	// r.PUT("/me",func(c*gin.Context){
	// 	// Email , Password
	// 	type MeRequest struct{
	// 		Email string    `json:"email" binding:"required"`   // backticks are used for binding with frontend
	// 		Password string `json:"password"`
	// 	}

	// 	var merequest MeRequest

	// 	if err:=c.BindJSON(&merequest);err!=nil{
	// 		c.JSON(http.StatusBadRequest,gin.H{
	// 			"error":err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.BindJSON(&merequest)

	// 	c.JSON(http.StatusOK,gin.H{
	// 		"email":merequest.Email,
	// 		"password":merequest.Password,
	// 	})

	// })


	// r.PATCH("/me",func(c*gin.Context){
	// 	// Email , Password
	// 	type MeRequest struct{
	// 		Email string    `json:"email" binding:"required"`   // backticks are used for binding with frontend
	// 		Password string `json:"password"`
	// 	}

	// 	var merequest MeRequest

	// 	if err:=c.BindJSON(&merequest);err!=nil{
	// 		c.JSON(http.StatusBadRequest,gin.H{
	// 			"error":err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.BindJSON(&merequest)

	// 	c.JSON(http.StatusOK,gin.H{
	// 		"email":merequest.Email,
	// 		"password":merequest.Password,
	// 	})

	// })


	
	// r.DELETE("/me/:id",func(c*gin.Context){
	
	// 	var id=c.Param("id")
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"id":id,
	// 		"message":"deleted!!",
	// 	})


	// })


    notesController:=&controllers.NotesController{}
	notesController.InitNotesControllerRoutes(r)

    


    

    r.Run(":8000")


}