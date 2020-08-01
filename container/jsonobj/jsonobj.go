package jsonobj

type Object interface {
	Decode(v interface{}) error
	Get(key string) (Object, error)
	DecodeByKeys(key string, v interface{}) error
	GetString(key string) (r string)
	GetInt(key string) (r int)
	String() string
}
