package main

import (
	"net/http"
)

type api struct{
	addr string 
}

func (s *api)ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case http.MethodGet:
		switch r.URL.Path{
			case "/" :
				w.Write([]byte("index page"))
				return
			case "/user":
				w.Write([]byte("User page"))
				return 

			case "/main-page":
				w.Write([]byte("Main page"))
				return
		}
	default:
		w.Write([]byte("Error 404"))
		return
		

	}
       
}

func main(){
	api:=&api{addr: ":8080"}

	srv:=&http.Server{
		Addr: api.addr,
		Handler: api,
	}
	srv.ListenAndServe()
}