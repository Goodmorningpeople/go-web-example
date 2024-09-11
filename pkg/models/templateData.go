package models

// data to be passed into template through render.go, pass as parameter into handler func 
type TemplateData struct {
	StringMap map[string]string
	IntMap map[int]int
	Floatmap map[float32]float32
	Data map[string]interface{}
	CSRFToken string	
	Flash string
	Warning string
	Error string
}