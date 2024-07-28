package gateway

import (
	"encoding/json"
	"fmt"
)

// Gateway is a combination of Body & Head representing a
// complete gateway
type Gateway struct {
	GatewayHead
	GatewayBody
}

func (g Gateway) String() string {
	return fmt.Sprintf("%s\n%s", g.GatewayHead.String(), g.GatewayBody.String())
}

// UnmarshalJSON invokes the default json bytes to struct transformation0n
// for both head and body.
//
// We need to handle it with a dedicated function because GatewayBody has a
// custom UnmarshalJSON which leads to the head not being parsed automatically
func (o *Gateway) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.GatewayHead); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &o.GatewayBody); err != nil {
		return err
	}
	return nil
}
