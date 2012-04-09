package mgodb

import (
	"launchpad.net/mgo"
)

var dialstring, database string

var session *mgo.Session

func Setup(ds string, db string) {
	dialstring = ds
	database = db
}

func getSession() (s *mgo.Session) {
	if session != nil {
		return session
	}

	if database == "" || dialstring == "" {
		panic("mgo: must setup database or dialstring with mgo.Setup(...)")
	}

	var err error
	session, err = mgo.Dial(dialstring)
	if err != nil {
		panic(err)
	}

	return session
}

func CollectionDo(name string, f func(c *mgo.Collection)) {
	s := getSession().Copy()
	defer s.Close()
	f(s.DB(database).C(name))
}

func CollectionsDo(f func(c ...*mgo.Collection), names ...string) {
	s := getSession().Copy()
	defer s.Close()
	cs := []*mgo.Collection{}
	for _, name := range names {
		cs = append(cs, s.DB(database).C(name))
	}
	f(cs...)
}
