package apiserver

import (
	"io"
	"net/http"

	"github.com/GuseynovAnar/rest_api.git/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.New(config.Store),
	}
}

// Start server

func (s *APIServer) Start() error {
	if err := s.configureLoger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAdd, s.router)
}

func (s *APIServer) configureLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}

func (s *APIServer) configureStore() error {
	if err := s.store.Open(); err != nil {
		return err
	}

	return nil
}
