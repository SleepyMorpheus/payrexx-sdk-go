package transaction

import "time"

type Transaction struct {
	Id                  string             `json:"id"`
	Uuid                string             `json:"uuid"`
	Status              string             `json:"status"`
	Time                time.Time          `json:"time"`
	Lang                string             `json:"lang"`
	PageUuid            string             `json:"pageUuid"`
	Payment             TransactionPayment `json:"payment"`
	PayoutUuid          string             `json:"payoutUuid"`
	Psp                 string             `json:"psp"`
	PspId               int32              `json:"pspId"`
	Mode                string             `json:"mode"`
	ReferenceId         string             `json:"referenceId"`
	Invoice             TransactionInvoice `json:"invoice"`
	Refundable          bool               `json:"refundable"`          // todo: default true
	PartiallyRefundable bool               `json:"partiallyRefundable"` // todo: default true
	// Contact todo: create struct for contact
}

type TransactionPayment struct {
	Brand                        string                                  `json:"brand"`
	Wallet                       string                                  `json:"wallet"`
	PurchaseOnInvoiceInformation TransactionPurchaseOnInvoiceInformation `json:"purchaseOnInvoiceInformation"`
}

type TransactionPurchaseOnInvoiceInformation struct {
	Zip       string `json:"zip"`
	Iban      string `json:"iban"`
	Place     string `json:"place"`
	Swift     string `json:"swift"`
	Address   string `json:"address"`
	Company   string `json:"company"`
	BankName  string `json:"bankName"`
	Reference string `json:"reference"`
}

type TransactionInvoice struct {
	CurrencyAlpha3 string                           `json:"currencyAlpha3"`
	Products       []TransactionInvoiceProduct      `json:"products"`
	Discount       TransactionInvoiceDiscount       `json:"discount"`
	ShippingAmount int32                            `json:"shippingAmount"`
	TotalAmount    int32                            `json:"totalAmount"`
	CustomFields   []TransactionInvoiceCustomFields `json:"customFields"` // todo: handle custom {"20": {"name...
}

type TransactionInvoiceProduct struct {
	Quantity int32  `json:"quantity"`
	Name     string `json:"name"`
	Amount   int32  `json:"amount"`
}

type TransactionInvoiceDiscount struct {
	Code       string `json:"code"`
	Percentage int32  `json:"percentage"`
	Amount     int32  `json:"amount"`
}

type TransactionInvoiceCustomFields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
