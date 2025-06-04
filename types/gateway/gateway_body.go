package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/SleepyMorpheus/payrexx-sdk-go/internal"
	"github.com/sosodev/duration"
)

// GatewayBody represents the data needed to create a gateway at payrexx
type GatewayBody struct {
	// (REQ) Amount of payment in cents
	Amount int32 `json:"amount"`
	// (OPT) VAT Rate Percentage
	VatRate float32 `json:"vatRate,omitempty"`
	// (REQ) Currency of payment (ISO code)
	Currency string `json:"currency"`
	// (OPT) Product stock keeping unit
	Sku string `json:"sku,omitempty"`
	// (OPT) The purpose of the payment
	Purpose string `json:"purpose,omitempty"`
	// (OPT) URL to redirect to after successful payment.
	SuccessRedirectUrl string `json:"successRedirectUrl,omitempty"`
	// (OPT) URL to redirect to after failed payment.
	FailedRedirectUrl string `json:"failedRedirectUrl,omitempty"`
	// (OPT) URL to redirect to after manual cancellation by shopper.
	CancelRedirectUrl string `json:"cancelRedirectUrl,omitempty"`
	// todo: implement basket
	// (OPT) List of PSPs to provide for payment. If empty all available PSPs are provided.
	Psp []int32 `json:"psp,omitempty"`
	// (OPT) List of payment mean names to display
	Pm []string `json:"pm,omitempty"`
	// (OPT) Whether charge payment manually at a later date (type authorization)
	PreAuthorization bool `json:"preAuthorization"`
	// (OPT) Whether charge payment manually at a later date (type reservation)
	Reservation bool `json:"reservation"`
	// (OPT) An internal reference id used by your system.
	ReferenceId string `json:"referenceId,omitempty"`
	// todo: implement fields
	// (OPT) Only available for Concardis PSP and if the custom ORDERID option is activated in PSP settings in Payrexx administration. This ORDERID will be transferred to the Payengine.
	ConcardisOrderId string `json:"concardisOrderId,omitempty"`
	// (OPT) Skip result page and directly redirect to success or failed URL
	SkipResultPage bool `json:"skipResultPage"`
	// (OPT) preAuthorization needs to be "true". This will charge the authorization during the first payment.
	ChargeOnAuthorization bool `json:"chargeOnAuthorization"`
	// (OPT) Gateway validity in minutes.
	Validity int32 `json:"validity"`
	// (OPT) Defines whether the payment should be handled as subscription.
	SubscriptionState bool `json:"subscriptionState"`
	// (OPT) Payment interval
	SubscriptionInterval duration.Duration `json:"subscriptionInterval"`
	// (OPT) Duration of the subscription
	SubscriptionPeriod duration.Duration `json:"subscriptionPeriod"`
	// (OPT) Defines the period, in which a subscription can be canceled.
	SubscriptionCancellationInterval duration.Duration `json:"subscriptionCancellationInterval"`
	// (OPT) Change the default button Text "Pay" to a custom String
	ButtonText []string `json:"buttonText,omitempty"`
	// (OPT) UUID of the look and feel profile to use.
	LookAndFeelProfile string `json:"lookAndFeelProfile,omitempty"`
	// (OPT) Custom success message on result page.
	SuccessMessage string `json:"successMessage,omitempty"`
	// (OPT) Holds the session ID of a scanned QR code. Only available and needed for Static QR Code with Twint.
	QrCodeSessionId string `json:"qrCodeSessionId,omitempty"`
}

func (g GatewayBody) String() string {
	outBody, err := json.Marshal(g)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("GatewayBody: %s", string(outBody))
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
		SubscriptionCancellationInterval: internal.DurationToString(&b.SubscriptionCancellationInterval),
		SubscriptionPeriod:               internal.DurationToString(&b.SubscriptionPeriod),
		Alias:                            (*Alias)(b),
	})
}

// UnmarshalJSON employes the default json bytes to struct transformation
// This is required, because payrexx has an error in their api docs
// for the field purpose, and we need to correctly parse that.
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

	b.SubscriptionPeriod = *subPerDur
	b.SubscriptionInterval = *subIntDur
	b.SubscriptionCancellationInterval = *subCanDur
	return nil
}
