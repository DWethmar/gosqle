package predicates

import (
	"errors"
	"strings"
	"testing"

	"github.com/dwethmar/gosqle/expressions"
	"github.com/dwethmar/gosqle/mock"
	"github.com/dwethmar/gosqle/postgres"
)

func TestWrap(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := Wrap{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("Group.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := Wrap{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("Group.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := Wrap{
			Predicates: []Predicate{
				EQ{
					Col:  expressions.NewColumn("id"),
					Expr: postgres.NewArgument("123", 1),
				},
				EQ{
					Col:  expressions.NewColumn("name"),
					Expr: postgres.NewArgument("John", 2),
				},
			},
		}
		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("Group.Write() error = %v", err)
		}
		got := sb.String()
		want := `(id = $1 AND name = $2)`
		if got != want {
			t.Errorf("Group.Write() = %v, want %v", got, want)
		}
	})
}

func TestEQ(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := EQ{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("EQ.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := EQ{Col: expressions.NewColumn("name"), Expr: postgres.NewArgument("dennis", 0), Logic: OR}

		if got := f.LogicOp(); got != OR {
			t.Errorf("EQ.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := EQ{
			Col:   expressions.NewColumn("id"),
			Expr:  postgres.NewArgument("123", 1),
			Logic: OR,
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("EQ.Write() error = %v", err)
		}
		got := sb.String()
		want := `id = $1`
		if got != want {
			t.Errorf("EQ.Write() = %v, want %v", got, want)
		}
	})
}

func TestNE(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := NE{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("NE.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := NE{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("NE.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := NE{
			Col:   expressions.NewColumn("id"),
			Expr:  postgres.NewArgument("123", 1),
			Logic: OR,
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("NE.Write() error = %v", err)
		}

		got := sb.String()
		want := `id != $1`
		if got != want {
			t.Errorf("NE.Write() = %v, want %v", got, want)
		}
	})
}

func TestGT(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := GT{Logic: AND}

		if got := f.LogicOp(); got != AND {
			t.Errorf("GT.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := GT{Logic: OR}

		if got := f.LogicOp(); got != OR {
			t.Errorf("GT.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := GT{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("123", 1),
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("GT.Write() error = %v", err)
		}
		got := sb.String()
		want := `id > $1`
		if got != want {
			t.Errorf("GT.Write() = %v, want %v", got, want)
		}
	})
}

func TestGTE(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := GTE{Logic: AND}

		if got := f.LogicOp(); got != AND {
			t.Errorf("GTE.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := GTE{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("GTE.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := GTE{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("123", 1),
		}
		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("GTE.Write() error = %v", err)

		}
		got := sb.String()
		want := `id >= $1`
		if got != want {
			t.Errorf("GTE.Write() = %v, want %v", got, want)
		}
	})
}

func TestLT(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := LT{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("123", 0),
		}
		if got := f.LogicOp(); got != AND {
			t.Errorf("LT.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := LT{
			Col:   expressions.NewColumn("id"),
			Expr:  postgres.NewArgument("123", 0),
			Logic: OR,
		}
		if got := f.LogicOp(); got != OR {
			t.Errorf("LT.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := LT{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("123", 1),
		}

		sb := new(strings.Builder)

		if err := f.WriteTo(sb); err != nil {
			t.Errorf("LT.Write() error = %v", err)
		}

		got := sb.String()
		want := `id < $1`
		if got != want {
			t.Errorf("LT.Write() = %q, want %q", got, want)
		}
	})
}

func TestLTE(t *testing.T) {
	t.Parallel()

	t.Run("LogicAnd", func(t *testing.T) {
		f := LTE{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("LTE.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := LTE{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("LTE.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := LTE{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("123", 1),
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("LTE.Write() error = %v", err)
		}

		got := sb.String()
		want := `id <= $1`

		if got != want {
			t.Errorf("LTE.Write() = %q, want %q", got, want)
		}
	})
}

func TestIn(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := In{
			Logic: AND,
		}
		if got := f.LogicOp(); got != AND {
			t.Errorf("IN.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := In{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("IN.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := In{
			Col: expressions.NewColumn("id"),
			Expressions: []expressions.Expression{
				postgres.NewArgument("123", 1),
				postgres.NewArgument("456", 2),
			},
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("IN.Write() error = %v", err)
		}

		got := sb.String()
		want := `id IN ($1, $2)`

		if got != want {
			t.Errorf("IN.Write() = %v, want %v", got, want)
		}
	})
}

func TestLike(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := Like{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("Like.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := Like{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("Like.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := Like{
			Col:  expressions.NewColumn("id"),
			Expr: postgres.NewArgument("%123", 1),
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("Like.WriteTo() error = %v", err)
		}
		got := sb.String()
		want := `id LIKE $1`
		if got != want {
			t.Errorf("Like.WriteTo() = %v, want %v", got, want)
		}
	})
}

func TestIsNull(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := IsNull{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("IsNull.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := IsNull{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("IsNull.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := IsNull{
			Col: expressions.NewColumn("id"),
		}

		sb := new(strings.Builder)

		if err := f.WriteTo(sb); err != nil {
			t.Errorf("IsNull.Write() error = %v", err)
		}
		got := sb.String()
		want := `id IS NULL`
		if got != want {
			t.Errorf("IsNull.Write() = %v, want %v", got, want)
		}
	})
}

func TestBetween(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := Between{Logic: AND}
		if got := f.LogicOp(); got != AND {
			t.Errorf("Between.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := Between{Logic: OR}
		if got := f.LogicOp(); got != OR {
			t.Errorf("Between.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		f := Between{
			Col:  expressions.NewColumn("id"),
			Low:  postgres.NewArgument("123", 1),
			High: postgres.NewArgument("456", 2),
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("Between.Write() error = %v", err)
		}
		got := sb.String()
		want := `id BETWEEN $1 AND $2`
		if got != want {
			t.Errorf("Between.Write() = %v, want %v", got, want)
		}
	})
}

func TestNot(t *testing.T) {
	t.Run("LogicAnd", func(t *testing.T) {
		f := Not{
			Predicate: EQ{
				Expr: postgres.NewArgument("whoohoo", 0),
			},
		}
		if got := f.LogicOp(); got != AND {
			t.Errorf("Not.LogicOp() = %v, want %v", got, AND)
		}
	})

	t.Run("LogicOr", func(t *testing.T) {
		f := Not{
			Predicate: EQ{
				Expr:  postgres.NewArgument("whoohoo", 0),
				Logic: OR,
			},
		}
		if got := f.LogicOp(); got != OR {
			t.Errorf("Not.LogicOp() = %v, want %v", got, OR)
		}
	})

	t.Run("write to sql", func(t *testing.T) {
		arg := postgres.NewArgument("123", 1)

		f := Not{
			Predicate: EQ{
				Col:  expressions.NewColumn("id"),
				Expr: arg,
			},
		}

		sb := new(strings.Builder)
		if err := f.WriteTo(sb); err != nil {
			t.Errorf("Not.Write() error = %v", err)
		}

		got := sb.String()
		want := `NOT id = $1`

		if got != want {
			t.Errorf("Not.Write() = %q, want %q", got, want)
		}
	})

	t.Run("should return io.writer error", func(t *testing.T) {
		f := Not{
			Predicate: EQ{
				Col:  expressions.NewColumn("id"),
				Expr: postgres.NewArgument("123", 0),
			},
		}

		sb := mock.StringWriterFn(func(s string) (n int, err error) {
			return 0, errors.New("error")
		})

		if err := f.WriteTo(sb); err == nil {
			t.Errorf("Not.Write() error = %v", err)
		}
	})
}
