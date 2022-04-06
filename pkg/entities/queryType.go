package entities

type Query struct {
	Type   string   `query:"type"`
	Fields []string `query:"fields"`
}
