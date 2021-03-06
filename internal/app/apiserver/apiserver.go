package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-rest/internal/app/store"
	"io"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Infof("Starting API server on port: %s", s.config.BindAddress[1:])
	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "hello")
	}
}
