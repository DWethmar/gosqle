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
	Write(writer io.StringWriter) error
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
func (f Wrap) Write(sw io.StringWriter) error {
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
func (f EQ) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing EQ column: %v", err)
	}

	if _, err := writer.WriteString(" = "); err != nil {
		return fmt.Errorf("error writing EQ operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f NE) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing NE column: %v", err)
	}

	if _, err := writer.WriteString(" != "); err != nil {
		return fmt.Errorf("error writing NE operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f GT) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing GT column: %v", err)
	}

	if _, err := writer.WriteString(" > "); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f GTE) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing GTE column: %v", err)
	}

	if _, err := writer.WriteString(" >= "); err != nil {
		return fmt.Errorf("error writing GTE operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f LT) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing LT column: %v", err)
	}

	if _, err := writer.WriteString(" < "); err != nil {
		return fmt.Errorf("error writing LT operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f LTE) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing LTE column: %v", err)
	}

	if _, err := writer.WriteString(" <= "); err != nil {
		return fmt.Errorf("error writing LTE operator")
	}

	if err := f.Expr.Write(writer); err != nil {
		return fmt.Errorf("error writing LTE argument: %v", err)
	}

	return nil
}

// In is an IN condition.
// For example: WHERE col IN (1, 2, 3)
// use expressions.NewList to create the list.
type In struct {
	Col   expressions.Expression
	Expr  expressions.Expression
	Logic Logic
}

func (f In) LogicOp() Logic { return f.Logic }
func (f In) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing IN column: %v", err)
	}

	if _, err := writer.WriteString(" IN ("); err != nil {
		return fmt.Errorf("error writing IN operator")
	}

	if err := f.Expr.Write(writer); err != nil {
		return fmt.Errorf("error writing IN expression: %v", err)
	}

	if _, err := writer.WriteString(")"); err != nil {
		return fmt.Errorf("error writing IN operator")
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
func (f Like) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing LIKE column: %v", err)
	}

	if _, err := writer.WriteString(" LIKE "); err != nil {
		return fmt.Errorf("error writing GT operator")
	}

	if err := f.Expr.Write(writer); err != nil {
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
func (f IsNull) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
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
func (f Between) Write(writer io.StringWriter) error {
	if err := f.Col.Write(writer); err != nil {
		return fmt.Errorf("error writing column: %v", err)
	}

	if _, err := writer.WriteString(" BETWEEN "); err != nil {
		return fmt.Errorf("error writing BETWEEN operator")
	}

	if err := f.Low.Write(writer); err != nil {
		return fmt.Errorf("error writing BETWEEN low argument: %v", err)
	}

	if _, err := writer.WriteString(" AND "); err != nil {
		return fmt.Errorf("error writing BETWEEN operator")
	}

	if err := f.High.Write(writer); err != nil {
		return fmt.Errorf("error writing BETWEEN high argument: %v", err)
	}

	return nil
}

// Not is a NOT condition.
type Not struct {
	Predicate Predicate
}

func (f Not) LogicOp() Logic { return f.Predicate.LogicOp() }
func (f Not) Write(writer io.StringWriter) error {
	if err := util.WriteStrings(writer, "NOT "); err != nil {
		return fmt.Errorf("error writing NOT opening: %v", err)
	}

	if err := f.Predicate.Write(writer); err != nil {
		return fmt.Errorf("error writing predicate after not: %v", err)
	}

	return nil
}
