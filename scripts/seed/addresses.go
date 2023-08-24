package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func seedAddressTypes(txn *sql.Tx) error {
	if txn == nil {
		return fmt.Errorf("transaction is nil")
	}

	stmt, err := txn.Prepare(pq.CopyIn("address_type", "id", "type"))
	if err != nil {
		log.Fatal(err)
	}

	for i, t := range addressTypes {
		_, err = stmt.Exec(i+1, t)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetAddressTypes(txn *sql.Tx) ([]AddressType, error) {
	rows, err := txn.Query("SELECT id, name FROM address_types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addressTypes []AddressType
	for rows.Next() {
		var addressType AddressType
		err := rows.Scan(&addressType.ID, &addressType.Name)
		if err != nil {
			return nil, err
		}

		addressTypes = append(addressTypes, addressType)
	}

	return addressTypes, nil
}

func seedAddresses(txn *sql.Tx) error {
	addressTypes, err := GetAddressTypes(txn)
	if err != nil {
		return err
	}

	if txn == nil {
		return fmt.Errorf("transaction is nil")
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
		log.Fatal(err)
	}

	for i, a := range addresses {
		var addressTypeID int
		for _, t := range addressTypes {
			if t.Name == a.Type {
				addressTypeID = t.ID
				break
			}
		}

		_, err = stmt.Exec(
			i+1,
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
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
