package gateway

import (
	"encoding/json"
	"testing"
	"time"
)

const gatewayHeadJson = `
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

func TestGatewayHead_UnmarshalJSON(t *testing.T) {

	var g = GatewayHead{}
	err := json.Unmarshal([]byte(gatewayHeadJson), &g)
	if err != nil {
		t.Error(err)
	}

	if g.CreatedAt != time.Unix(1475578052, 0) {
		t.Error("Failed to parse CreatedAt")
	}

}
