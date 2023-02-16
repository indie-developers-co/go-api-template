package entities

type Entities []interface{}

func Provide() Entities {
	return []interface{}{
		&User{},
	}
}
