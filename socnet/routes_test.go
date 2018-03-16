package socnet

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDistanceRoute(t *testing.T) {
	fetchAPI(t, "GET", "/distance/40/382")
}

func TestDistanceNoRoute(t *testing.T) {
	fetchAPI(t, "GET", "/distance/10/4780")
}

func TestDistanceRouteInvalidId(t *testing.T) {
	fetchAPI(t, "GET", "/distance/34/9999999")
}

func TestFriendsRoute(t *testing.T) {
	fetchAPI(t, "GET", "/friends/4017/3980")
}

func TestNoFriendsRoute(t *testing.T) {
	fetchAPI(t, "GET", "/friends/6764/3980")
}

func TestFriendsRouteInvalidID(t *testing.T) {
	fetchAPI(t, "GET", "/friends/4017/9999999")
}

func fetchAPI(t *testing.T, method string, path string) {
	req, _ := http.NewRequest(method, path, nil)

	resp := httptest.NewRecorder()
	router := NewRouter()
	router.ServeHTTP(resp, req)

	if http.StatusInternalServerError == resp.Code {
		t.Errorf("NOT OK %d\n", resp.Code)
	}
	fmt.Println(resp.Body.String())

}
