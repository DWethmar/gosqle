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
	/**
	SELECT country, COUNT(id) AS address_count
	FROM addresses
	GROUP BY country
	ORDER BY address_count DESC;
	**/
	err := gosqle.NewSelect(
		clauses.Selectable{Expr: expressions.NewColumn("country")},
		clauses.Selectable{Expr: expressions.NewCount(expressions.NewColumn("id")), As: "address_count"},
	).From(expressions.Table{
		Name: "addresses",
	}).GroupBy(groupby.ColumnGrouping{
		expressions.NewColumn("country"),
	}).OrderBy(orderby.Sort{
		Column:    expressions.NewColumn("address_count"),
		Direction: orderby.DESC,
	}).WriteTo(sb)
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
