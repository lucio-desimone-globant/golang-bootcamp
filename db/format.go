package db

import "encoding/json"

//ToJSON TODO: add JSON serialization to db
func ToJSON(record Record) []byte {
	jsonRecord, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}
	return jsonRecord
}
