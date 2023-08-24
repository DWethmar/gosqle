package predicates

import (
	"fmt"
	"io"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/util"
)

var logicStrings = map[Logic]string{
	AND: "AND",
	OR:  "OR",
}

var ( // type check
	_ Predicate = EQ{}
	_ Predicate = NE{}
	_ Predicate = GT{}
	_ Predicate = GTE{}
	_ Predicate = LT{}
	_ Predicate = LTE{}
	_ Predicate = In{}
	_ Predicate = Like{}
	_ Predicate = IsNull{}
	_ Predicate = Between{}
	_ Predicate = Wrap{}
	_ Predicate = Not{}
)

// Predicates are conditions that can be evaluated to SQL three-valued logic (3VL) (true/false/unknown)
type Predicate interface {
	LogicOp() Logic // LogicOp returns the logic operator: AND or OR.
	WriteTo(writer io.StringWriter) error
}

// Logic is a logic operator.
type Logic int

const (
	AND Logic = iota // LogicAnd is the AND operator.
	OR               // LogicOr is the OR operator.
)

// Wrap wraps a group of predicates in parenthesis.
type Wrap struct {
	Predicates []Predicate
	Logic      Logic
}

func (f Wrap) LogicOp() Logic { return f.Logic }

// Write writes the group to the given writer.
func (f Wrap) WriteTo(sw io.StringWriter) error {
	if err := util.WriteStrings(sw, "("); err != nil {
		return fmt.Errorf("error writing opening parenthesis: %v", err)
	}

	if err := WriteAll(sw, f.Predicates); err != nil {
		return fmt.Errorf("error writing predicates: %v", err)
	}

	if err := util.WriteStrings(sw, ")"); err != nil {
		return fmt.Errorf("error writing closing parenthesis: %v", err)
	}

	return nil
}

// EQ is an equal condition.
type EQ struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f EQ) LogicOp() Logic { return f.Logic }
func (f EQ) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing EQ column: %v", err)
	}

	if _, err := writer.WriteString(" = "); err != nil {
		return fmt.Errorf("error writing EQ operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing EQ argument: %v", err)
	}

	return nil
}

// NE is a not equal condition.
type NE struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f NE) LogicOp() Logic { return f.Logic }
func (f NE) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing NE column: %v", err)
	}

	if _, err := writer.WriteString(" != "); err != nil {
		return fmt.Errorf("error writing NE operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing NE argument: %v", err)
	}

	return nil
}

// GT is a greater than condition.
type GT struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f GT) LogicOp() Logic { return f.Logic }
func (f GT) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing GT column: %v", err)
	}

	if _, err := writer.WriteString(" > "); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing GT argument: %v", err)
	}

	return nil
}

// GTE is a greater than or equal to condition.
type GTE struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f GTE) LogicOp() Logic { return f.Logic }
func (f GTE) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing GTE column: %v", err)
	}

	if _, err := writer.WriteString(" >= "); err != nil {
		return fmt.Errorf("error writing GTE operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing GTE argument: %v", err)
	}

	return nil
}

// LT is a less than condition.
type LT struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f LT) LogicOp() Logic { return f.Logic }
func (f LT) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LT column: %v", err)
	}

	if _, err := writer.WriteString(" < "); err != nil {
		return fmt.Errorf("error writing LT operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LT argument: %v", err)
	}

	return nil
}

// LTE is a less than or equal to condition.
type LTE struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f LTE) LogicOp() Logic { return f.Logic }
func (f LTE) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LTE column: %v", err)
	}

	if _, err := writer.WriteString(" <= "); err != nil {
		return fmt.Errorf("error writing LTE operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LTE argument: %v", err)
	}

	return nil
}

// In is an IN condition.
type In struct {
	Col         expressions.Expression
	Expressions []expressions.Expression
	Logic       Logic
}

func (f In) LogicOp() Logic { return f.Logic }
func (f In) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing GT column: %v", err)
	}

	if _, err := writer.WriteString(" IN ("); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	for i, argument := range f.Expressions {
		if err := argument.WriteTo(writer); err != nil {
			return fmt.Errorf("could not write argument at index %d: %w", i, err)
		}

		if i < len(f.Expressions)-1 {
			if _, err := writer.WriteString(", "); err != nil {
				return fmt.Errorf("error writing comma: %w", err)
			}
		}
	}

	if _, err := writer.WriteString(")"); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	return nil
}

// Like is a LIKE condition.
type Like struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f Like) LogicOp() Logic { return f.Logic }
func (f Like) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LIKE column: %v", err)
	}

	if _, err := writer.WriteString(" LIKE "); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	if err := f.Expr.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing LIKE argument: %v", err)
	}
	return nil
}

// IsNull is a IS NULL condition.
type IsNull struct {
	Col   expressions.Expression
	Logic Logic
}

func (f IsNull) LogicOp() Logic { return f.Logic }
func (f IsNull) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing IsNull field: %v", err)
	}

	if err := util.WriteStrings(writer, " IS NULL"); err != nil {
		return fmt.Errorf("error writing IS NULL predicate: %v", err)
	}

	return nil
}

type Between struct {
	Col   expressions.Expression
	Low   expressions.Expression
	High  expressions.Expression
	Logic Logic
}

func (f Between) LogicOp() Logic { return f.Logic }
func (f Between) WriteTo(writer io.StringWriter) error {
	if err := f.Col.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing column: %v", err)
	}

	if _, err := writer.WriteString(" BETWEEN "); err != nil {
		return fmt.Errorf("error writing BETWEEN operator")
	}

	if err := f.Low.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing BETWEEN low argument: %v", err)
	}

	if _, err := writer.WriteString(" AND "); err != nil {
		return fmt.Errorf("error writing BETWEEN operator")
	}

	if err := f.High.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing BETWEEN high argument: %v", err)
	}

	return nil
}

// Not is a NOT condition.
type Not struct {
	Predicate Predicate
}

func (f Not) LogicOp() Logic { return f.Predicate.LogicOp() }
func (f Not) WriteTo(writer io.StringWriter) error {
	if err := util.WriteStrings(writer, "NOT "); err != nil {
		return fmt.Errorf("error writing NOT opening: %v", err)
	}

	if err := f.Predicate.WriteTo(writer); err != nil {
		return fmt.Errorf("error writing predicate after not: %v", err)
	}

	return nil
}
