package structs

// FieldInfo configuration of a field
type FieldInfo struct {
	// Name for field
	Name string `json:"name"`
	// Type for filed as: varchar, int...
	Type string `json:"type"`
	// IsNull allow the field is null
	IsNull bool `json:"is_null"`
}


