package httpserver

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	server struct {
		*http.Server
		*Pipe
	}
	Pipe struct {
		ErorLog *chan error //канал вывода лога ошибок из пакета
	}
)

// принимает аргументом указатель на структуру содержащую каналы
// возвращает указатель на структуру данных server
func New(pipe *Pipe) *server {
	return &server{&http.Server{Addr: ":8080"}, pipe}
}

// метод запускает сервер
// если возникает ошибка - отправляет в канал ErorLog и завершает работу
func (srv *server) Run() {
	http.HandleFunc("/", srv.handlerMain)
	if err := srv.ListenAndServe(); err != nil {
		err = fmt.Errorf("http server runing error: %w", err)
		*srv.Pipe.ErorLog <- err
		return
	}
}

// базовый хендлер
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

// обработчики методов
func (srv *server) methodGET()    {}
func (srv *server) methodPOST()   {}
func (srv *server) methodPUT()    {}
func (srv *server) methodDELETE() {}
