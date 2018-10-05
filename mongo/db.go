package mongo

import (
	"github.com/globalsign/mgo"
)

type DB struct {
	session *mgo.Session
	DB      *mgo.Database
}

func (db *DB) Connect(url string, database string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	db.session = session
	db.DB = session.DB(database)
	session.SetMode(mgo.Monotonic, true)
	return nil
}

func (db *DB) Close() {
	if db.session != nil {
		db.session.Close()
	}
}
