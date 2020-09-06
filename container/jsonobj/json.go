package jsonobj

import (
	"encoding/json"
	"strings"
)

type object struct {
	json.RawMessage
	Split string
}

func NewObject(v []byte) Object {
	return NewObjectWithSplit(v, ".")
}

func NewObjectWithSplit(v []byte, sep string) Object {
	return &object{
		RawMessage: v,
		Split:      sep,
	}
}

// Decode decode object to target
func (o object) Decode(v interface{}) error {
	return json.Unmarshal(o.RawMessage, v)
}

func (o object) get(key string) (json.RawMessage, error) {
	ma := make(map[string]json.RawMessage)
	err := json.Unmarshal(o.RawMessage, &ma)
	if err != nil {
		return nil, err
	}
	return ma[key], nil
}

// Get get from key like "a.b.c.d"
// TODO: improve performance
func (o object) Get(key string) (Object, error) {
	keys := strings.Split(key, o.Split)
	s := o.RawMessage
	for _, k := range keys {
		var err error
		s, err = object{
			RawMessage: s,
			Split:      o.Split,
		}.get(k)
		if err != nil {
			return object{}, err
		}
	}

	return object{
		RawMessage: s,
		Split:      o.Split,
	}, nil
}

func (o object) DecodeByKeys(key string, v interface{}) error {
	s, err := o.Get(key)
	if err != nil {
		return err
	}

	return s.Decode(v)
}

// more type go to DecodeByKeys
func (o object) GetString(key string) (r string) {
	_ = o.DecodeByKeys(key, &r)

	return
}

// more type go to DecodeByKeys
func (o object) GetInt(key string) (r int) {
	_ = o.DecodeByKeys(key, &r)

	return
}

func (o object) String() string {
	return string(o.RawMessage)
}
