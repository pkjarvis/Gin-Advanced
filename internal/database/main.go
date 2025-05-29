package internal

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


func InitDB()*gorm.DB{
		dsn:="host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
		db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})

		if err!=nil{
			fmt.Print("Unalbe to connec to the database")
			return nil;
		}
		return db;
}


