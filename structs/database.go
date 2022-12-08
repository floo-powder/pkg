package structs

import "github.com/floo-powder/pkg/consts"

type DatabaseConfigField struct {
	// Name for database config field
	Name string `json:"name"`
	// Kind for database config field
	Kind consts.DatabaseConfigFieldKind `json:"kind"`
	// Description for database config field
	Description string `json:"description"`
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
