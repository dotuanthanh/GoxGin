package mdb

type IRedisHandler interface {
	Set(key string, data *[]interface{}, expTime int) error
	Get(key string) (interface{}, error)
}
