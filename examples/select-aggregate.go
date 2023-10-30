package main

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/clauses"
	"github.com/dwethmar/gosqle/clauses/groupby"
	"github.com/dwethmar/gosqle/clauses/orderby"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/functions"
)

// SelectAmountOfAddressesPerCountry select amount of addresses per country
func SelectAmountOfAddressesPerCountry() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		clauses.Selectable{
			Expr: &expressions.Column{Name: "country"},
		},
		clauses.Selectable{
			Expr: functions.NewCount(&expressions.Column{Name: "id"}),
			As:   "address_count",
		},
	).FromTable("addresses", nil).
		GroupBy(groupby.ColumnGrouping{
			&expressions.Column{Name: "country"},
		}).
		OrderBy(orderby.Sort{
			Column:    &expressions.Column{Name: "address_count"},
			Direction: orderby.DESC,
		}).Write(sb)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
