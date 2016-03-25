package main

import (
    "fmt"
    "encoding/json"
    )

type User struct {
    Name string
    Parents struct {
        Mother string
        Father string
    }
}

func main() {
    encoded := `[{"name":"C1","parents":{"mother":"E1","father":"A1"}}, {"name":"C2","parents":{"mother":"E2","father":"A2"}}]`
	//encoded := `{"name":"Cain","parents":{"mother":"Eve","father":"Adam"}}`

    // Decode the json object
    u := &[]User{} // &User{}
    err := json.Unmarshal([]byte(encoded), &u)
    if err != nil {
        panic(err)
    }

   // fmt.Println(u.Parents.Mother)
  //  fmt.Println(u.Parents.Father)
	fmt.Println(u)
}
