package kv

type KEY_TYPE int64
type VAL_TYPE string

type Map interface {
	Get(k KEY_TYPE) (VAL_TYPE, bool)
	Set(k KEY_TYPE, v VAL_TYPE)
}
