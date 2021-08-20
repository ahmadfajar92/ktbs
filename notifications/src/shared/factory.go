package shared

import "notifications/config"

// usecaseFactory struct
type (
	usecaseFactory struct {
		usecases map[string]interface{}
	}
)

func UsecaseFactory(cfg config.Config) Usecases {
	factory := new(usecaseFactory)
	factory.usecases = make(map[string]interface{}, 0)
	return factory
}

func (ufc *usecaseFactory) Call(u string) interface{} {
	return ufc.usecases[u]
}

func (ufc *usecaseFactory) Add(n string, u interface{}) {
	ufc.usecases[n] = u
}
