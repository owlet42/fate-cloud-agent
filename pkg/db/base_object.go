package db

type BaseObject interface{
	func Save(baseObject BaseObject) (string, error)

	func FindByUUID(uuid string) (*BaseObject, error)

	func Find() (*[]BaseObject, error)

	func UpdateByUUID(uuid string, curd BaseObject) (*BaseObject, error)

	func Delete(uuid string) error
}