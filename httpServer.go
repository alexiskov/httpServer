package httpserver

import (
	"errors"
	"net/http"
)

type (
	server struct {
		*http.Server
		*Pipe
	}
	Pipe struct {
		ErorLog *chan error
	}
)

func New(pipe *Pipe) *server {
	return &server{&http.Server{Addr: ":8080"}, pipe}
}

func (srv *server) Run() {
	http.HandleFunc("/", srv.handlerMain)
}

func (srv *server) handlerMain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		srv.methodGET()
	case http.MethodPost:
		srv.methodPOST()
	case http.MethodPut:
		srv.methodPUT()
	case http.MethodDelete:
		srv.methodDELETE()
	default:
		http.Error(w, "HTTP-method is not allowed", http.StatusMethodNotAllowed)
		*srv.Pipe.ErorLog <- errors.New("http-method is not allowed")
	}
}

func (srv *server) methodGET()    {}
func (srv *server) methodPOST()   {}
func (srv *server) methodPUT()    {}
func (srv *server) methodDELETE() {}
