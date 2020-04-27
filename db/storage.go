package db

//DataBase represents the database internally
var DataBase GlobalControl

//Allocate memory to the map and sets the baseId initial value
func init() {
	DataBase = GlobalControl{
		DbStorage: make(map[id]Record),
		baseID:    0,
	}
}

//Internal management of the id assigned
//simulating the SQL AI
func autoIncrement(a *id) id {
	*a = *a + 1
	return *a
}

//Internal management of the store
func addRecord(newRecord Record) Record {
	DataBase.DbStorage[newRecord.ID] = newRecord
	return newRecord
}
