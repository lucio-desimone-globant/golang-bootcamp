package db

//Create a new record using the data received
func Create(data SampleData) Record {
	var newRecord = Record{
		ID:        autoIncrement(&DataBase.baseID),
		Data:      data,
		IsDeleted: false,
	}

	//allocate to store in the database the new record
	addRecord(newRecord)

	return newRecord
}

//FindByID a
func FindByID(id id) (Record, error) {
	record, ok := DataBase.DbStorage[id]
	if !ok {
		return Record{}, ErrFind()
	}
	return record, nil
}

//FindByFirstName find a record by firstName if doesn't exist
//return an error describing the situation
func FindByFirstName(field string) (Record, error) {
	for _, record := range DataBase.DbStorage {
		if record.Data.FirstName == field {
			return record, nil
		}
	}
	return Record{}, ErrFind()
}

//DestroyByID deletes a record physically with the ID
//received from the database if doesn't exist return false
//if it's successfully return true
func DestroyByID(id id) bool {
	_, ok := DataBase.DbStorage[id]
	if !ok {
		return ok
	}
	delete(DataBase.DbStorage, id)
	return ok
}

//DeleteByID deletes a record logically with the ID received
//always changing the isDeleted value to true from the database
//this will have no effect on one that is already deleted
//returns a boolean describing the result
func DeleteByID(id id) bool {
	value, ok := DataBase.DbStorage[id]
	if !ok {
		return ok
	}
	value.IsDeleted = true
	DataBase.DbStorage[id] = value
	return ok
}

//UpdateByID rewrites the register in the ID with the data entered
func UpdateByID(id int, data SampleData) {

}
