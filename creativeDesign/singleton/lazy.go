package singleton

import "sync"

//懒汉模式

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazySingleton() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}

	return lazySingleton
}
