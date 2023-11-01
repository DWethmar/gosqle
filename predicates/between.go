package predicates

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
)

type Range struct {
	Target expressions.Expression
	Low    expressions.Expression
	High   expressions.Expression
}

func (r *Range) Write(writer io.StringWriter) error {
	if r.Target == nil {
		return fmt.Errorf("range target is nil")
	}

	if r.Low == nil {
		return fmt.Errorf("range low is nil")
	}

	if r.High == nil {
		return fmt.Errorf("range high is nil")
	}

	if err := r.Target.Write(writer); err != nil {
		return fmt.Errorf("error writing range target: %v", err)
	}

	if _, err := writer.WriteString(" BETWEEN "); err != nil {
		return fmt.Errorf("error writing BETWEEN: %v", err)
	}

	if err := r.Low.Write(writer); err != nil {
		return fmt.Errorf("error writing range low: %v", err)
	}

	if _, err := writer.WriteString(" AND "); err != nil {
		return fmt.Errorf("error writing AND: %v", err)
	}

	if err := r.High.Write(writer); err != nil {
		return fmt.Errorf("error writing range high: %v", err)
	}

	return nil
}

func Between(target expressions.Expression, low expressions.Expression, high expressions.Expression) *Range {
	return &Range{
		Target: target,
		Low:    low,
		High:   high,
	}
}
