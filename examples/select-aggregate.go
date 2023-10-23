package main

import (
	"database/sql"
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/postgres"
)

type AmountOfAddressesPerCountry struct {
	Country string
	Count   int64
}

// SelectAmountOfAddressesPerCountry select amount of addresses per country
func SelectAmountOfAddressesPerCountry(db *sql.DB) ([]AmountOfAddressesPerCountry, string, error) {
	sb := new(strings.Builder)
	args := postgres.NewArguments()
	err := gosqle.NewSelect(
		clauses.Selectable{
			Expr: &expressions.Column{Name: "country"},
		},
		clauses.Selectable{
			Expr: expressions.NewCount(&expressions.Column{Name: "id"}),
			As:   "address_count",
		},
	).FromTable("addresses", nil).GroupBy(groupby.ColumnGrouping{
		&expressions.Column{Name: "country"},
	}).OrderBy(orderby.Sort{
		Column:    &expressions.Column{Name: "address_count"},
		Direction: orderby.DESC,
	}).Write(sb)
	if err != nil {
		return nil, "", err
	}

	rows, err := db.Query(sb.String(), args.Args...)
	if err != nil {
		return nil, "", err
	}

	var r []AmountOfAddressesPerCountry
	for rows.Next() {
		var a AmountOfAddressesPerCountry
		err = rows.Scan(&a.Country, &a.Count)
		if err != nil {
			return nil, "", err
		}
		r = append(r, a)
	}

	return r, sb.String(), nil
}
