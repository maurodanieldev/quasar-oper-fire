package apm

import (
	"sync"

	i "github.com/maurodanieldev/quasar-oper-fire/interfaces"
)

var (
	instance i.APM
	once     sync.Once
)

func Get() i.APM {
	once.Do(func() {
		instance = getInstance()
	})

	return instance
}

func getInstance() i.APM {
	return createEmptyAPM()
}
