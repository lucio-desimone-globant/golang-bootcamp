package main

import (
	"fmt"
	"github/luciods/golang-bootcamp/src/db"
	"os"
)

var data1 = db.SampleData{
	FirstName: "Alfio \"Coco\"",
	LastName:  "Basile",
	Email:     "cocob@gmail.com",
	Booze:     "whisky",
}
var data2 = db.SampleData{
	FirstName: "Bob",
	LastName:  "Marley",
	Email:     "bobm@gmail.com",
	Booze:     "rum",
}
var data3 = db.SampleData{
	FirstName: "Vladimir",
	LastName:  "Putin",
	Email:     "vladp@gmail.com",
	Booze:     "vodka",
}

//Data sample to test the db
var data = []db.SampleData{data1, data2, data3}

func main() {

	added := db.Create(data[0])
	added1 := db.Create(data[1])
	added2 := db.Create(data[2])

	a, err0 := db.FindByFirstName(added.Data.FirstName)
	b, err1 := db.FindByFirstName(added1.Data.FirstName)
	c, err2 := db.FindByFirstName(added2.Data.FirstName)

	if err0 != nil {
		panic(err0)
	}
	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	d, err3 := db.FindByID(added.ID)
	e, err4 := db.FindByID(added1.ID)
	f, err5 := db.FindByID(added2.ID)

	if err3 != nil {
		panic(err3)
	}
	if err4 != nil {
		panic(err4)
	}

	if err5 != nil {
		panic(err5)
	}

	fmt.Printf("Name Alfio > %+v\n", a)
	fmt.Printf("Name Bob > %+v\n", b)
	fmt.Printf("Name Vladimir > %+v\n", c)
	fmt.Printf("ID 1 > %+v\n", d)
	fmt.Printf("ID 2 > %+v\n", e)
	fmt.Printf("ID 3 > %+v\n", f)

	db.DeleteByID(added.ID)

	fmt.Printf("ID 2 deleted db result > %v\n", db.DataBase.DbStorage)

	os.Exit(0)
}
