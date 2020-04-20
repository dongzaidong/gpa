package databasesad

func init() {

}

type MysqlDB struct {
}

func NewMysqlDB() *MysqlDB {
	return nil
}

func (m *MysqlDB) Insert(con interface{}) (bool, error) {
	return false, nil
}

func (m *MysqlDB) Update(con interface{}) (bool, error) {
	return false, nil
}

func (m *MysqlDB) Delete(con interface{}) (bool, error) {
	return false, nil
}

func (m *MysqlDB) Select(con interface{}) (interface{}, error) {
	return nil, nil
}
