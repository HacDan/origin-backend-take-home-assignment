package api

import "net/http"

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) Server {
	return Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("POST /user/register", s.HandlePostRegisterUser)
	http.HandleFunc("POST /business/register", s.HandlePostRegisterBusiness)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) HandlePostRegisterUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandlePostRegisterBusiness(w http.ResponseWriter, r *http.Request) {

}
