package main

import (
	"log"
	"net/http"
)

type server struct{
	addr string 
}

func (s *server)ServeHTTP(w http.ResponseWriter,r *http.Request){
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
	s:=&server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr,s); err!= nil{
		log.Fatal(err)
	}
}