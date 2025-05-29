package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	// "net/http"
	"go-tutorial/services"
)


type NotesController struct {
    notesService services.NotesService
}

func (n*NotesController) InitController(notesService services.NotesService)*NotesController{
    
    n.notesService=notesService;
    return n;
    
}

func (n*NotesController) InitRoutes(router*gin.Engine){
    notes:=router.Group("/notes")
    notes.GET("/",n.GetNotes())
    notes.POST("/",n.CreateNotes())
    notes.PUT("/",n.UpdateNotes())
    notes.DELETE("/:id",n.DeleteNotes())
    
}


func (n*NotesController) GetNotes() gin.HandlerFunc{
    return func(c*gin.Context){

        status:=c.Query("status")

        if status==""{
            status="false"
        }
        // since query returns string but here it is bool so we need to parse it
        actualStatus,err:=strconv.ParseBool(status)  // here strconv function represents conversion from string

        if err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }

        notes,err:=n.notesService.GetNotesServices(actualStatus);

        if err!=nil{
            c.JSON(400,gin.H{
            "message":err.Error(),
            })
            return
        }
         c.JSON(200,gin.H{
            "notes":notes,
        })
        
    }
}

func (n*NotesController) CreateNotes() gin.HandlerFunc{
    type NoteBody struct{
        Title string `json:"title" binding:"required"`
        Status bool `json:"status"`
    }


    return func(c*gin.Context){
        var noteBody NoteBody
        if err:=c.BindJSON(&noteBody);err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }
        // n.notesService.CreateNotesService()

        note,err:=n.notesService.CreateNotesService(noteBody.Title,noteBody.Status)
        if err!=nil{
            c.JSON(400,gin.H{
                "message":err,
            })
            return
        }
        
        c.JSON(200,gin.H{
            "notes":note,
        })
    }
}

func (n*NotesController) UpdateNotes() gin.HandlerFunc{
    type NoteBody struct{
        Title string `json:"title" binding:"required"`
        Status bool `json:"status"`
        Id int `json:"id" binding:"required"`
    }


    return func(c*gin.Context){
        var noteBody NoteBody
        if err:=c.BindJSON(&noteBody);err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }
        // n.notesService.CreateNotesService()

        note,err:=n.notesService.UpdateNotesService(noteBody.Title,noteBody.Status,noteBody.Id)
        if err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }
        
        c.JSON(200,gin.H{
            "notes":note,
        })
    }
}


func (n*NotesController) DeleteNotes() gin.HandlerFunc{
    type NoteBody struct{
        Title string `json:"title" binding:"required"`
        Status bool `json:"status"`
        Id int `json:"id" binding:"required"`
    }


    return func(c*gin.Context){
        // var noteBody NoteBody
        // if err:=c.BindJSON(&noteBody);err!=nil{
        //     c.JSON(400,gin.H{
        //         "message":err.Error(),
        //     })
        //     return
        // }
        // n.notesService.CreateNotesService()

        id:=c.Param("id")
        noteId,err:=strconv.ParseInt(id,10,64)   
        
        if err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }

        err=n.notesService.DeleteNotesService(noteId)
        if err!=nil{
            c.JSON(400,gin.H{
                "message":err.Error(),
            })
            return
        }
        
        c.JSON(200,gin.H{
            "notes":"Note Deleted!",
        })
    }
}