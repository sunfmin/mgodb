package mgodb

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"testing"
)

const (
	ALLUSERS = "allusers"
	DB1      = "mogdb_test_1"
	DB2      = "mogdb_test_2"
)

func TestSave(t *testing.T) {
	db := NewDatabase("localhost", DB1)

	db.Save(ALLUSERS, &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User
	db.CollectionDo(ALLUSERS, func(uc *mgo.Collection) {
		uc.Find(bson.M{"email": "sunfmin@gmail.com"}).One(&found)
	})
	if found == nil {
		t.Error("Can not find user after saved")
	}

	db = NewDatabase("localhost", DB2)

	db.Save(ALLUSERS, &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var u2 *User
	db.CollectionDo(ALLUSERS, func(uc *mgo.Collection) {
		uc.Find(bson.M{"email": "sunfmin@gmail.com"}).One(&u2)
	})
	if u2 == nil {
		t.Error("Can not find user after saved")
	}

}
