package bloom

type Storager interface {
	Init(errRate float64, elements uint) error
	Add(key string, value interface{}) error
	Exist(key string, value interface{}) (bool, error)
}
