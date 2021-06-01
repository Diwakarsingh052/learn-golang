package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

var xmlData = `<data>
    <person>
        <firstname>Nic</firstname>
        <lastname>Raboy</lastname>
        <address>
            <city>San Francisco</city>
            <state>CA</state>
        </address>
    </person>
    <person>
        <firstname>Maria</firstname>
        <lastname>Raboy</lastname>
		<address>
            
            
        </address>
    </person>
</data>
`

type Data struct {
	XMLName    xml.Name `xml:"data"`
	PersonList []Person `xml:"person"`
}

type Person struct {
	XMLName   xml.Name  `xml:"person"`
	FirstName string    `xml:"firstname"`
	LastName  string    `xml:"lastname"`
	Address   []Address `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	var data Data
	err := xml.Unmarshal([]byte(xmlData), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

}
