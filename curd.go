package mgodb

type Id interface {
	MakeId() interface{}
}

func Save(collectionName string, obj Id) (err error) {
	err = DefaultDatabase.Save(collectionName, obj)
	return
}

func DropCollection(collectionName string) (err error) {
	err = DefaultDatabase.DropCollection(collectionName)
	return
}

func DropCollections(collectionNames ...string) (err error) {
	err = DefaultDatabase.DropCollections(collectionNames...)
	return
}

func Delete(collectionName string, id interface{}) (err error) {
	err = DefaultDatabase.Delete(collectionName, id)
	return
}

func Update(collectionName string, obj Id) (err error) {
	err = DefaultDatabase.Update(collectionName, obj)
	return
}

func FindAll(collectionName string, query interface{}, result interface{}) (err error) {
	err = DefaultDatabase.FindAll(collectionName, query, result)
	return
}

func FindOne(collectionName string, query interface{}, result interface{}) (err error) {
	err = DefaultDatabase.FindOne(collectionName, query, result)
	return
}

func FindById(collectionName string, id interface{}, result interface{}) (err error) {
	err = DefaultDatabase.FindById(collectionName, id, result)
	return
}
