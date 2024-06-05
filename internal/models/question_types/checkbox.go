/**
 * Question Type | Model | Multiple Choice
*/
package question_types;

import (
	elementValidation "iform/pkg/helpers/element/validation"
	elementConfig "iform/pkg/helpers/element/config"
)

var defaultElementCheckbox *ElementBase;

func CreateDefaultCheckbox () *ElementBase {

	const (
		ELEMENT_ID = "elm_checkbox";
		ELEMENT_TITLE = "Checkbox";
		DEFAULT_OPTION_TEXT_LENGTH = 1000;
	);
	
	if defaultElementCheckbox != nil {
		return defaultElementCheckbox;
	}

	elm := &ElementBase {
		Id: ELEMENT_ID,
		Title: ELEMENT_TITLE,
		Required: false,
	};

	validations := elementValidation.New();
	validations.AddOptionMaxSelectLength();

	config := elementConfig.New();
	config.AddHasOptions(true);

	elm.Validations = validations.ToArray();
	elm.Config = config.ToMap();

	// We don't want to create every time	
	defaultElementCheckbox = elm;

	return elm;
}