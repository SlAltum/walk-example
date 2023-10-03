package databinding

import (
	"time"
)

type DurationField struct {
	p *time.Duration
}

func (*DurationField) CanSet() bool       { return true }
func (f *DurationField) Get() interface{} { return f.p.String() }
func (f *DurationField) Set(v interface{}) error {
	x, err := time.ParseDuration(v.(string))
	if err == nil {
		*f.p = x
	}
	return err
}
func (f *DurationField) Zero() interface{} { return "" }
