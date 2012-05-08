package mgodb

import (
	"launchpad.net/mgo"
	"launchpad.net/mgo/bson"
	"reflect"
)

type Id interface {
	MakeId() interface{}
}

func Save(collectionName string, obj Id) (err error) {
	CollectionDo(collectionName, func(rc *mgo.Collection) {
		_, err = rc.Upsert(bson.M{"_id": obj.MakeId()}, obj)
	})
	return
}

func Delete(collectionName string, id interface{}) (err error) {
	CollectionDo(collectionName, func(rc *mgo.Collection) {
		err = rc.RemoveAll(bson.M{"_id": id})
	})
	return
}

func Update(collectionName string, obj Id) (err error) {
	CollectionDo(collectionName, func(rc *mgo.Collection) {
		v := reflect.ValueOf(obj)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		found := reflect.New(v.Type()).Interface()
		rc.Find(bson.M{"_id": obj.MakeId()}).One(found)

		originalValue := reflect.ValueOf(found)
		if originalValue.Kind() == reflect.Ptr {
			originalValue = originalValue.Elem()
		}

		for i := 0; i < v.NumField(); i++ {
			fieldValue := v.Field(i)
			if !reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
				continue
			}

			fieldValue.Set(originalValue.Field(i))
		}

		rc.Upsert(bson.M{"_id": obj.MakeId()}, v.Interface())
	})
	return
}

func FindAll(collectionName string, query interface{}, result interface{}) (err error) {
	CollectionDo(collectionName, func(c *mgo.Collection) {
		err = c.Find(query).All(result)
	})
	return
}

func FindOne(collectionName string, query interface{}, result interface{}) (err error) {
	CollectionDo(collectionName, func(c *mgo.Collection) {
		err = c.Find(query).One(result)
	})
	return
}

func FindById(collectionName string, id interface{}, result interface{}) (err error) {
	CollectionDo(collectionName, func(c *mgo.Collection) {
		err = c.Find(bson.M{"_id": id}).One(result)
	})
	return
}
