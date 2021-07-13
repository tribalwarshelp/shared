package twdataloader

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

func getDefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func uncompressAndReadCsvLines(r io.Reader) ([][]string, error) {
	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		return [][]string{}, err
	}
	defer uncompressedStream.Close()
	return csv.NewReader(uncompressedStream).ReadAll()
}

type handlers struct {
	getServers http.HandlerFunc
}

func (h *handlers) init() {
	noop := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	if h.getServers == nil {
		h.getServers = noop
	}
}

func prepareTestServer(h *handlers) *httptest.Server {
	if h == nil {
		h = &handlers{}
	}
	h.init()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EndpointGetServers:
			h.getServers(w, r)
			return
		case EndpointKillAll:
			return
		case EndpointKillAtt:
			return
		case EndpointKillDef:
			return
		case EndpointKillSup:
			return
		case EndpointKillAllTribe:
			return
		case EndpointKillAttTribe:
			return
		case EndpointKillDefTribe:
			return
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func createWriteStringHandler(resp string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(resp))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
