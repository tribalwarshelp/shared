package twdataloader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func prepareTestServer(resp string) *httptest.Server {
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EndpointGetServers:
			_, err := fmt.Fprintln(w, resp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func TestLoadServers(t *testing.T) {
	t.Run("invalid response", func(t *testing.T) {
		type scenario struct {
			resp           string
			expectedErrMsg string
		}

		scenarios := []scenario{
			{
				resp:           `:"https://pl165.plemiona.pl";s:5:"pl166";s:25:"https://pl166.plemiona.pl";s:5:"pl167";s:25:"https://pl167.plemiona.pl";}`,
				expectedErrMsg: "couldn't decode the response body into a go value",
			},
			{
				resp:           `a:19:{s:5:"pl150"s:25"https://pl150.plemiona.pl";}`,
				expectedErrMsg: "expected string as the value of the map, got <nil>",
			},
			{
				resp:           "a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}",
				expectedErrMsg: "expected string as the key of the map, got int64",
			},
			{
				resp:           `O:8:"stdClass":0:{}`,
				expectedErrMsg: "expected map, got *phpserialize.PhpObject",
			},
			{
				resp:           `a:2:{s:3:"asd";i:123;s:4:"asd2";i:123;}`,
				expectedErrMsg: "expected string as the value of the map, got int64",
			},
		}

		for _, scenario := range scenarios {
			ts := prepareTestServer(scenario.resp)

			dl := NewVersionDataLoader(&VersionDataLoaderConfig{
				Host:   strings.ReplaceAll(ts.URL, "https://", ""),
				Client: ts.Client(),
			})

			_, err := dl.LoadServers()
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)

			ts.Close()
		}
	})

	t.Run("success", func(t *testing.T) {
		resp := `a:19:{s:5:"pl150";s:25:"https://pl150.plemiona.pl";s:5:"pl151";s:25:"https://pl151.plemiona.pl";s:5:"pl152";s:25:"https://pl152.plemiona.pl";s:5:"pl153";s:25:"https://pl153.plemiona.pl";s:5:"pl154";s:25:"https://pl154.plemiona.pl";s:5:"pl155";s:25:"https://pl155.plemiona.pl";s:5:"pl156";s:25:"https://pl156.plemiona.pl";s:5:"pl157";s:25:"https://pl157.plemiona.pl";s:5:"pl158";s:25:"https://pl158.plemiona.pl";s:5:"pl159";s:25:"https://pl159.plemiona.pl";s:5:"pl160";s:25:"https://pl160.plemiona.pl";s:5:"pl161";s:25:"https://pl161.plemiona.pl";s:5:"pl162";s:25:"https://pl162.plemiona.pl";s:5:"pl163";s:25:"https://pl163.plemiona.pl";s:5:"pl164";s:25:"https://pl164.plemiona.pl";s:4:"plp7";s:24:"https://plp7.plemiona.pl";s:5:"pl165";s:25:"https://pl165.plemiona.pl";s:5:"pl166";s:25:"https://pl166.plemiona.pl";s:5:"pl167";s:25:"https://pl167.plemiona.pl";}`
		ts := prepareTestServer(resp)
		defer ts.Close()

		expectedLength, err := strconv.Atoi(strings.Split(resp, ":")[1])
		assert.Nil(t, err)

		dl := NewVersionDataLoader(&VersionDataLoaderConfig{
			Host:   strings.ReplaceAll(ts.URL, "https://", ""),
			Client: ts.Client(),
		})

		servers, err := dl.LoadServers()
		assert.Nil(t, err)
		cnt := 0
		for _, server := range servers {
			if strings.Contains(resp, server.URL) && strings.Contains(resp, server.Key) {
				cnt++
			}
		}
		assert.Equal(t, expectedLength, cnt)
	})
}
