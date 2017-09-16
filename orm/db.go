package orm

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {

	// sqlConnection := fmt.Sprintf("host=%v user=%v port=%v dbname=%v sslmode=disable password=%v", conf.Conf.DB.Host, conf.Conf.DB.User, conf.Conf.DB.Port, conf.Conf.DB.DBName, conf.Conf.DB.Password)
	// db, err := gorm.Open(conf.Conf.DB.Type, sqlConnection)
	db, err := gorm.Open("sqlite3", "notice.db")

	if err != nil {
		return nil, err
	}
	return db, nil
}
