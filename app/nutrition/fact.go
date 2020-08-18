package nutrition

type carbohydrates string
type proteins string
type name string
type description string
type recommendendAmount string
type upperLimitPerDay string

// vitamin defines an vitamin structure
type vitamin struct {
	Name               name
	Description        description
	RecommendendAmount recommendendAmount
	UpperLimitPerDay   upperLimitPerDay
}

type vitamins []vitamin

type facts struct {
	Vitamins      vitamins
	Carbohydrates carbohydrates
	Proteins      proteins
}
