package paylink

import "payrexxsdk/types/shared"

/*
If active, the field will appear on the paylink page to be filled in by the user.
Mandatory defines if a user must fill in the field.
The custom fields require a name translated in all 4 languages (DE;EN;FR;IT)
*/

type PaylinkBodyField struct {
	Active    bool `json:"active"`
	Mandatory bool `json:"mandatory"`
}

type PaylinkBodyFieldTranslatable struct {
	Names     shared.Translation `json:"names"`
	Active    bool               `json:"active"`
	Mandatory bool               `json:"mandatory"`
}

type PaylinkBodyFields struct {
	Title         PaylinkBodyField             `json:"title,omitempty"`
	Forename      PaylinkBodyField             `json:"forename,omitempty"`
	Surname       PaylinkBodyField             `json:"surname,omitempty"`
	Company       PaylinkBodyField             `json:"company,omitempty"`
	Street        PaylinkBodyField             `json:"street,omitempty"`
	Postcode      PaylinkBodyField             `json:"postcode,omitempty"`
	Place         PaylinkBodyField             `json:"place,omitempty"`
	Country       PaylinkBodyField             `json:"country,omitempty"`
	Phone         PaylinkBodyField             `json:"phone,omitempty"`
	Email         PaylinkBodyField             `json:"email,omitempty"`
	DateOfBirth   PaylinkBodyField             `json:"date_of_birth,omitempty"`
	Terms         PaylinkBodyField             `json:"terms,omitempty"`
	PrivacyPolicy PaylinkBodyField             `json:"privacyPolicy,omitempty"`
	CustomField1  PaylinkBodyFieldTranslatable `json:"custom_field_1,omitempty"`
	CustomField2  PaylinkBodyFieldTranslatable `json:"custom_field_2,omitempty"`
	CustomField3  PaylinkBodyFieldTranslatable `json:"custom_field_3,omitempty"`
}
