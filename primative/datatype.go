package primative

import (
	"strings"
	"time"
)

// BCDate is a struct representing a BigCommerce date resource
type BCDate struct {
	time.Time
}

// UnmarshalJSON will unmarshall date into time object
func (bcD *BCDate) UnmarshalJSON(input []byte) error {
	strTime := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(time.RFC1123Z, strTime)
	if err != nil {
		return err
	}
	bcD.Time = newTime
	return nil
}
