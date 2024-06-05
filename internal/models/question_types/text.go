/**
 * Question Type | Model | Text
*/
package question_types;

import (
	elementValidation "iform/pkg/helpers/element/validation"
	elementConfig "iform/pkg/helpers/element/config"
)

var defaultElementText *ElementBase;

func CreateDefaultText () *ElementBase {
	
	const (
		ELEMENT_ID = "elm_text";
		ELEMENT_TITLE = "Text";
		DEFAULT_TEXT_LENGTH = 1000;
	);

	if defaultElementText != nil {
		return defaultElementText;
	}

	elm := &ElementBase {
		Id: ELEMENT_ID,
		Title: ELEMENT_TITLE,
		Required: false,
	};

	validations := elementValidation.New();
	validations.AddMaxCharecterLength(DEFAULT_TEXT_LENGTH);

	config := elementConfig.New();
	config.AddHasOptions(false);

	elm.Validations = validations.ToArray();
	elm.Config = config.ToMap();

	// We don't want to create every time	
	defaultElementText = elm;

	return elm;
}