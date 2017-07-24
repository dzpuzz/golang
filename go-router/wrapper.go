package router

import "net/http"

type HttpRequest struct {
	*http.Request
	PathParams map[string]string
}

type HttpResponse struct {
	http.ResponseWriter
}

func newHttpRequest(r *http.Request) *HttpRequest {
	return &HttpRequest{r, make(map[string]string)}

}

func newHttpResponse(r http.ResponseWriter) HttpResponse {
	return HttpResponse{r}
}
