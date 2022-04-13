package api

import (
	"encoding/json"
	"net/http"

	"github.com/chorro/openesp/service"
	"github.com/gorilla/mux"
)

type RestAPI interface {
	CreateDevice(w http.ResponseWriter, r *http.Request)
	GetDevices(w http.ResponseWriter, r *http.Request)
	Router() http.Handler
}

type restApiImpl struct {
	router  *mux.Router
	service service.DeviceRule
}

func NewRestAPI(service service.DeviceRule) RestAPI {
	apiHandler := &restApiImpl{
		router:  mux.NewRouter(),
		service: service,
	}

	// http://<ip>:<port>/service/v0
	subRoute := apiHandler.router.PathPrefix("/service/v0").Subrouter()
	// POST http://<ip>:<port>/service/v0/devices
	subRoute.HandleFunc("/devices", apiHandler.CreateDevice).Methods(http.MethodPost)
	// GET http://<ip>:<port>/service/v0/devices
	subRoute.HandleFunc("/devices", apiHandler.GetDevices).Methods(http.MethodGet)
	return apiHandler
}

func (s *restApiImpl) GetDevices(w http.ResponseWriter, r *http.Request) {

	deviceList, _ := s.service.GetDevices()

	json.NewEncoder(w).Encode(deviceList)

}

func (s *restApiImpl) CreateDevice(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Name string `json:"name"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.service.CreateDevice(payload.Name)

	// devicesList, _ := s.service.GetDevices()
	// fmt.Println(devicesList)
}

func (s *restApiImpl) Router() http.Handler {
	return s.router
}
