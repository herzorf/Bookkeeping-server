package database

import (
	"bookkeeping-server/unit"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ValidationCode struct {
	ID        uint       `gorm:"primaryKey"`
	Code      string     `gorm:"size:20;not null"`
	Email     string     `gorm:"size:255;not null"`
	UsedAt    *time.Time // 可以为空，代表未使用
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

var DB *gorm.DB

func Connect() {
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	db, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	unit.HandleError("数据库连接失败", err)
	DB = db
	// 如果数据库不存在就创建数据库
	DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " CHARSET utf8mb4 COLLATE utf8mb4_general_ci;")
	DB.Exec(fmt.Sprintf("USE %s", dbname))
}

func Migrate() {
	DB.AutoMigrate(&ValidationCode{})
}
