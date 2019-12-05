///////////////////////////////////////////////////////////////////////////
// Copyright 2019 Roku, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//////////////////////////////////////////////////////////////////////////

package httpServer

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	ecp "ecpClient"
	"bytes"
)

var sessionId = "test"

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}


func checkSuccessStatus(code int, t *testing.T) {
	if  code != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, code)
	}
}

func GetMockedServerWithSession(res *string) *Server {
	sessions := make(map[string]*SessionInfo)
	sessions[sessionId] = &SessionInfo{
		capability: &Capability{},
		client: ecp.GetMockedClient(res),
	}
	server := &Server{
		router: mux.NewRouter(),
		sessions: sessions,
	}	
	return server
}

func TestAppsHandlerWithValidGet(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/session/test/apps", nil)
    server := GetMockedServerWithSession(&ecp.AppsResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/apps").Handler(Middleware(server.GetAppsHandler())).Methods("GET")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validAppsResponse, rr.Body.String(), "Response body differs")
}

func TestCurrentAppHandlerWithValidGet(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/session/test/apps", nil)
    server := GetMockedServerWithSession(&ecp.AppResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/apps").Handler(Middleware(server.GetCurrentAppHandler())).Methods("GET")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validAppResponse, rr.Body.String(), "Response body differs")
}

func TestPressButtonHandlerWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/press",  bytes.NewBuffer([]byte(`{"button" : "up"}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/press").Handler(Middleware(server.GetPressButtonHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestPressButtonSequenceHandlerWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/press",  bytes.NewBuffer([]byte(`{"button_sequence" : ["up", "right"]}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/press").Handler(Middleware(server.GetPressButtonHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestInstallHandlerWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/install",  bytes.NewBuffer([]byte(`{"channelId" : "dev"}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/install").Handler(Middleware(server.GetInstallHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestLaunchHandlerWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/launch",  bytes.NewBuffer([]byte(`{"channelId" : "dev"}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/launch").Handler(Middleware(server.GetLaunchHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestDeleteSessionWithValidDelete(t *testing.T) {
    req, err := http.NewRequest(http.MethodDelete, "/session/test",  nil)
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}").Handler(Middleware(server.GetSessionDeleteHandler())).Methods("DELETE")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestSessionWithValidGet(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/session/test",  nil)
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}").Handler(Middleware(server.GetSessionHandler())).Methods("GET")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validSessionResponse, rr.Body.String(), "Response body differs")
}

func TestSessionsInfoWithValidGet(t *testing.T) {
    req, err := http.NewRequest(http.MethodGet, "/sessions",  nil)
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/sessions").Handler(Middleware(server.GetSessionsInfoHandler())).Methods("GET")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validSessionsResponse, rr.Body.String(), "Response body differs")
}

func TestTimeoutsWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/sessions/test/timeouts",  bytes.NewBuffer([]byte(`{"type" : "implicit", "ms": 10}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/sessions/{sessionId}/timeouts").Handler(Middleware(server.GetTimeoutsHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestTimeoutImplicitWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/sessions/test/timeouts/implicit_wait",  bytes.NewBuffer([]byte(`{"ms": 10}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/sessions/{sessionId}/timeouts/implicit_wait").Handler(Middleware(server.GetImplicitTimeoutHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}


func TestTimeoutPressWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/sessions/test/timeouts/press_wait",  bytes.NewBuffer([]byte(`{"ms": 10}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/sessions/{sessionId}/timeouts/press_wait").Handler(Middleware(server.GetImplicitTimeoutHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validResponseWithNullValue, rr.Body.String(), "Response body differs")
}

func TestElementWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/element",  bytes.NewBuffer([]byte(`{"elementData": [{
		"using": "tag",
		"value": "Poster"
	}]}`)))
    server := GetMockedServerWithSession(&ecp.UiResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/element").Handler(Middleware(server.GetElementHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validElementResponse, rr.Body.String(), "Response body differs")
}

func TestElementsWithValidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/elements",  bytes.NewBuffer([]byte(`{ "elementData": [{
		"using": "tag",
		"value": "Poster"
	}]}`)))
    server := GetMockedServerWithSession(&ecp.UiResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/elements").Handler(Middleware(server.GetElementsHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	checkSuccessStatus(rr.Code, t)
	assert.JSONEq(t, validElementsResponse, rr.Body.String(), "Response body differs")
}


func TestElementWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/element",  bytes.NewBuffer([]byte(`{}`)))
    server := GetMockedServerWithSession(&ecp.UiResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/element").Handler(Middleware(server.GetElementHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestElementsWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/elements",  bytes.NewBuffer([]byte(`{}`)))
    server := GetMockedServerWithSession(&ecp.UiResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/elements").Handler(Middleware(server.GetElementsHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestTimeoutImplicitWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/sessions/test/timeouts/implicit_wait",  bytes.NewBuffer([]byte(`{"ms": -5}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/sessions/{sessionId}/timeouts/implicit_wait").Handler(Middleware(server.GetImplicitTimeoutHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestInstallHandlerWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/install",  bytes.NewBuffer([]byte(`{}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/install").Handler(Middleware(server.GetInstallHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestLaunchHandlerWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/launch",  bytes.NewBuffer([]byte(`{}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/launch").Handler(Middleware(server.GetLaunchHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPressButtonHandlerWithInvalidPost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPost, "/session/test/press",  bytes.NewBuffer([]byte(`{}`)))
    server := GetMockedServerWithSession(&ecp.SuccessResponseMock)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.Path("/session/{sessionId}/press").Handler(Middleware(server.GetPressButtonHandler())).Methods("POST")
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}