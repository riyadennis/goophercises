package signal

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	arg := make([]int, 10)
	for i := 0; i < 10; i++ {
		arg[i] = r.Int()
	}
	got := Pick(arg)
	for _, v := range arg {
		if got == v {
			return
		}
	}
	t.Errorf("Pick(%d) = %d", seed, got)
}

func TestHandler(t *testing.T) {
	reader := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err %v", err)
	}
	Handler(reader, request)
	resp := reader.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Handler() response %v", resp.StatusCode)
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %v", contentType)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll err %v", err)
	}
	p := &Person{}
	err = json.Unmarshal(data, p)
	if err != nil {
		t.Errorf("json.Unmarshal err %v", err)
	}
	if p.Name != "Sam" {
		t.Errorf("name got %v, want %v", "Sam", p.Name)
	}
	if p.Address != "128-138 High Road" {
		t.Errorf("address got %v, want %v", "128-138 High Road", p.Address)
	}
}
