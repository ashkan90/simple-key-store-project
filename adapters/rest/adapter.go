package rest

import (
	"fmt"
	"log"
	"net/http"
	"ys-project/adapters/handlers"
	"ys-project/domain/key"
)

type Adapter struct {
	Port string
	Host string
}

func (a *Adapter) Serve() {
	var keyRepository, err = key.NewFileRepository("/tmp/TIMESTAMP-data.json")
	if err != nil {
		panic(err)
	}
	var keyService = key.NewService(keyRepository, 1)
	var keyRest = &handlers.KeyHandler{Service: keyService}

	RegisterRequests("/key", keyRest)
	//RegisterRequests("/foo", fooRest)
	//RegisterRequests("/baz", bazRest)
	//RegisterRequests("/some", someRest)

	log.Println("Server has started on port" + a.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", a.Host, a.Port), nil))
}

func RegisterRequests(basePath string, handler http.Handler) {
	http.Handle(basePath, handler)
}
