package question_types;


type ElementBase struct {
	Id string `json:"id"`
	Title string `json:"title"`
	DefaultValue string `json:"default_value"`
	Required bool `json:"required"`
	Validations []map[string]interface{} `json:"validations"`
	Config map[string]interface{} `json:"config"`
};