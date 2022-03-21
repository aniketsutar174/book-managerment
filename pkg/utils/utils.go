//this code used to unmarshal the data from JSON
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//to pass the body in create function in controlls,
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { //reading a body
		if err := json.Unmarshal([]byte(body), x); err != nil { //here we unmarshling json
			return
		}
	}
}
