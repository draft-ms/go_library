package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreateAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,string(data))
}

func OrderHandle(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Order:", r.URL.Query().Get("param"))
}

func NewHttpHandler() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello")
	})

	mux.HandleFunc("/order", OrderHandle)

	mux.Handle("/testHandle", &testHandler{})

	return mux
}

func TestWebReqeust(t *testing.T) {
	http.ListenAndServe(":3000",NewHttpHandler())
}

func TestMockHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/order", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	//OrderHandle(res, req)

	assert.Equal(http.StatusOK, res.Code)
}