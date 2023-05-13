package models

type TemplateData struct {
	Title     string
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	OtherMap  map[string]interface{}
	CSRFToken string
	Message   string
	Error     string
}
