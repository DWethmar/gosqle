package examples

import (
	"strings"

	"github.com/dwethmar/gosqle"
	"github.com/dwethmar/gosqle/alias"
	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/logic"
	"github.com/dwethmar/gosqle/predicates"
)

// WhereIsNull selects addresses where phone is null
func WhereIsNull() (string, error) {
	sb := new(strings.Builder)
	err := gosqle.NewSelect(
		alias.New(expressions.Column{Name: "id"}),
	).From(alias.NewStr("users")).
		Where(
			logic.And(predicates.IsNull(expressions.Column{Name: "phone"})),
		).Write(sb)

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
