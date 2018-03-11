package models

import (
	"../settings"
	_ "github.com/mattn/go-sqlite3"
	//"github.com/go-xorm/xorm"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	//"fmt"
	//"github.com/go-xorm/xorm"
	//"github.com/go-pg/pg"
	//"fmt"
	//"reflect"
)

type User struct {
	Id uint64
	Username string `sql:"type:varchar(32),unique,index,notnull"`
	Memorizations []*Memorization
}

type Language struct {
	Id uint16
	Code string `sql:"type:varchar(32),unique,index,notnull"`
	Words []*Word
}

type Word struct {
	Id uint32
	Word string `sql:",notnull,unique,index"`
	LanguageId uint16 `sql:",notnull,index"`
	Language *Language
}

type Memorization struct {
	Id uint64
	UserId uint64 `sql:",notnull,unique:user_id__word_id"`
	User *User
	WordId uint32 `sql:",notnull,unique:user_id__word_id"`
	Word *Word
	MemorizationCoefficient float32 `sql:",notnull,default:0.0"`
	LastUpdateTimestamp uint64 `sql:",notnull"`
}

var DB *pg.DB;

func InitDb() error {
	DB = pg.Connect(&pg.Options{
		Database: settings.DB_NAME,
		User: settings.DB_USERNAME,
		Password: settings.DB_PASSWORD,
	})

	err := createSchema(DB)

	if err != nil {
		return err
	}

	initUsers := []interface{}{
		&User{Username: "admin"},
		&User{Username: "abmin"}}

	for _, model := range initUsers {
		_, err := DB.Model(model).OnConflict("(username) DO NOTHING").Insert()
		if err != nil {
			return err
		}
	}

	initLanguages := []interface{}{
		&Language{Code: "eng"},
		&Language{Code: "rus"}}

	for _, model := range initLanguages {
		_, err := DB.Model(model).OnConflict("(code) DO NOTHING").Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

func createSchema(db *pg.DB) error {
	initTables := []interface{}{
		&User{},
		&Language{},
		&Word{},
		&Memorization{},}

	for _, model := range initTables {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			FKConstraints: true,})
		if err != nil {
			return err
		}
	}
	return nil
}
