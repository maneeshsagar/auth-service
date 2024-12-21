package migrations

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/maneeshsagar/auth-service/db"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// this functio to run the migration
func RunMigrations() {
	// getting the current working directory
	pwd, _ := os.Getwd()

	// creating the dsn to connect to the DB so that migration can run
	dataSource := db.GetDSN() + "?multiStatements=true"

	// here we may improve this by single time connection
	db, _ := sql.Open("mysql", dataSource)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+pwd+"/migrations/sql",
		"mysql",
		driver,
	)

	if err != nil {
		fmt.Println("RunMigrations: ", err)
		return
	}

	// taking the migration version from the env
	version := cast.ToInt(viper.Get("MIGRATION_VERSION"))
	err = m.Steps(cast.ToInt(version))
	if err != nil {
		fmt.Println("RunMigrations: ", err)
		return
	}
}
