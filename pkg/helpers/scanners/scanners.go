package scanners;

import (
  "encoding/json"
  "database/sql/driver"
)

type JSONB map[string]interface{}
type JSONBArray []map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
  return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
  return json.Unmarshal(value.([]byte), &j)
}

func (j JSONBArray) Value() (driver.Value, error) {
  return json.Marshal(j)
}

func (j *JSONBArray) Scan(value interface{}) error {
  return json.Unmarshal(value.([]byte), &j)
}