package common

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

// FieldInfo configuration of a field
type FieldInfo struct {
	// Name for field
	Name string `json:"name"`
	// Type for filed as: varchar, int...
	Type string `json:"type"`
	// IsNull allow the field is null
	IsNull bool `json:"is_null"`
}

// TableInfo Information about the tables used
type TableInfo struct {
	// Name for table
	Name string `json:"name"`
	// FieldConfig field configuration
	FieldConfig []FieldInfo `json:"field_config"`
}

// ActuatorConfig configuration for Actuator
type ActuatorConfig struct {
	// SourceName the name of the data source
	SourceName string `json:"source_name"`
	// TableConfig Use table information
	TableConfig TableInfo `json:"table_config"`
}
