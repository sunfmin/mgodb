package mgodb

import (
	"fmt"
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
)

const (
	USERS = "users"
)

type User struct {
	Id      string `bson:"_id"`
	Email   string
	Name    string
	Gender  string
	Company *Company
}

type Company struct {
	Name string
}

func (u *User) MakeId() interface{} {
	u.Id = u.Email
	return u.Id
}

func ExampleSave() {
	Setup("localhost", "mgodb_test")
	Save(USERS, &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User
	CollectionDo("users", func(uc *mgo.Collection) {
		uc.Find(bson.M{"email": "sunfmin@gmail.com"}).One(&found)
	})
	fmt.Println(found.Name)

	//Output: Felix Sun

}

func ExampleFindOne() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User
	FindOne(USERS, bson.M{"email": "sunfmin@gmail.com"}, &found)
	fmt.Println(found.Name)

	//Output: Felix Sun
}

func ExampleFindById() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User
	FindById(USERS, "sunfmin@gmail.com", &found)
	fmt.Println(found.Email)

	// Output: sunfmin@gmail.com
}

func ExampleUpdate() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	Update("users", &User{Email: "sunfmin@gmail.com", Gender: "Male"})

	var found *User
	FindById("users", "sunfmin@gmail.com", &found)
	fmt.Println(found.Name)

	u := &User{Email: "sunfmin@gmail.com", Company: &Company{Name: "The Plant"}}
	Update("users", u)
	FindById("users", "sunfmin@gmail.com", &found)
	fmt.Println(found.Company.Name)
	fmt.Println(u.Gender)
	//Output:
	//Felix Sun
	//The Plant
	//Male
}
