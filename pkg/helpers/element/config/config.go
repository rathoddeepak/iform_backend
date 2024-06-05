package config;

type ElementConfig struct {
	currentConfig map[string]interface{}
}

const (
	HAS_OPTIONS = "hasOptions"
)


func New () *ElementConfig {
	return &ElementConfig{
		currentConfig: map[string]interface{}{},
	};
}

/**
 * Has Mulitple options
 * used for types like mulitple choice options
*/
func (ec *ElementConfig) AddHasOptions (val bool) {
	ec.currentConfig[HAS_OPTIONS] = val;
}

/**
 * Returns format that can be converted to json
*/
func (ec *ElementConfig) ToMap () map[string]interface{} {
	return ec.currentConfig;
}