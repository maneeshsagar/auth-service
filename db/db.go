package db

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// to return the basic DSN from the env
func GetDSN() string {
	dbuser := cast.ToString(viper.Get("DB_USER"))
	dbPassword := cast.ToString(viper.Get("DB_PASSWORD"))
	host := cast.ToString(viper.Get("DB_HOST"))
	port := cast.ToString(viper.Get("DB_PORT"))
	db := cast.ToString(viper.Get("DB_NAME"))
	dataSource := dbuser + ":" + dbPassword + "@tcp(" + host + ":" + port + ")/" + db
	return dataSource
}

func SetUpMySql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := GetDSN() + "?charset=utf8"
	fmt.Println(dataSource)
	orm.RegisterDataBase("default", "mysql", dataSource)
	db_debug := os.Getenv("DB_DEBUG")
	orm.Debug = cast.ToBool(db_debug)
	fmt.Println("Database connected...")
}
