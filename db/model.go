package db

type id int

// Record structure that provides a connection
// between the business logic of the model used
// and the database.
type Record struct {
	ID        id
	Data      SampleData
	IsDeleted bool
}

//Model implements the structure used to define
//a physical table in the database
type Model struct {
}

//GlobalControl Internal data control register
type GlobalControl struct {
	//autoincremented id to manage registers
	baseID id

	//Storage for our database records
	DbStorage map[id]Record
}

//SampleData table used to test manually the functionality
type SampleData struct {
	FirstName string
	LastName  string
	Email     string
	Booze     string
}
