package probes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	testutils "github.com/RHEcosystemAppEng/DepDraw/src/k8sclient/pkg/test"
	"github.com/gorilla/mux"
)

func Test_heackthCheck(t *testing.T) {
	config := testutils.InitTest()
	health := NewHealthCheckForConfig(config)
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(health.Check)
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
	t.Logf("Received expected HTTP status code %d", http.StatusOK)
}
