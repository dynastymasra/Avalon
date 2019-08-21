package repository

type Query struct {
	Model     string
	Limit     int
	Offset    int
	Filters   []*Filter
	Orderings []*Ordering
}

type Filter struct {
	Condition string
	Property  string
	Value     interface{}
}

type Ordering struct {
	Property  string
	Direction string
}

const (
	FilterEqual            = "eq"
	FilterLessThan         = "lt"
	FilterGreaterThan      = "gt"
	FilterGreaterThanEqual = "gte"
	FilterLessThanEqual    = "lte"
	FilterJsonb            = "json"

	Descending = "desc"
	Ascending  = "asc"
)

var (
	validOrdering = map[string]bool{
		Descending: true,
		Ascending:  true,
	}
)

func NewQuery(model string) *Query {
	return &Query{
		Model: model,
	}
}

// Filter adds a filter to the query
func (q *Query) Filter(property, condition string, value interface{}) *Query {
	filter := NewFilter(property, condition, value)
	q.Filters = append(q.Filters, filter)
	return q
}

// Order adds a sort order to the query
func (q *Query) Ordering(property, direction string) *Query {
	order := NewOrdering(property, direction)
	q.Orderings = append(q.Orderings, order)
	return q
}

func (q *Query) Slice(offset, limit int) *Query {
	q.Offset = offset
	q.Limit = limit

	return q
}

// NewFilter creates a new property filter
func NewFilter(property, condition string, value interface{}) *Filter {
	return &Filter{
		Property:  property,
		Condition: condition,
		Value:     value,
	}
}

func NewOrdering(property, direction string) *Ordering {
	d := direction

	if !validOrdering[direction] {
		d = Descending
	}

	return &Ordering{
		Property:  property,
		Direction: d,
	}
}
