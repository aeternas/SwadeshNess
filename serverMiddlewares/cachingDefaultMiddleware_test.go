package middlewares_test

import (
	"fmt"
	Api "github.com/aeternas/SwadeshNess/apiClient"
	Conf "github.com/aeternas/SwadeshNess/configuration"
	m "github.com/aeternas/SwadeshNess/serverMiddlewares"
	Wrappers "github.com/aeternas/SwadeshNess/wrappers"
	"net/http"
	"net/url"
	"testing"
)

func TestNoop(t *testing.T) {
	fmt.Println("")
}

func TestGetKey(t *testing.T) {
	var wrapper = Wrappers.New(new(Wrappers.OsWrapper))
	var reader *Conf.Reader = &Conf.Reader{Path: "../configuration/db.json", OsWrapper: wrapper}
	config, _ := reader.ReadConfiguration()
	mdlwr := m.NewCachingDefaultServerMiddleware(&config)
	request := &Api.Request{Data: []byte{}, Cached: false, NetRequest: &http.Request{URL: &url.URL{RawQuery: "translate=translation"}}}
	str := mdlwr.GetKey(request)
	if str != "translate=translation&v=4" {
		t.Errorf("Key is not equal to expected: %s", str)
	}
}
