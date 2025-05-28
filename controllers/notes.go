package controllers

import (
    "github.com/gin-gonic/gin"
    // "net/http"
    "go-tutorial/services"
)


type NotesController struct {
    notesService services.NotesService
}

func (n*NotesController) InitNotesControllerRoutes(router*gin.Engine){
    notes:=router.Group("/notes")
    notes.GET("/",n.GetNotes())
    notes.POST("/",n.CreateNotes())
}


func (n*NotesController) GetNotes() gin.HandlerFunc{
    return func(c*gin.Context){
        c.JSON(200,gin.H{
            "notes": n.notesService.GetNotesServices(),
        })
    }
}

func (n*NotesController) CreateNotes() gin.HandlerFunc{
    return func(c*gin.Context){
        c.JSON(200,gin.H{
            "notes":n.notesService.CreateNotesService(),
        })
    }
}