package services

import (
	"go-tutorial/internal/model"
	"errors"
	"gorm.io/gorm"
)




type AuthService struct{
	db*gorm.DB
}


func InitAuthService(db*gorm.DB)*AuthService{
	db.AutoMigrate(&internal.User{})
	return &AuthService{
		db:db,
	}
}

// Method
func (a*AuthService) Login(email *string,password *string)(*internal.User,error){ // here types are pointer bcz here null values could be possible
	if email==nil{
		return nil,errors.New("email can't be null")
	}

	if password==nil{
		return nil,errors.New("password can't be empty")
	}

	var user internal.User

	if err:=a.db.Where("email=?",email).Where("password=?",password).Find(&user).Error;err!=nil{
		return nil,err;
	}


	return &user,nil;
}


func (a*AuthService) Register(email *string,password *string)(*internal.User,error){ // here types are pointer bcz here null values could be possible
	if email==nil{
		return nil,errors.New("email can't be null")
	}

	if password==nil{
		return nil,errors.New("password can't be empty")
	}

	var user internal.User

	user.Email=*email
	user.Password=*password

	if err:=a.db.Create(&user).Error;err!=nil{
		return nil,err
	}


	return &user,nil;
}