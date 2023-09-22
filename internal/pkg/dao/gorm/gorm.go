package gorm

import (
	"fmt"
	"os"
	"strconv"

	"eirc.app/internal/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver" // 引入 SQL Server 驅動程序
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// func New() (*gorm.DB, error) {
// 	dsn := "sqlserver://ManuOrder:nkust215@163.18.42.223:1433?database=XMLY5000"
// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func New() (*gorm.DB, error) {
	const config = "sqlserver://%s:%s@%s:%s?database=%s"
	sources := fmt.Sprintf(config,
		os.Getenv("SOURCES_USER"),
		os.Getenv("SOURCES_PASSWORD"),
		os.Getenv("SOURCES_HOST"),
		os.Getenv("SOURCES_PORT"),
		os.Getenv("SOURCES_DATABASE"),
	)

	db, err := gorm.Open(sqlserver.Open(sources), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, err
	}

	dbReplicas, err := strconv.Atoi(os.Getenv("DB_REPLICAS"))
	if err != nil {
		log.Debug(err)
	}
	if dbReplicas >= 1 {
		var dialectics []gorm.Dialector
		for i := 1; i <= dbReplicas; i++ {
			replicas := fmt.Sprintf(config,
				os.Getenv("REPLICAS_USER_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_PASSWORD_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_HOST_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_PORT_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_DATABASE_"+strconv.Itoa(i)),
			)
			director := sqlserver.Open(replicas) // 使用 SQL Server 驅動程序
			dialectics = append(dialectics, director)
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{Replicas: dialectics}))
		if err != nil {
			log.Error(err)
		}
	}

	return db, nil
}
func NewSQLite() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
