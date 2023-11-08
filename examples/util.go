package examples

import "database/sql"

func InsertAddress(db *sql.DB, address *Address) (int64, error) {
	r := db.QueryRow(
		`INSERT INTO addresses (
			user_id,
			recipient_name,
			address_line1,
			address_line2,
			address_line3,
			city,
			state_province_region,
			postal_code,
			country,
			phone,
			address_type_id
		) VALUES (
			(SELECT id FROM users ORDER BY id ASC LIMIT 1),
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			(SELECT id FROM address_types ORDER BY id ASC LIMIT 1)
		) RETURNING id;`,
		address.RecipientName,
		address.AddressLine1,
		address.AddressLine2,
		address.AddressLine3,
		address.City,
		address.StateProvinceRegion,
		address.PostalCode,
		address.Country,
		address.Phone,
	)

	var id int64
	if err := r.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
