package json

import (
	"encoding/json"
	"strings"
)

type Object struct {
	json.RawMessage
	Split string
}

func NewObject(v []byte) *Object {
	return &Object{
		RawMessage: v,
		Split:      ".",
	}
}

// Decode decode object to target
func (o Object) Decode(v interface{}) error {
	return json.Unmarshal(o.RawMessage, v)
}

func (o Object) get(key string) (json.RawMessage, error) {
	ma := make(map[string]json.RawMessage)
	err := json.Unmarshal(o.RawMessage, &ma)
	if err != nil {
		return nil, err
	}
	return ma[key], nil
}

// Get get from key like "a.b.c.d"
// TODO: improve performance
func (o Object) Get(key string) (Object, error) {
	keys := strings.Split(key, o.Split)
	s := o.RawMessage
	for _, k := range keys {
		var err error
		s, err = NewObject(s).get(k)
		if err != nil {
			return Object{}, err
		}
	}

	return Object{
		RawMessage: s,
		Split:      o.Split,
	}, nil
}

func (o Object) DecodeByKeys(key string, v interface{}) error {
	s, err := o.Get(key)
	if err != nil {
		return err
	}

	return s.Decode(v)
}

// more type go to DecodeByKeys
func (o Object) GetString(key string) (r string) {
	_ = o.DecodeByKeys(key, &r)

	return
}

// more type go to DecodeByKeys
func (o Object) GetInt(key string) (r int) {
	_ = o.DecodeByKeys(key, &r)

	return
}

func (o Object) String() string {
	return string(o.RawMessage)
}
