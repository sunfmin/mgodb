package mgodb

import (
	"fmt"
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
)

type User struct {
	Email string
	Name  string
}

func (u *User) IdByForeignKeys() string {
	return u.Email
}

func ExampleSave() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{"sunfmin@gmail.com", "Felix Sun"})

	var found *User

	CollectionDo("users", func(uc *mgo.Collection) {
		uc.Find(bson.M{"email": "sunfmin@gmail.com"}).One(&found)
	})
	fmt.Println(found.Name)

	//Output: Felix Sun

}
