package db

type Database interface {
	Connect()
	Disconnect()
}

type DatabaseImpl struct{}

func New() *DatabaseImpl { return &DatabaseImpl{} }

// TODO
func (d *DatabaseImpl) Connect() {}

// TODO
func (d *DatabaseImpl) Disconnect() {}

// UpdateDefinitions is
func (d *DatabaseImpl) UpdateDefinitions() {}
