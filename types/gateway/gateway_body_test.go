package gateway

import (
	"encoding/json"
	"github.com/sosodev/duration"
	"testing"
)

const gatewayBodyJson = `
	{
      "id": 1,
      "status": "waiting",
      "hash": "db8c618c87dc91f100292f6ffd9c5044",
      "referenceId": "975382",
      "link": "https://demo.payrexx.com/?payment=db8c618c87dc91f100292f6ffd9c5044",
      "invoices": [],
      "preAuthorization": false,
      "reservation": false,
      "fields": {
        "title": {
          "active": false,
          "mandatory": true
        },
        "forename": {
          "active": false,
          "mandatory": true
        },
        "surname": {
          "active": false,
          "mandatory": true
        },
        "company": {
          "active": false,
          "mandatory": true
        },
        "street": {
          "active": false,
          "mandatory": true
        },
        "postcode": {
          "active": false,
          "mandatory": true
        },
        "place": {
          "active": false,
          "mandatory": true
        },
        "country": {
          "active": false,
          "mandatory": true
        },
        "phone": {
          "active": false,
          "mandatory": true
        },
        "email": {
          "active": false,
          "mandatory": true
        },
        "date_of_birth": {
          "active": false,
          "mandatory": true,
          "names": {
            "de": "",
            "en": "",
            "fr": "",
            "it": ""
          }
        },
        "terms": {
          "active": true,
          "mandatory": true
        },
        "privacy_policy": {
          "active": true,
          "mandatory": true
        },
        "text": {
          "active": false,
          "mandatory": true,
          "names": {
            "de": "Benutzerdefiniertes Feld (DE)",
            "en": "Benutzerdefiniertes Feld (EN)",
            "fr": "Benutzerdefiniertes Feld (FR)",
            "it": "Benutzerdefiniertes Feld (IT)"
          }
        }
      },
      "psp": [],
      "pm": [],
      "amount": 8925,
      "vatRate": 7.7,
      "currency": "CHF",
      "sku": "P01122000",
      "createdAt": 1475578052
    }
`

func TestGatewayBody_UnmarshalJSON(t *testing.T) {

	var g = GatewayBody{}
	err := json.Unmarshal([]byte(gatewayBodyJson), &g)
	if err != nil {
		t.Error(err)
	}

}

func TestGatewayBody_MarshalJSON(t *testing.T) {

	var g = GatewayBody{
		SubscriptionInterval:             duration.Duration{Months: 1},
		SubscriptionPeriod:               duration.Duration{Months: 2},
		SubscriptionCancellationInterval: duration.Duration{Months: 3},
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

}
