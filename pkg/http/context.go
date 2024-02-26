package http

type Context struct {
	Request *Request
	*Response
	store map[string]interface{}
}

func NewContext(request *Request, response *Response) *Context {
	return &Context{
		Request:  request,
		Response: response,
		store:    map[string]interface{}{},
	}
}

func (c *Context) Prop(name string, defaultValue ...string) string {
	val, ok := c.Request.Props[name]
	if ok {
		return val
	}
	if len(defaultValue) == 1 {
		return defaultValue[0]
	}
	return ""
}

func (c *Context) Set(key string, value interface{}) {
	c.store[key] = value
}

func (c *Context) Get(key string) interface{} {
	val, ok := c.store[key]
	if ok {
		return val
	}
	return nil
}
