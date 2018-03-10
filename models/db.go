package models

import (
	_ "github.com/mattn/go-sqlite3"
	//"github.com/go-xorm/xorm"
	//"github.com/go-pg/pg"
	//"fmt"
	"github.com/go-xorm/xorm"
	//"github.com/go-pg/pg"
)

type User struct {
	Id uint64  `xorm:"'id' pk autoincr"`
	Username string `xorm:"varchar(32) unique index not null"`

	// Memorizations []Memorization `xorm:"foreignkey:UserId"`
}

type Word struct {
	Id uint32  `xorm:"'id' pk autoincr"`
	Word string `xorm:"not null unique index"`

	// Memorizations []Memorization `xorm:"foreignkey:WordId"`
}

type Memorization struct {
	Id uint64 `xorm:"'id' pk autoincr"`
	UserId int64 `xorm:"not null unique('idx__user_id__word_id')"`
	WordId int32 `xorm:"not null unique('idx__user_id__word_id')"`  // index:idx__user_id__word_id
	MemorizationCoefficient float32 `xorm:"not null"`  // default=0.0
	LastUpdateTimestamp int64 `xorm:"not null"`
}

var DB *xorm.Engine;

func InitDb() {
	engine, err := xorm.NewEngine("sqlite3", "./db.sqlite3")

	DB = engine
	PanicIfError(err)
	//if err != nil {
	//	panic("failed to connect database")
	//}
	err = engine.Sync2(new(User))
	PanicIfError(err)

	err = engine.Sync2(new(Word))
	PanicIfError(err)
	err = engine.Sync2(new(Memorization))
	PanicIfError(err)

	// defer DB.Close()
	engine.ShowSQL(true)

	user := User{Username: "admin"}
	ProcessDBResult(DB.Insert(user))
	ProcessDBResult(DB.Insert(User{Username: "abmin"}))
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ProcessDBResult(res interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return res
}
