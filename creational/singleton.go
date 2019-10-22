package creational

import "sync"

type Object struct {}

var instance *Object
var once sync.Once

func GetInstance() *Object {
	once.Do(func() {
		instance = &Object{}
	})

	return instance
}