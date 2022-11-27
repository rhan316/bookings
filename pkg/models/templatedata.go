package models

// TemplateData przechowuje dane przesyłane z handlers do templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	// Cross Site Request Forgery Token
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
