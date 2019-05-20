package middlewares_test

import (
	"fmt"
	Api "github.com/aeternas/SwadeshNess/apiClient"
	m "github.com/aeternas/SwadeshNess/clientMiddlewares"
	Conf "github.com/aeternas/SwadeshNess/configuration"
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
	var reader *Conf.Reader = &Conf.Reader{Path: "db/db.json", OsWrapper: wrapper}
	config, _ := reader.ReadConfiguration()
	mdlwr := m.NewCachingDefaultClientMiddleware(&config)
	request := &Api.Request{Data: []byte{}, Cached: false, NetRequest: &http.Request{URL: &url.URL{RawQuery: "translate=translation"}}}
	str := mdlwr.GetKey(request)
	if str != "translate=translation" {
		t.Errorf("Key is not equal to expected: %s", str)
	}
}
