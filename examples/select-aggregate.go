package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/functions"
)

// SelectAmountOfAddressesPerCountry select amount of addresses per country
func SelectAmountOfAddressesPerCountry() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "country"}),
		alias.New(functions.NewCount(&expressions.Column{Name: "id"})).SetAs("address_count"),
	).FromTable("addresses", nil).
		GroupBy(groupby.ColumnGrouping{
			&expressions.Column{Name: "country"},
		}).
		OrderBy(orderby.Sort{
			Column:    &expressions.Column{Name: "address_count"},
			Direction: orderby.DESC,
		}).
		Write(sb)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
