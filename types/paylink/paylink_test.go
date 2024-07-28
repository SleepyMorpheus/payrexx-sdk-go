package paylink

import (
	"encoding/json"
	"testing"
	"time"
)

const jsonPaylink = `

{
  "id": 1,
  "hash": "382c85eab7a86278e3c3b06a23af2358",
  "referenceId": "Order number of my online shop application",
  "link": "https://demo.payrexx.com/?payment=382c85eab7a86278e3c3b06a23af2358",
  "invoices": [],
  "preAuthorization": false,
  "reservation": false,
  "name": "Online-Shop payment #001",
  "api": true,
  "fields": {
	"title": {
	  "active": true,
	  "mandatory": true
	},
	"forename": {
	  "active": true,
	  "mandatory": true
	},
	"surname": {
	  "active": true,
	  "mandatory": true
	},
	"company": {
	  "active": true,
	  "mandatory": true
	},
	"street": {
	  "active": false,
	  "mandatory": false
	},
	"postcode": {
	  "active": false,
	  "mandatory": false
	},
	"place": {
	  "active": false,
	  "mandatory": false
	},
	"country": {
	  "active": true,
	  "mandatory": true
	},
	"phone": {
	  "active": false,
	  "mandatory": false
	},
	"email": {
	  "active": true,
	  "mandatory": true
	},
	"date_of_birth": {
	  "active": false,
	  "mandatory": false
	},
	"terms": {
	  "active": true,
	  "mandatory": true
	},
	"privacy_policy": {
	  "active": true,
	  "mandatory": true
	},
	"custom_field_1": {
	  "active": true,
	  "mandatory": true,
	  "names": {
		"de": "This is a field",
		"en": "This is a field",
		"fr": "This is a field",
		"it": "This is a field"
	  }
	},
	"custom_field_2": {
	  "active": false,
	  "mandatory": false,
	  "names": {
		"de": "",
		"en": "",
		"fr": "",
		"it": ""
	  }
	},
	"custom_field_3": {
	  "active": false,
	  "mandatory": false,
	  "names": {
		"de": "",
		"en": "",
		"fr": "",
		"it": ""
	  }
	}
  },
  "psp":[
    36
  ],
  "pm": [],
  "purpose": {
	"1": "Test Zahlung"
  },      
  "amount": 590,
  "vatRate": 7.7,
  "currency": "CHF",
  "sku": "P01122000",
  "subscriptionState": false,
  "subscriptionInterval": "",
  "subscriptionPeriod": "",
  "subscriptionPeriodMinAmount": "",
  "subscriptionCancellationInterval": "",
  "createdAt": 1418392958
}
`

func TestPaylink_UnmarshalJSON(t *testing.T) {

	var g = Paylink{}
	err := json.Unmarshal([]byte(jsonPaylink), &g)
	if err != nil {
		t.Error(err)
	}

	if g.CreatedAt != time.Unix(1418392958, 0) {
		t.Error("Failed to parse CreatedAt from Paylink")
	}

	if g.Currency != "CHF" {
		t.Error("Failed to parse Currency from Paylink")
	}

}
