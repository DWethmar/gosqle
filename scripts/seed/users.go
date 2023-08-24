package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

func seedUsers(txn *sql.Tx) error {
	if txn == nil {
		return fmt.Errorf("transaction is nil")
	}

	emailDomains, err := readLines("scripts/seed/email_domain.txt")
	if err != nil {
		return fmt.Errorf("error reading email domains: %w", err)
	}

	emailDomainsMap := make(map[string]string)
	// for every letter after the @ we create a map
	for _, domain := range emailDomains {
		emailDomainsMap[domain[1:2]] = domain
	}

	stmt, err := txn.Prepare(pq.CopyIn("users", "id", "name", "email"))
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	for i, a := range addresses {
		email := strings.ToLower(a.RecipientName)
		email = strings.ReplaceAll(email, " ", ".")
		email = fmt.Sprintf("%s%s", email, emailDomainsMap[email[0:1]])

		_, err = stmt.Exec(i+1, strings.ToLower(a.RecipientName), email)
		if err != nil {
			return fmt.Errorf("error executing statement: %w", err)
		}
	}

	if err = stmt.Close(); err != nil {
		return fmt.Errorf("error closing statement: %w", err)
	}

	return nil
}
