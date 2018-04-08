package models

import . "config"

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       uint64 `json:"id"`
	Username string `sql:"type:varchar(32),unique,index,notnull" json:"username"`
	Password []byte `sql:",notnull" json:"-"`

	Memorizations []*Memorization `json:"memorizations,omitempty"`
}

type Language struct {
	Id    uint16  `json:"id,omitempty"`
	Code  string  `sql:"type:varchar(32),unique,index,notnull" json:"code"`
	Words []*Word `json:"words,omitempty"`
}

type Word struct {
	Id               uint32    `json:"id,omitempty"`
	Word             string    `sql:",notnull,unique:word__language_id,index" json:"word,omitempty"`
	LanguageId       uint16    `sql:",notnull,unique:word__language_id,index" json:"languageId,omitempty"`
	Language         *Language `json:"language"`
	ApprovedByUserId uint64    `sql:",notnull,default=0" json:"approvedByUserId"`
	ApprovedByUser   *User     `json:"approvedByUser"`
	Frequency        float32   `sql:",notnull,default=0.0" json:"frequency"`
}

type Memorization struct {
	Id                      uint64  `json:"id"`
	UserId                  uint64  `sql:",notnull,unique:user_id__word_id" json:"userId"`
	User                    *User   `json:"user"`
	WordId                  uint32  `sql:",notnull,unique:user_id__word_id" json:"wordId"`
	Word                    *Word   `json:"word"`
	MemorizationCoefficient float64 `sql:",notnull,default:0.0" json:"memorizationCoefficient"`
	LastUpdateTimestamp     int64   `sql:",notnull" json:"lastUpdateTimestamp"`
	NextShowTimestamp       int64   `sql:",notnull,default:0" json:"nextShowTimestamp"`
}

var DB *pg.DB

func InitDb() error {
	dbConfig := Settings.Database
	DB = pg.Connect(&pg.Options{
		Database: dbConfig["Name"],
		User:     dbConfig["Username"],
		Password: dbConfig["Password"],
	})

	DB.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		if Settings.Debug {
			fmt.Printf("SQL: %s %s\n", time.Since(event.StartTime), query)
		}
	})

	err := createSchema(DB)

	if err != nil {
		return err
	}

	//initUsers := []interface{}{
	//	&User{Username: "admin"},
	//	&User{Username: "abmin"}}
	//
	//for _, model := range initUsers {
	//	_, err := DB.Model(model).OnConflict("(username) DO NOTHING").Insert()
	//	if err != nil {
	//		return err
	//	}
	//}

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
		&Memorization{}}

	for _, model := range initTables {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true})
		if err != nil {
			return err
		}
	}
	return nil
}
