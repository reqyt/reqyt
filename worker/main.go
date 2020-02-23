package main

import (
	" reqyt.run/worker/container_manager"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Handler struct {
	containerManager *container_manager.ContainerManager
	config           Config
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	image := req.URL.String()[1:]
	config := handler.config.Images.ForImage(image)
	if config == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Image not registered!"))
		return
	}
	container := handler.containerManager.GetContainer(image)
	url := "http://"+ container.IP.String()+":8000"
	proxyReq, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		// handle error
	}

	proxyReq.Header.Set("Host", req.Host)
	proxyReq.Header.Set("X-Forwarded-For", req.RemoteAddr)

	for header, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{}
	proxyRes, err := client.Do(proxyReq)

	for header, values := range proxyRes.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}

	io.Copy(w, proxyRes.Body)
}

func main() {
	config := readConfig()
	cm, err := container_manager.NewContainerManager(config.Images)
	if err != nil {
		panic(err)
	}
	defer cm.Close()
	server := &http.Server{Addr: ":8000", Handler: &Handler{cm, config}}
	go func() {
		c := make(chan os.Signal, 1)

		// Passing no signals to Notify means that
		// all signals will be sent to the channel.
		signal.Notify(c, os.Interrupt)

		// Block until any signal is received.
		s := <-c
		log.Println("Got signal:", s)
		server.Close()
	}()
	err = server.ListenAndServe()
	log.Println(err)

}
