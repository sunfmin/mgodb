package mgodb

import (
	"fmt"
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
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

func (u *User) MakeId() string {
	u.Id = u.Email
	return u.Id
}

func ExampleSave() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User

	CollectionDo("users", func(uc *mgo.Collection) {
		uc.Find(bson.M{"email": "sunfmin@gmail.com"}).One(&found)
	})
	fmt.Println(found.Name)

	//Output: Felix Sun

}

func ExampleFind() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	var found *User

	FindOne("users", "sunfmin@gmail.com", &found)
	fmt.Println(found.Name)

	//Output: Felix Sun

}

func ExampleUpdate() {
	Setup("localhost", "mgodb_test")
	Save("users", &User{Email: "sunfmin@gmail.com", Name: "Felix Sun"})

	Update("users", &User{Email: "sunfmin@gmail.com", Gender: "Male"})

	var found *User
	FindOne("users", "sunfmin@gmail.com", &found)
	fmt.Println(found.Name)

	u := &User{Email: "sunfmin@gmail.com", Company: &Company{Name: "The Plant"}}
	Update("users", u)
	FindOne("users", "sunfmin@gmail.com", &found)
	fmt.Println(found.Company.Name)
	fmt.Println(u.Gender)
	//Output:
	//Felix Sun
	//The Plant
	//Male
}
