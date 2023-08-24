package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

// embed the addresses.json file
//
//go:embed addresses.json
var addressesFile []byte
var addresses []Address

// embed the address_types.json file
//
//go:embed address_types.json
var addressTypesFile []byte
var addressTypes []AddressType

type AddressType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	ID                  int    `json:"id"`
	UserID              int    `json:"user_id"`
	RecipientName       string `json:"recipient_name"`
	AddressLine1        string `json:"address_line1"`
	AddressLine2        string `json:"address_line2"`
	AddressLine3        string `json:"address_line3"`
	City                string `json:"city"`
	StateProvinceRegion string `json:"state_province_region"`
	PostalCode          string `json:"postal_code"`
	Country             string `json:"country"`
	Phone               string `json:"phone"`
	Type                string `json:"type"`
}

type User struct {
	ID    int64
	Name  string
	Email string
}

func init() {
	err := json.Unmarshal(addressesFile, &addresses)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(addressTypesFile, &addressTypes)
	if err != nil {
		log.Fatal(err)
	}
}
