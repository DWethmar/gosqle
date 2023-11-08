package examples

import "database/sql"

type User struct {
	ID    int64
	Name  string
	Email string
}

type AmountOfAddressesPerCountry struct {
	Country      string
	AddressCount int
}

type Address struct {
	ID                  int
	UserID              int
	RecipientName       string
	AddressLine1        string
	AddressLine2        sql.NullString
	AddressLine3        sql.NullString
	City                string
	StateProvinceRegion sql.NullString
	PostalCode          sql.NullString
	Country             string
	Phone               sql.NullString
	AddressTypeID       int
}
