package database


type Database struct {
	db map[string] interface{}
}

func NewDatabase() *Database {
	return &Database{}
}


