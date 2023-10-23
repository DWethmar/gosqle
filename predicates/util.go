package predicates

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/util"
)

// WriteAll writes predicates to a string writer.
// It will ignore the logic operator of the first predicate.
func WriteAll(sw io.StringWriter, predicates []Predicate) error {
	for i, predicate := range predicates {
		if i != 0 {
			if err := util.WriteStrings(sw, " ", logicStrings[predicate.LogicOp()], " "); err != nil {
				return fmt.Errorf("error writing logic operator: %v", err)
			}
		}

		if err := predicate.Write(sw); err != nil {
			return fmt.Errorf("writing predicate failed: %w", err)
		}
	}

	return nil
}
