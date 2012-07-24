package mgodb

import (
	"labix.org/v2/mgo"
)

var DefaultDatabase *Database

func Setup(ds string, db string) {
	DefaultDatabase = NewDatabase(ds, db)
}

func CollectionDo(name string, f func(c *mgo.Collection)) {
	DefaultDatabase.CollectionDo(name, f)
}

func DatabaseDo(f func(db *mgo.Database)) {
	DefaultDatabase.DatabaseDo(f)
}

func CollectionsDo(f func(c ...*mgo.Collection), names ...string) {
	DefaultDatabase.CollectionsDo(f, names...)
}
