package transaction

import (
	"github.com/SleepyMorpheus/payrexx-sdk-go/types/shared"
	"strconv"
	"time"
)

type RetrieveManyArguments struct {
	GreaterThan        time.Time
	LessThan           time.Time
	MyTransactionsOnly bool
	OrderByTime        shared.SortOrder
	Offset             int32
	Limit              int32
}

func (args *RetrieveManyArguments) ToMap() map[string]string {
	var argMap = make(map[string]string)

	if !args.GreaterThan.IsZero() {
		argMap["filterDatetimeUtcGreaterThan"] = args.GreaterThan.Format("2006-01-02 15:04:05")
	}

	if !args.LessThan.IsZero() {
		argMap["filterDatetimeUtcLessThan"] = args.LessThan.Format("2006-01-02 15:04:05")
	}

	// same behaviour if set or not
	argMap["filterMyTransactionsOnly"] = strconv.FormatBool(args.MyTransactionsOnly)

	if args.OrderByTime != "" {
		argMap["orderByTime"] = string(args.OrderByTime)
	} else {
		argMap["orderByTime"] = string(shared.SortOrderAsc)
	}

	// same behaviour if set or not
	argMap["offset"] = string(args.Offset)

	if args.Limit > 0 {
		argMap["limit"] = string(args.Limit)
	}

	return argMap
}
