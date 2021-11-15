package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ..

type APIServer struct {
	config * Config
	logger * logrus.Logger
	router * mux.Router
}

// Mew Instance ...

func New(config * Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start Server ...

func (s * APIServer) Start() error {
	if error := s.configureLogger(); error != nil {
		return error
	}

	s.configureRouter()
	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s * APIServer) configureLogger() error {
	level, error := logrus.ParseLevel(s.config.LogLevel)

	if error != nil {
		return  error
	}

	s.logger.SetLevel(level)

	return nil
}

func (s * APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s * APIServer) handleHello() http.HandlerFunc {

	// ... переменные хендлера или типы

	return func (w http.ResponseWriter, r * http.Request) {
		io.WriteString(w, "Hello!")
	}
}