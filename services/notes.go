package services

import (
     "gorm.io/gorm"
     "go-tutorial/internal/model"
     "fmt"
)


type NotesService struct{
    db*gorm.DB
}

func (n*NotesService) InitService(database *gorm.DB){
    n.db=database;
    n.db.AutoMigrate(&internal.Notes{})
}

type Note struct{
    Id int
    Name string
}

func (n*NotesService) GetNotesServices(status bool) ([]*internal.Notes,error) {
    var notes []*internal.Notes

    if err:=n.db.Where("status=?",status).Find(&notes).Error;err!=nil{
        return nil,err;
    }
    return notes,nil
}


func (n*NotesService) CreateNotesService(title string,status bool)(*internal.Notes,error) {
    
    note:=&internal.Notes{
        Title:title,
        Status: status,
    }

    if err:=n.db.Create(note).Error;err!=nil{
        fmt.Print(err)
    }

    // data:=Note{
    //     Id:3,Name:"Note 3",
    // }

    // err:=n.db.Create(&internal.Notes{
    //     Id:1,
    //     Title:"Notes",
    //     Status:false,
    // });

    // if err!=nil{
    //     fmt.Print(err)
    // }

    return  note,nil
}


func (n*NotesService) UpdateNotesService(title string,status bool,id int)(*internal.Notes,error) {
    
    var note*internal.Notes

    if err:=n.db.Where("id=?",id).First(&note).Error;err!=nil{
        return nil,err;
    }

    note.Title=title
    note.Status=status

    if err:=n.db.Save(&note).Error;err!=nil{
        fmt.Print(err)
        return nil,err;
    }

    // data:=Note{
    //     Id:3,Name:"Note 3",
    // }

    // err:=n.db.Create(&internal.Notes{
    //     Id:1,
    //     Title:"Notes",
    //     Status:false,
    // });

    // if err!=nil{
    //     fmt.Print(err)
    // }

    return  note,nil;
}


func (n*NotesService) DeleteNotesService(id int64)(error) {
    
    var note*internal.Notes

    if err:=n.db.Where("id=?",id).First(&note).Error;err!=nil{
        return err;
    }

    // note.Title=title
    // note.Status=status

    if err:=n.db.Where("id=?",id).Delete(&note).Error;err!=nil{
        fmt.Print(err);
        return nil;
    }

    

    return  nil;
}

