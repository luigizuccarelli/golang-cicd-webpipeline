package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/microlib/logger/pkg/multi"
	"lmzsoftware.com/lzuccarelli/golang-cicd-webconsole/pkg/connectors"
)

func TestAllMiddleware(t *testing.T) {

	logger := multi.NewLogger(multi.COLOR, multi.TRACE)
	os.Setenv("TEMPLATE_FILE", "../../html-templates/template.html")
	os.Setenv("CICD_CONSOLE_DIR", "../../tests/console")

	t.Run("IsAlive : should pass", func(t *testing.T) {
		var STATUS int = 200
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v2/sys/info/isalive", nil)
		connectors.NewTestConnectors(logger)
		handler := http.HandlerFunc(IsAlive)
		handler.ServeHTTP(rr, req)

		body, e := ioutil.ReadAll(rr.Body)
		if e != nil {
			t.Fatalf("Should not fail : found error %v", e)
		}
		logger.Trace(fmt.Sprintf("Response %s", string(body)))
		// ignore errors here
		if rr.Code != STATUS {
			t.Errorf(fmt.Sprintf("Handler %s returned with incorrect status code - got (%d) wanted (%d)", "IsAlive", rr.Code, STATUS))
		}
	})

	t.Run("ResultsHandler : should pass", func(t *testing.T) {
		var STATUS int = 200
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/cicd/golang-test", nil)
		con := connectors.NewTestConnectors(logger)
		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ResultsHandler(w, req, con)
		})
		//Hack to try to fake gorilla/mux vars
		vars := map[string]string{
			"project": "golang-test",
		}
		req = mux.SetURLVars(req, vars)
		handler.ServeHTTP(rr, req)
		body, e := ioutil.ReadAll(rr.Body)
		if e != nil {
			t.Fatalf("Should not fail : found error %v", e)
		}
		logger.Trace(fmt.Sprintf("Response %s", string(body)))
		// ignore errors here
		if rr.Code != STATUS {
			t.Errorf(fmt.Sprintf("Handler %s returned with incorrect status code - got (%d) wanted (%d)", "ResultsHandler", rr.Code, STATUS))
		}
	})

	t.Run("ResultsHandler : should fail (empty param)", func(t *testing.T) {
		var STATUS int = 500
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/cicd/golang-test", nil)
		con := connectors.NewTestConnectors(logger)
		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ResultsHandler(w, req, con)
		})
		handler.ServeHTTP(rr, req)
		body, e := ioutil.ReadAll(rr.Body)
		if e != nil {
			t.Fatalf("Should not fail : found error %v", e)
		}
		logger.Trace(fmt.Sprintf("Response %s", string(body)))
		// ignore errors here
		if rr.Code != STATUS {
			t.Errorf(fmt.Sprintf("Handler %s returned with incorrect status code - got (%d) wanted (%d)", "ResultsHandler", rr.Code, STATUS))
		}
	})

	t.Run("ResultsHandler : should fail (incorrect directory)", func(t *testing.T) {
		var STATUS int = 500
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/cicd/golang-fake", nil)
		con := connectors.NewTestConnectors(logger)
		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ResultsHandler(w, req, con)
		})
		//Hack to try to fake gorilla/mux vars
		vars := map[string]string{
			"project": "golang-fake",
		}
		req = mux.SetURLVars(req, vars)

		handler.ServeHTTP(rr, req)
		body, e := ioutil.ReadAll(rr.Body)
		if e != nil {
			t.Fatalf("Should not fail : found error %v", e)
		}
		logger.Trace(fmt.Sprintf("Response %s", string(body)))
		// ignore errors here
		if rr.Code != STATUS {
			t.Errorf(fmt.Sprintf("Handler %s returned with incorrect status code - got (%d) wanted (%d)", "ResultsHandler", rr.Code, STATUS))
		}
	})

	t.Run("ResultsHandler : should fail (incorrect template directory)", func(t *testing.T) {
		var STATUS int = 500
		os.Setenv("TEMPLATE_FILE", "html-templates/template.html")
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/cicd/golang-test", nil)
		con := connectors.NewTestConnectors(logger)
		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ResultsHandler(w, req, con)
		})
		//Hack to try to fake gorilla/mux vars
		vars := map[string]string{
			"project": "golang-test",
		}
		req = mux.SetURLVars(req, vars)

		handler.ServeHTTP(rr, req)
		body, e := ioutil.ReadAll(rr.Body)
		if e != nil {
			t.Fatalf("Should not fail : found error %v", e)
		}
		logger.Trace(fmt.Sprintf("Response %s", string(body)))
		// ignore errors here
		if rr.Code != STATUS {
			t.Errorf(fmt.Sprintf("Handler %s returned with incorrect status code - got (%d) wanted (%d)", "ResultsHandler", rr.Code, STATUS))
		}
	})
}
