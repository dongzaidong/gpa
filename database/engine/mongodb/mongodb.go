package database

type MongoDB struct {
}

func NewMongoDB() *MongoDB {
	return nil
}

func (m *MongoDB) Insert(con interface{}) (bool, error) {
	return false, nil
}

func (m *MongoDB) Update(con interface{}) (bool, error) {
	return false, nil
}

func (m *MongoDB) Delete(con interface{}) (bool, error) {
	return false, nil
}

func (m *MongoDB) Select(con interface{}) (interface{}, error) {
	return nil, nil
}
