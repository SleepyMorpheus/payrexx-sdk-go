package paylink

import (
	"encoding/json"
	"fmt"
)

// Paylink is a combination of Body & Head representing a
// complete Paylink
type Paylink struct {
	PaylinkHead
	PaylinkBody
	Purpose map[string]string
}

func (g Paylink) String() string {
	return fmt.Sprintf("%s\n%s", g.PaylinkHead.String(), g.PaylinkBody.String())
}

// UnmarshalJSON invokes the default json bytes to struct transformation0n
// for both head and body.
//
// We need to handle it with a dedicated function because PaylinkBody has a
// custom UnmarshalJSON which leads to the head not being parsed automatically
func (o *Paylink) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.PaylinkHead); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &o.PaylinkBody); err != nil {
		return err
	}
	return nil
}
