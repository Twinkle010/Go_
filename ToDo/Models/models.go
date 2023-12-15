package Models

type ToDo struct {
	ID uint            `json:"id"`
	Title string       `json:"title"`
	Description string `json:"description"`
}

//  json:"id" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
// We can add validations to this struct as well, for eg: binding:”required” will return an error if it has empty value while binding.
func (b *ToDo) ToDo() string {
	return "todo"
}