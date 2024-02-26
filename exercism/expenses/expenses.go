package expenses

import "errors"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

func ByDaysPeriod(period DaysPeriod) func(Record) bool {
	return func(rec Record) bool {
		return period.From <= rec.Day && rec.Day <= period.To
	}
}

func ByCategory(category string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == category
	}
}

func TotalByPeriod(in []Record, period DaysPeriod) float64 {
	periodExpenses := Filter(in, ByDaysPeriod(period))
	var total float64
	for _, rec := range periodExpenses {
		total += rec.Amount
	}
	return total
}

func CategoryExpenses(in []Record, period DaysPeriod, cat string) (float64, error) {
	categoryExpenses := Filter(in, ByCategory(cat))
	if len(categoryExpenses) == 0 {
		return 0, errors.New("unknown category")
	}
	return TotalByPeriod(categoryExpenses, period), nil
}
