package migrations

type Migrations interface {
	Run()
}

type Entities []interface{}
