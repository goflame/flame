package response

type Err struct {
	err    error
	status int
}

func NewError(err error) *Err {
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
	return r.err
}

func (r *Err) GetStatus() int {
	return r.status
}
