package filters

import "time"

type TimeRange struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}
type StringFilter struct {
	Eq        *string `json:"eq"`
	Contains  *string `json:"contains"`
	Neq       *string `json:"neq"`
	StartWith *string `json:"startWith"`
	EndWith   *string `json:"endWith"`
}

// type SortsFilter struct {
//	SortFilter SortFilter `json:"sortFilter"`
//}

// Number is a union type for all number types.
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type IntFilter struct {
	Eq  *int
	NEq *int
	Gt  *int
	Gte *int
	Lt  *int
	Lte *int
}

type FloatFilter struct {
	Eq  *float64
	NEq *float64
	Gt  *float64
	Gte *float64
	Lt  *float64
	Lte *float64
}
