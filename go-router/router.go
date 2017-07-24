package router

import (
	"fmt"
	"net/http"
	"regexp"
)

//照抄gorest  修改版支持同一URL下的http方法
type Router struct {
	handlers   map[string]map[string]func(r *HttpRequest, w HttpResponse) error
	patterns   []string
	methods    map[string]map[string]string
	regexps    map[string]*regexp.Regexp
	pathparams map[string][]string
	errHandler func(err error, r *HttpRequest, w HttpResponse)
}

type hodler struct {
	router *Router
}

func newRouter() *Router {
	return &Router{
		handlers:   make(map[string]map[string]func(r *HttpRequest, w HttpResponse) error),
		patterns:   make([]string, 0),
		methods:    make(map[string]map[string]string),
		regexps:    make(map[string]*regexp.Regexp),
		pathparams: make(map[string][]string),
		errHandler: func(err error, r *HttpRequest, w HttpResponse) {
			w.Write([]byte(err.Error()))
		},
	}
}
func (r *Router) handle(method string, pattern string, handler func(r *HttpRequest, w HttpResponse) error) {
	if f, exist := r.handlers[pattern]; exist {
		f[method] = handler
	} else {
		f := make(map[string]func(r *HttpRequest, w HttpResponse) error)
		f[method] = handler
		r.handlers[pattern] = f
	}
	if c, exist := r.methods[pattern]; exist {
		c[method] = method
	} else {
		c := make(map[string]string)
		c[method] = method
		r.methods[pattern] = c
	}
	r.regexps[pattern], r.pathparams[pattern] = convertPatterntoRegex(pattern)
	for _, s := range r.patterns {
		if s == pattern {
			return
		}
	}
	r.patterns = append(r.patterns, pattern)

}

func (r *Router) Get(pattern string, handler func(r *HttpRequest, w HttpResponse) error) {
	r.handle("GET", pattern, handler)
}

func (r *Router) Post(pattern string, handler func(r *HttpRequest, w HttpResponse) error) {
	r.handle("POST", pattern, handler)
}

func (r *Router) Delete(pattern string, handler func(r *HttpRequest, w HttpResponse) error) {
	r.handle("DELETE", pattern, handler)
}

func (r *Router) Put(pattern string, handler func(r *HttpRequest, w HttpResponse) error) {
	r.handle("PUT", pattern, handler)
}

func (r *Router) Error(handler func(err error, r *HttpRequest, w HttpResponse)) {
	r.errHandler = handler
}

func (r *Router) Run(address string) error {
	fmt.Printf("Server listens on %s", address)
	err := http.ListenAndServe(address, &hodler{router: r})
	if err != nil {
		return err
	}
	return nil
}
func (h *hodler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := newHttpRequest(r)
	response := newHttpResponse(w)
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				h.router.errHandler(InternalError{e, ""}, request, response)
			}
			if e, ok := err.(string); ok {
				h.router.errHandler(InternalError{nil, e}, request, response)
			}
		}
	}()
	for _, p := range h.router.patterns {
		if reg, ok := h.router.regexps[p]; ok {
			if method, ok := h.router.methods[p]; ok && r.Method == method[r.Method] {
				if reg.Match([]byte(r.URL.Path)) {
					matchers := reg.FindSubmatch([]byte(r.URL.Path))
					pathParamMap := make(map[string]string)
					if len(matchers) > 1 {
						if pathParamNames, ok := h.router.pathparams[p]; ok {
							for i := 1; i < len(matchers); i++ {
								pathParamMap[pathParamNames[i]] = string(matchers[i])
							}
						}
					}
					request.PathParams = pathParamMap
					if handler, ok := h.router.handlers[p][method[r.Method]]; ok {
						err := handler(request, response)
						if err != nil {
							h.router.errHandler(err, request, response)
						}
						return
					}
				}
			}
		}
	}
}
func convertPatterntoRegex(pattern string) (*regexp.Regexp, []string) {
	b := regexp.MustCompile(`:[a-zA-Z0-9]+`).ReplaceAll([]byte(pattern), []byte(`([a-zA-Z1-9]+)`))
	reg := regexp.MustCompile("^" + string(b) + "$")
	b1 := regexp.MustCompile(`:[a-zA-Z0-9]+`).ReplaceAll([]byte(pattern), []byte(`:([a-zA-Z1-9]+)`))
	reg1 := regexp.MustCompile(string(b1))
	matchers := reg1.FindSubmatch([]byte(pattern))
	pathparamnames := make([]string, 0)
	if len(matchers) > 0 {
		for i := 0; i < len(matchers); i++ {
			pathparamnames = append(pathparamnames, string(matchers[i]))
		}
	}
	return reg, pathparamnames
}
