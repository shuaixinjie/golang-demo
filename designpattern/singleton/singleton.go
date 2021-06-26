package singleton

import "sync"

// 饿汉模式就没什么就好写的了，我就写写懒汉模式

type Singleton struct {
}

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
