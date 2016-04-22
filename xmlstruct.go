package main

import (
	"encoding/xml"
	"fmt"
)

const str = `<?xml version="1.0" encoding="UTF-8" ?>
<countries>
	<country>
		<code>CH</code>
		<name>Schweiz</name>
		<regions>
			<region>
				<code>10</code>
				<name>Mittelland</name>
				<subregions>
					<subregion>
						<code>11</code>
						<name>Blah blah</name>
					</subregion>
				</subregions>
			</region>
		</regions>
	</country>
</countries>`

type Destination struct {
	CountryCode string    `xml:"code"`
	CountryName string    `xml:"name"`
	Regions     []*Region `xml:"regions>region"`
}

type Region struct {
	Code       int          `xml:"code"`
	Name       string       `xml:"name"`
	Subregions []*Subregion `xml:"subregions>subregion"`
}

type Subregion struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

type Destinations struct {
	Destinations []Destination `xml:"country"`
}

func main() {
	d := Destinations{}

	data := []byte(str) //	data, _ := ioutil.ReadFile("xml/countryregion.xml")

	err := xml.Unmarshal(data, &d)
	if err != nil {
		fmt.Printf("error: ", err)
		return
	}

	for _, e := range d.Destinations {
		fmt.Printf("Country: %s - %s\n", e.CountryCode, e.CountryName)
		for _, r := range e.Regions {
			fmt.Printf("Region: %d - %s\n", r.Code, r.Name)
			for _, subr := range r.Subregions {
				//Если len(r.Subregions) == 0, то и так ничего не будет
				fmt.Printf("\tSubregion of %s: %d - %s\n", r.Name, subr.Code, subr.Name)
			}
		}
	}
}

/*
Country: CH - Schweiz
Region: 10 - Mittelland
	Subregion of Mittelland: 11 - Blah blah
*/
