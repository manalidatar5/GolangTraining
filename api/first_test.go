package api_test

import (
	"bytes"
	"github.com/you/hello/api"
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestGetUser(t *testing.T){
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Getuser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestLogin(t *testing.T){
	var jsonStr = []byte(`{"username":"Admin","password":"Admin123"}`)
	req, err := http.NewRequest("GET", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateTokenEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateEntry(t *testing.T) {
	var jsonStr = []byte(`{"id":993,"name":"xyz","age":22,"department":"xyz@pqr"}`)

	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.UserCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Status":201,"message":"Created"}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}

}

func TestEditEntry(t *testing.T) {

	var jsonStr = []byte(`{"id":59,"name":"xyzzz","age":224,"department":"xyzpqr"}`)
	req, err := http.NewRequest("PUT", "/update", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Userupdate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Status":200,"message":"Updated"}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}
func TestDeleteEntry(t *testing.T) {
	var jsonStr = []byte(`{"id":60}`)
	req, err := http.NewRequest("DELETE", "/delete", bytes.NewBuffer(jsonStr) )
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "60")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Userdelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Status":200,"message":"Deleted"}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}