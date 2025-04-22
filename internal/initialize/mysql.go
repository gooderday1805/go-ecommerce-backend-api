package initialize

import (
	"fmt"
	"time"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
	"github.com/gooderday1805/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	// check error
	checkErrorPanic(err, "Initialize mysql failed")

	global.Logger.Info("Initialize mysql success")
	global.DB = db

	// set Pool
	SetPool()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.DB.DB()
	if err != nil {
		fmt.Println("mysql err:: %s", err)
	}

	// optimize db
	// kết nối rảnh rỗi sẽ đc giữ lại
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))

	// quanlity max connect
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)

	// time max của 1 kêts nối dù nó rảnh rỗi hay không
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

func migrateTable() {
	err := global.DB.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("Migration DB Error",err)
	}
}
