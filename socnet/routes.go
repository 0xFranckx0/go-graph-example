package socnet

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

// this function instanciates a new HTTP router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Methods("GET").Path("/distance/{id0:[0-9]+}/{id1:[0-9]+}").Handler(logger(HandlerFunc(distanceHandler)))
	router.Methods("GET").Path("/friends/{id0:[0-9]+}/{id1:[0-9]+}").Handler(logger(HandlerFunc(friendsHandler)))

	return router
}

// distanceHandler returns the distance and the path between 2 points
func distanceHandler(w http.ResponseWriter, r *http.Request, log *logrus.Entry) (int, string, []byte, string) {

	//first we gather the two points
	id0, ok := mux.Vars(r)["id0"]
	if !ok {
		return http.StatusInternalServerError, "", nil, "id0 failure "
	}
	id1, ok := mux.Vars(r)["id1"]
	if !ok {
		return http.StatusInternalServerError, "", nil, "id1 failure "
	}

	// We need to store the path and distance into a structure that will be
	// returned as a JSON
	type distanceJSON struct {
		Distance int   `json:"distance"`
		Path     []int `json:"path"`
	}

	var distance distanceJSON
	distance.Distance, distance.Path = GetDistance(id0, id1)

	if distance.Distance == 0 {
		return http.StatusNotFound, "", nil, "No Path found"
	}

	if distance.Distance == -1 {
		// If distance is negative that means that something went wrong while
		// processing the path.
		return http.StatusInternalServerError, "", nil, "Dijkstra Error while processing the path"
	}

	// Once we filled out the structure we can Marshal it and return
	respByte, err := json.Marshal(distance)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "json.Marshal failed: " + err.Error()
	}

	return http.StatusOK, "application/json", respByte, "SUCCESS"
}

// friendsHandler returns the list of common friends between points
func friendsHandler(w http.ResponseWriter, r *http.Request, log *logrus.Entry) (int, string, []byte, string) {

	//first we gather the two points
	id0, ok := mux.Vars(r)["id0"]
	if !ok {
		return http.StatusInternalServerError, "", nil, "id0 failure "
	}
	id1, ok := mux.Vars(r)["id1"]
	if !ok {
		return http.StatusInternalServerError, "", nil, "id1 failure "
	}

	// We need to store the friend list into a structure that will be
	// returned as a JSON
	type friendsJSON struct {
		Friends []int `json:"friends"`
	}
	var f friendsJSON
	f.Friends = GetCommonFriends(id0, id1)

	// Then we marshal the result and return
	respByte, err := json.Marshal(f)
	if err != nil {
		return http.StatusInternalServerError, "", nil, "json.Marshal failed: " + err.Error()
	}

	return http.StatusOK, "application/json", respByte, "SUCCESS"

}
