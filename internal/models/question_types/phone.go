/**
 * Question Type | Model | Phone
*/
package question_types;

import (
	elementValidation "iform/pkg/helpers/element/validation"
	elementConfig "iform/pkg/helpers/element/config"
)

var defaultElementPhone *ElementBase;

func CreateDefaultPhone () *ElementBase {
	const (
		ELEMENT_ID = "elm_phone";
		ELEMENT_TITLE = "Phone Number";
		DEFAULT_TEXT_LENGTH = 13;
	);

	if defaultElementPhone != nil {
		return defaultElementPhone;
	}

	elm := &ElementBase {
		Id: ELEMENT_ID,
		Title: ELEMENT_TITLE,
		Required: false,
	};

	validations := elementValidation.New();
	validations.AddMaxCharecterLength(DEFAULT_TEXT_LENGTH);
	validations.AddValidPhone();

	config := elementConfig.New();
	config.AddHasOptions(false);

	elm.Validations = validations.ToArray();
	elm.Config = config.ToMap();

	// We don't want to create every time	
	defaultElementPhone = elm;

	return elm;
}