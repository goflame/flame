package response

import "errors"

type Err struct {
	err    string
	status int
}

func NewError(err string) *Err {
	return &Err{err, 500}
}

func (r *Err) Status(s int) *Err {
	r.status = s
	return r
}

func (r *Err) GetError() error {
	if r == nil {
		return nil
	}
	return errors.New(r.err)
}

func (r *Err) GetStatus() int {
	return r.status
}
