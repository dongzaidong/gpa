package database


type Executer struct {
	Insert(interface{}) (bool,error)

	Update(interface{}) (bool,error)

	Delete(interface{}) (bool,error)

	Select(interface{}) (interface{},error)
}


func CreateExecuter(dbType string) Executer{
	var db Executer
	switch dbType{
	case "mysql":
		db = NewMysqlDB()
	case "mongo":
		db = NewMongoDB()
	default:
		panic("No such data storage")
	}
	return db
}