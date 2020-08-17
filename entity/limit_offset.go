package entity

// LimitOffset definition
type LimitOffset struct {
	Limit  int `json:"limit" query:"limit"`
	Offset int `json:"offset"  query:"offset"`
}
