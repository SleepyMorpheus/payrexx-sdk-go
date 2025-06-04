package paylink

import (
	"encoding/json"
	"github.com/SleepyMorpheus/payrexx-sdk-go/types/shared"
	"github.com/sosodev/duration"
	"testing"
)

const jsonPaylinkBody = `

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

func TestPaylinkBody_UnmarshalJSON(t *testing.T) {

	var g = PaylinkBody{}
	err := json.Unmarshal([]byte(jsonPaylinkBody), &g)
	if err != nil {
		t.Error(err)
	}

	if g.Fields.CustomField1.Names.De != "This is a field" {
		t.Error("CustomField1.Names.De should be 'This is a field'")
	}

}

func TestPaylinkBody_MarshalJSON(t *testing.T) {

	var g = PaylinkBody{
		SubscriptionInterval:             duration.Duration{Months: 1},
		SubscriptionPeriod:               duration.Duration{Months: 2},
		SubscriptionCancellationInterval: duration.Duration{Months: 3},
		Fields: PaylinkBodyFields{
			CustomField2: PaylinkBodyFieldTranslatable{
				Active:    true,
				Mandatory: false,
				Names:     shared.Translation{De: "De", En: "En", Fr: "Fr", It: "It"},
			},
		},
	}

	jout, err := json.Marshal(g)
	if err != nil {
		t.Error(err)
	}

	// reparse it as interface
	var i = map[string]interface{}{}
	err = json.Unmarshal(jout, &i)
	if err != nil {
		t.Error(err)
	}

	if i["subscriptionInterval"] != "P1M" {
		t.Error("SubscriptionInterval: P1M !=", i["subscriptionInterval"])
	}

	if i["subscriptionPeriod"] != "P2M" {
		t.Error("SubscriptionPeriod: P2M !=", i["subscriptionPeriod"])
	}

	if i["subscriptionCancellationInterval"] != "P3M" {
		t.Error("SubscriptionCancellationInterval: P3M !=", i["subscriptionCancellationInterval"])
	}

	customField2 := i["fields"].(map[string]interface{})["custom_field_2"].(map[string]interface{})

	if customField2["active"] != true || customField2["mandatory"] != false {
		t.Error("Custom field 2 has invalid boolean")
	}

	names := customField2["names"].(map[string]interface{})
	if names["de"] != "De" {
		t.Error("Custom field 2 has invalid de")
	}

}
