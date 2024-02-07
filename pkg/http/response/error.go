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
	return nil
}

func (r *Err) GetStatus() int {
	return r.status
}