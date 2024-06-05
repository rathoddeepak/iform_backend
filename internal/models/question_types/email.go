/**
 * Question Type | Model | Email
*/
package question_types;

import (
	elementValidation "iform/pkg/helpers/element/validation"
	elementConfig "iform/pkg/helpers/element/config"
)

var defaultElementEmail *ElementBase;

func CreateDefaultEmail () *ElementBase {

	const (
		ELEMENT_ID = "elm_email";
		ELEMENT_TITLE = "Email";
		DEFAULT_TEXT_LENGTH = 320;
	);
	
	if defaultElementEmail != nil {
		return defaultElementEmail;
	}

	elm := &ElementBase {
		Id: ELEMENT_ID,
		Title: ELEMENT_TITLE,
		Required: false,
	};

	validations := elementValidation.New();
	validations.AddMaxCharecterLength(DEFAULT_TEXT_LENGTH);
	validations.AddValidEmail();

	config := elementConfig.New();
	config.AddHasOptions(false);

	elm.Validations = validations.ToArray();
	elm.Config = config.ToMap();

	// We don't want to create every time	
	defaultElementEmail = elm;

	return elm;
}