// from https://habrahabr.ru/post/255043/
package main

import (
	"encoding/json"
	"fmt"
)

type X struct {
	Time json.Number `json:"time,omitempty"`
	Name json.Number `json:"name,omitempty"` // haven't in json
	Ya   map[string]interface{}      `json:"y,omitempty"`
	// Yb   string      `json:"z,omitempty"`
}

func main() {
	//var m1 map[string]interface{}
	//_ = json.Unmarshal([]byte(`{"x":1,"y":{}}`), &m1)
	//fmt.Println(m1)
	var x X
	_ = json.Unmarshal([]byte(`{"x":1, "y":{"a":5}}`), &x)
	fmt.Println(x.Time, x.Ya)
	_ = json.Unmarshal([]byte(`{"time":123, "x":1, "y":{"a":5}}`), &x)
	fmt.Println(x.Time, x.Ya)
	_ = json.Unmarshal([]byte(`{"time":"123", "x":1, "y":{"b":7}}`), &x)
	fmt.Println(x.Time, x.Ya["b"])
}

/* output
 map[a:5]
123 map[a:5]
123 7
*/
