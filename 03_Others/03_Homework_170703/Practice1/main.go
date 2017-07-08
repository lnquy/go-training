/* Write an application that read the json content from homework.txt file with content as below.
	Unmarshall/decode/deserialize that data to an struct object,
	change some property values (anything you want) then marshall/encode/serialize that struct object to
	json and save to homework_new.txt.
	Remember to use defer to handle resource, handle errors, and practice with embedded/anonymous struct.

// homework.txt
// Hint: Create a Name struct then embeds that struct into parent struct.
{
  "id": "1",
  "name": {
    "first_name": "Alan",
    "last_name": "Turing"
  },
  "class": "BXDA123",
  "school": "MIT",
  "message": "Hello World"
}
*/

package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"gitlab.com/lnquy/go-training/03_Others/03_Homework_170703/Practice1/models"
)

func main() {
	// Read file
	// TODO (for you :)): Use os.Open to read file line by line, then use defer to handle closing resource
	raw, err := ioutil.ReadFile("homework.txt")
	if err != nil {
		fmt.Printf("Cannot read homework.txt file: %s", err)
		return
	}

	// Unmarshal
	s := &models.Student{} // s is the pointer to Student struct (*models.Student)
	if err := json.Unmarshal(raw, s); err != nil {
		fmt.Printf("Cannot unmarshal raw []byte data to Student struct: %s", err)
		return
	}
	fmt.Printf("Student after decoded from file: %v\n", s)

	// Change struct values
	s.ChangeFirstName("Fermat")
	s.ChangeMessage("Value changed via ChangeMessage function")
	s.Class = "Value changed via directly access Class field"
	fmt.Printf("Student after changed values: %v\n", s)

	// Marshal and write to file
	if sByte, err := json.MarshalIndent(s, "", "    "); err != nil { // Pretty print JSON with 4 spaces
		fmt.Printf("Cannot marshall Student struct to []byte data: %s", err)
	} else {
		if err := ioutil.WriteFile("homework_new.txt", sByte, 0666); err != nil {
			fmt.Printf("Cannot write data to file: %s", err)
		}
	}
}

