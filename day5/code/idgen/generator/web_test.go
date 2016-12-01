package generator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func testIdGenHandlerFunc(idgen idGenerator, t *testing.T) {
	server := httptest.NewServer(NewIdGenHandler(idgen))

	defer server.Close()

	for i := 1; i <= 10; i++ {
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		expected := idGenResponseGet{ID: int32(i)}
		var actual idGenResponseGet
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&actual)
		resp.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		if expected != actual {
			t.Errorf("Expected the message %q, recieved %q\n", expected, actual)
		}
	}
}

func testIdGenHandlerOtherWayFunc(idgen idGenerator, t *testing.T) {
	handler := NewIdGenHandler(idgen)

	for i := 1; i <= 10; i++ {
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)
		if resp.Code != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.Code)
		}
		expected := "{\"id\":" + strconv.Itoa(i) + "}\n"
		actual, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(actual) != expected {
			t.Errorf("Expected the message %q, recieved %q\n", expected, string(actual))
		}
	}
}

func TestIdGenHandlerAtomicOtherWay(t *testing.T) {
	testIdGenHandlerOtherWayFunc(NewIdGeneratorAtomic(), t)
}

func TestIdGenHandlerAtomic(t *testing.T) {
	testIdGenHandlerFunc(NewIdGeneratorAtomic(), t)
}

func TestIdGenHandlerMutex(t *testing.T) {
	testIdGenHandlerFunc(NewIdGeneratorMutex(), t)
}

func TestIdGenHandlerChan(t *testing.T) {
	testIdGenHandlerFunc(NewIdGeneratorChan(), t)
}
