package payrexxsdk

import (
	"encoding/json"
	"fmt"
	"payrexxsdk/internal"
)

func (g GatewayBody) String() string {
	outBody, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("GatewayBody: %s", string(outBody))
}

func (g GatewayHead) String() string {
	outHead, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("GatewayHead: %s", string(outHead))
}
func (g Gateway) String() string {
	return fmt.Sprintf("%s\n%s", g.GatewayHead.String(), g.GatewayBody.String())
}

func (b *GatewayBody) MarshalJSON() ([]byte, error) {
	// Create a shadow type to avoid infinite recursion
	type Alias GatewayBody
	return json.Marshal(&struct {
		Purpose                          map[string]string `json:"purpose"`
		SubscriptionInterval             string            `json:"subscriptionInterval,omitempty"`
		SubscriptionPeriod               string            `json:"subscriptionPeriod,omitempty"`
		SubscriptionCancellationInterval string            `json:"subscriptionCancellationInterval,omitempty"`
		*Alias
	}{
		Purpose:                          map[string]string{"1": b.Purpose},
		SubscriptionInterval:             internal.DurationToString(&b.SubscriptionInterval),
		SubscriptionCancellationInterval: internal.DurationToString(&b.subscriptionCancellationInterval),
		SubscriptionPeriod:               internal.DurationToString(&b.SubscriptionPeriod),
		Alias:                            (*Alias)(b),
	})
}

// UnmarshalJSON employes the default json bytes to struct transformation
// This is required, because payrexx has an error in their api docs
// for the field purpose and we need to correctly parse that.
func (b *GatewayBody) UnmarshalJSON(data []byte) error {
	// Create a shadow type to avoid infinite recursion
	type Alias GatewayBody

	aux := &struct {
		SubscriptionInterval             string            `json:"subscriptionInterval,omitempty"`
		SubscriptionPeriod               string            `json:"subscriptionPeriod,omitempty"`
		SubscriptionCancellationInterval string            `json:"subscriptionCancellationInterval,omitempty"`
		Purpose                          map[string]string `json:"purpose"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Extract the date string from the map
	purposeStr, ok := aux.Purpose["1"]
	if !ok {
		return fmt.Errorf("GatewayBody parsing error: purpose key '1' not found")
	}

	subIntDur, err := internal.StringToDuration(aux.SubscriptionInterval)
	if err != nil {
		return fmt.Errorf("GatewayBody parsing error: Failed to unmarshall 'subscriptionInterval' due to %w", err)
	}

	subPerDur, err := internal.StringToDuration(aux.SubscriptionPeriod)
	if err != nil {
		return fmt.Errorf("GatewayBody parsing error: Failed to unmarshall 'subscriptionPeriod' due to %w", err)
	}

	subCanDur, err := internal.StringToDuration(aux.SubscriptionCancellationInterval)
	if err != nil {
		return fmt.Errorf("GatewayBody parsing error: Failed to unmarshall 'subscriptionCancellationInterval' due to %w", err)
	}

	b.Purpose = purposeStr
	b.SubscriptionPeriod = *subPerDur
	b.SubscriptionInterval = *subIntDur
	b.subscriptionCancellationInterval = *subCanDur
	return nil
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
