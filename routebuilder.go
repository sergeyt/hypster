package hypster

// RouteBuilder holds route handlers
type RouteBuilder struct {
	app *AppBuilder
	// handlers
	get    func(*Context)
	post   func(*Context)
	put    func(*Context)
	update func(*Context)
	patch  func(*Context)
	del    func(*Context)
}

// Get registers GET handler
func (r *RouteBuilder) Get(handler Handler) *RouteBuilder {
	r.get = handler
	return r
}

// Post registers POST handler
func (r *RouteBuilder) Post(handler Handler) *RouteBuilder {
	r.post = handler
	return r
}

// Put registers PUT handler
func (r *RouteBuilder) Put(handler Handler) *RouteBuilder {
	r.put = handler
	return r
}

// Update registers UPDATE handler
func (r *RouteBuilder) Update(handler Handler) *RouteBuilder {
	r.update = handler
	return r
}

// Patch registers PATCH handler
func (r *RouteBuilder) Patch(handler Handler) *RouteBuilder {
	r.patch = handler
	return r
}

// Delete registers DELETE handler
func (r *RouteBuilder) Delete(handler Handler) *RouteBuilder {
	r.del = handler
	return r
}
