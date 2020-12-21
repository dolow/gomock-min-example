package dao

type User interface {
	Create() error
}

type UserImpl struct {
	tableName string
}

func (*UserImpl) Create() error {
	return nil
}
