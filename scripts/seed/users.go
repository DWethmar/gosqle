package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/lib/pq"
)

func seedUsers(txn *sql.Tx) error {
	if txn == nil {
		return fmt.Errorf("transaction is nil")
	}

	emailDomains, err := readLines("scripts/seed/email_domain.txt")
	if err != nil {
		return err
	}

	emailDomainsMap := make(map[string]string)
	// for every letter after the @ we create a map
	for _, domain := range emailDomains {
		emailDomainsMap[domain[1:2]] = domain
	}

	stmt, err := txn.Prepare(pq.CopyIn("users", "id", "name", "email"))
	if err != nil {
		log.Fatal(err)
	}

	for i, a := range addresses {
		email := strings.ToLower(a.RecipientName)
		email = strings.ReplaceAll(email, " ", ".")
		email = fmt.Sprintf("%s%s", email, emailDomainsMap[email[0:1]])

		_, err = stmt.Exec(i+1, strings.ToLower(a.RecipientName), email)
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
