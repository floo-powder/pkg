package consts

type DatabaseKind string

const (
	Mysql    DatabaseKind = "Mysql"
	Hive     DatabaseKind = "Hive"
	Postgres DatabaseKind = "Postgres"
	Oracle   DatabaseKind = "Oracle"
)

type DatabaseConfigFieldKind string

const (
	IntConfig      DatabaseConfigFieldKind = "int"
	BoolConfig     DatabaseConfigFieldKind = "bool"
	PasswordConfig DatabaseConfigFieldKind = "password"
	StringConfig   DatabaseConfigFieldKind = "string"
)

type DatabaseConfigField struct {
	// Name for database config field
	Name string `json:"name"`
	// Kind for database config field
	Kind DatabaseConfigFieldKind `json:"kind"`
	// Description for database config field
	Description string `json:"description"`
}
