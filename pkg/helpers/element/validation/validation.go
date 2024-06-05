package validation;

type ElementValidations struct {
	list []map[string]interface{}
}

const (	
	TEXT_CHAR_LEN = "text_char_len"
	VALID_PHONE = "valid_phone"
	VALID_EMAIL = "valid_email"
	OPTION_SELECT_LIMIT = "option_select_limit"
)


func createValidation (id string, defaultVal interface{}) map[string]interface{} {
	validation := map [string] interface {} {};
	validation["id"] = id;
	validation["defaultValue"] = defaultVal;
	return validation;
}

func New () *ElementValidations {
	return &ElementValidations{
		list: []map[string]interface{}{},
	};
}

/**
 * @Element: Text
*/
func (ev *ElementValidations) AddMaxCharecterLength (val int) {
	ev.list = append(
		ev.list,
		createValidation(TEXT_CHAR_LEN, val),
	);
}

/**
 * @Element: Phone
*/
func (ev *ElementValidations) AddValidPhone () {
	ev.list = append(
		ev.list,
		createValidation(VALID_PHONE, ""),
	);
}

/**
 * @Element: Email
*/
func (ev *ElementValidations) AddValidEmail () {
	ev.list = append(
		ev.list,
		createValidation(VALID_EMAIL, ""),
	);
}

/**
 * @Element: Text, Phone, Email
*/
func (ev *ElementValidations) AddOptionMaxSelectLength () {
	ev.list = append(
		ev.list,
		createValidation(OPTION_SELECT_LIMIT, ""),
	);
}

/**
 * Returns format that can be converted to json
*/
func (ev *ElementValidations) ToArray () []map[string]interface{} {
	return ev.list;
}
