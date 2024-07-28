package paylink

import (
	"encoding/json"
	"fmt"
	"github.com/sosodev/duration"
	"payrexxsdk/internal"
)

// PaylinkBody represents the data needed to create a Paylink at payrexx
type PaylinkBody struct {
	// (REQ) Title of the payment shown on the page
	Title string `json:"title"`
	// (REQ) Description of the payment shown on the page
	Description string `json:"description"`
	// (REQ) Your internal reference id
	ReferenceId string `json:"referenceId"`
	// (REQ) The Purpose of the payment
	Purpose string `json:"purpose"`
	// (REQ) Amount of payment in cents
	Amount int32 `json:"amount"`
	// (OPT) VAT Rate Percentage (nil meaning not set)
	VatRate float32 `json:"vatRate,omitempty"`
	// (REQ) Currency of payment (ISO code)
	Currency string `json:"currency"`
	// (OPT) Psp represents a list of which payment method is allowed
	Psp []int32 `json:"psp,omitempty"`
	// (OPT) PM is list of payment mean names to display
	Pm []string `json:"pm,omitempty"`
	// (OPT) Sku meaning product stock keeping unit
	Sku string `json:"sku,omitempty"`
	// (opt) Whether charge payment manually at a later date (type authorization).
	PreAuthorization bool `json:"preAuthorization"`
	// (opt) Whether charge payment manually at a later date (type reservation).
	Reservation bool `json:"reservation"`
	// (opt) This is an internal name of the payment page. This name will be displayed to the administrator only.
	Name string `json:"name,omitempty"`
	// (OPT) The contact data fields which should be displayed
	Fields PaylinkBodyFields `json:"fields,omitempty"`
	// (OPT) Hide the whole contact fields section on invoice page
	HideFields bool `json:"hideFields"`
	// (OPT) Only available for Concardis PSP and if the custom ORDERID option is activated in PSP settings in Payrexx administration. This ORDERID will be transferred to the Payengine.
	ConcardisOrderId string `json:"concardisOrderId,omitempty"`
	// (OPT) Custom pay button text.
	ButtonText string `json:"buttonText"`
	// (OPT) Expiration date for link. (Format Y-m-d)
	ExpirationDate string `json:"expirationDate,omitempty"`
	// (OPT) URL to redirect to after successful payment.
	SuccessRedirectUrl string `json:"successRedirectUrl,omitempty"`
	// (OPT) URL to redirect to after failed payment.
	FailedRedirectUrl string `json:"failedRedirectUrl,omitempty"`
	// (OPT) Defines whether the payment should be handled as subscription.
	SubscriptionState bool `json:"subscriptionState"`
	// (OPT) Payment interval (converted from ISO 8601)
	SubscriptionInterval duration.Duration `json:"subscriptionInterval,omitempty"`
	// (OPT) Duration of subscription (converted from ISO 8601)
	SubscriptionPeriod duration.Duration `json:"subscriptionPeriod,omitempty"`
	// (OPT) Defines the period, in which a subscription can be canceled. (converted from ISO 8601)
	SubscriptionCancellationInterval duration.Duration `json:"subscriptionCancellationInterval,omitempty"`
	// (OPT) Provide your customers file attachments.
	// todo: add attachments of type file
}

func (g PaylinkBody) String() string {
	outBody, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("PaylinkBody: %s", string(outBody))
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
