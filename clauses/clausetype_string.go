// Code generated by "stringer -type=ClauseType"; DO NOT EDIT.

package clauses

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownType-0]
	_ = x[FromType-1]
	_ = x[JoinType-2]
	_ = x[WhereType-3]
	_ = x[LimitType-4]
	_ = x[OffsetType-5]
	_ = x[OrderByType-6]
	_ = x[ReturningType-7]
	_ = x[ValuesType-8]
	_ = x[SetType-9]
}

const _ClauseType_name = "UnknownClauseFromClauseJoinClauseWhereClauseLimitClauseOffsetClauseOrderByClauseReturningClauseValuesClauseSetClause"

var _ClauseType_index = [...]uint8{0, 13, 23, 33, 44, 55, 67, 80, 95, 107, 116}

func (i ClauseType) String() string {
	if i < 0 || i >= ClauseType(len(_ClauseType_index)-1) {
		return "ClauseType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ClauseType_name[_ClauseType_index[i]:_ClauseType_index[i+1]]
}
