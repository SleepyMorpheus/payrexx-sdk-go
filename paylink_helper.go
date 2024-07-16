package payrexxsdk

import (
	"encoding/json"
	"fmt"
	"payrexxsdk/internal"
)

func (g PaylinkBody) String() string {
	outBody, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("PaylinkBody: %s", string(outBody))
}

func (g PaylinkHead) String() string {
	outHead, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("PaylinkHead: %s", string(outHead))
}

func (g Paylink) String() string {
	return fmt.Sprintf("%s\n%s", g.PaylinkHead.String(), g.PaylinkBody.String())
}

func (b *PaylinkBody) MarshalJSON() ([]byte, error) {
	// Create a shadow type to avoid infinite recursion
	type Alias PaylinkBody
	return json.Marshal(&struct {
		Purpose                          map[string]string `json:"purpose"`
		SubscriptionInterval             string            `json:"subscriptionInterval,omitempty"`
		SubscriptionPeriod               string            `json:"subscriptionPeriod,omitempty"`
		SubscriptionCancellationInterval string            `json:"subscriptionCancellationInterval,omitempty"`
		*Alias
	}{
		Purpose:                          map[string]string{"1": b.Purpose},
		SubscriptionInterval:             internal.DurationToString(&b.SubscriptionInterval),
		SubscriptionCancellationInterval: internal.DurationToString(&b.SubscriptionCancellationInterval),
		SubscriptionPeriod:               internal.DurationToString(&b.SubscriptionPeriod),
		Alias:                            (*Alias)(b),
	})
}

// UnmarshalJSON employes the default json bytes to struct transformation
// This is required, because payrexx has an error in their api docs
// for the field purpose and we need to correctly parse that.
func (b *PaylinkBody) UnmarshalJSON(data []byte) error {
	// Create a shadow type to avoid infinite recursion
	type Alias PaylinkBody
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
		return fmt.Errorf("PaylinkBody parsing error: purpose key '1' not found")
	}

	subIntDur, err := internal.StringToDuration(aux.SubscriptionInterval)
	if err != nil {
		return fmt.Errorf("PaylinkBody parsing error: Failed to unmarshall 'subscriptionInterval' due to %w", err)
	}

	subPerDur, err := internal.StringToDuration(aux.SubscriptionPeriod)
	if err != nil {
		return fmt.Errorf("PaylinkBody parsing error: Failed to unmarshall 'subscriptionPeriod' due to %w", err)
	}

	subCanDur, err := internal.StringToDuration(aux.SubscriptionCancellationInterval)
	if err != nil {
		return fmt.Errorf("PaylinkBody parsing error: Failed to unmarshall 'subscriptionCancellationInterval' due to %w", err)
	}

	b.Purpose = purposeStr
	b.SubscriptionPeriod = *subPerDur
	b.SubscriptionInterval = *subIntDur
	b.SubscriptionCancellationInterval = *subCanDur
	return nil
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
