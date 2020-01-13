package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	cpb "storage-service/internal/storage"
)

// Contacts -
type Contacts struct {
	Contacts []*cpb.Contact `json:"contacts"`
}

var contacts Contacts

// ReadData reads contact data from json file
func ReadData() []*cpb.Contact {
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/internal/db/contacts.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err2 := json.Unmarshal(byteValue, &contacts)
	if err2 != nil {
		fmt.Println(err)
	}

	// for i := 0; i < len(contacts.Contacts); i++ {
	// 	fmt.Println("ID: " + strconv.Itoa(int(contacts.Contacts[i].Id)))
	// 	fmt.Println("Contact FirstName: " + contacts.Contacts[i].FirstName)
	// 	fmt.Println("Contact LastName: " + contacts.Contacts[i].LastName)
	// 	fmt.Println("Contact Email: " + contacts.Contacts[i].Email)
	// 	fmt.Println("Contact PhoneNumber: " + contacts.Contacts[i].PhoneNumber)
	// }

	// fmt.Println(contacts.Contacts)

	return contacts.Contacts

}
