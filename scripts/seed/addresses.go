package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

func seedAddressTypes(txn *sql.Tx) error {
	if txn == nil {
		return fmt.Errorf("transaction is nil")
	}

	stmt, err := txn.Prepare(pq.CopyIn("address_types", "name"))
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	for _, t := range addressTypes {
		_, err = stmt.Exec(t.Name)
		if err != nil {
			return fmt.Errorf("error executing statement: %w", err)
		}
	}

	if err = stmt.Close(); err != nil {
		return fmt.Errorf("error closing statement: %w", err)
	}

	return nil
}

func GetAddressTypes(txn *sql.Tx) ([]AddressType, error) {
	rows, err := txn.Query("SELECT id, name FROM address_types")
	if err != nil {
		return nil, fmt.Errorf("error querying address_types: %w", err)
	}
	defer rows.Close()

	var addressTypes []AddressType
	for rows.Next() {
		var addressType AddressType
		err := rows.Scan(&addressType.ID, &addressType.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning address_types: %w", err)
		}

		addressTypes = append(addressTypes, addressType)
	}

	return addressTypes, nil
}

func seedAddresses(txn *sql.Tx) error {
	addressTypes, err := GetAddressTypes(txn)
	if err != nil {
		return fmt.Errorf("error getting address types: %w", err)
	}

	if txn == nil {
		return fmt.Errorf("transaction is nil")
	}

	users, err := getUsers(txn)
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"addresses",
		"user_id",
		"recipient_name",
		"address_line1",
		"address_line2",
		"address_line3",
		"city",
		"state_province_region",
		"postal_code",
		"country",
		"phone",
		"address_type_id",
	))
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	for i, a := range addresses {
		var addressTypeID int
		for _, t := range addressTypes {
			if t.Name == a.Type {
				addressTypeID = t.ID
				break
			}
		}

		var user User
		if i < len(users) {
			user = users[i]
		} else {
			user = users[0]
		}

		_, err = stmt.Exec(
			user.ID,
			a.RecipientName,
			a.AddressLine1,
			a.AddressLine2,
			a.AddressLine3,
			a.City,
			a.StateProvinceRegion,
			a.PostalCode,
			a.Country,
			a.Phone,
			addressTypeID,
		)
		if err != nil {
			return fmt.Errorf("error executing statement: %w", err)
		}
	}

	err = stmt.Close()
	if err != nil {
		return fmt.Errorf("error closing statement: %w", err)
	}

	return nil
}
