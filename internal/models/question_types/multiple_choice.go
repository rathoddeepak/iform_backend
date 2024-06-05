/**
 * Question Type | Model | Multiple Choice
*/
package question_types;

import (
	elementValidation "iform/pkg/helpers/element/validation"
	elementConfig "iform/pkg/helpers/element/config"
)

var defaultElementChoice *ElementBase;

func CreateDefaultMultipleChoice () *ElementBase {

	const (
		ELEMENT_ID = "elm_choice";
		ELEMENT_TITLE = "Multiple Choice";
		DEFAULT_OPTION_TEXT_LENGTH = 1000;
	);
	
	if defaultElementChoice != nil {
		return defaultElementChoice;
	}

	elm := &ElementBase {
		Id: ELEMENT_ID,
		Title: ELEMENT_TITLE,
		Required: false,
	};

	validations := elementValidation.New();

	config := elementConfig.New();
	config.AddHasOptions(true);

	elm.Validations = validations.ToArray();
	elm.Config = config.ToMap();

	// We don't want to create every time	
	defaultElementChoice = elm;

	return elm;
}