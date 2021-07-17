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
	getServers   http.HandlerFunc
	killAll      http.HandlerFunc
	killAtt      http.HandlerFunc
	killDef      http.HandlerFunc
	killSup      http.HandlerFunc
	killAllTribe http.HandlerFunc
	killAttTribe http.HandlerFunc
	killDefTribe http.HandlerFunc
	getPlayers   http.HandlerFunc
	getTribes    http.HandlerFunc
	getVillages  http.HandlerFunc
}

func (h *handlers) init() {
	noop := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	if h.getServers == nil {
		h.getServers = noop
	}
	if h.killAll == nil {
		h.killAll = noop
	}
	if h.killAtt == nil {
		h.killAtt = noop
	}
	if h.killDef == nil {
		h.killDef = noop
	}
	if h.killSup == nil {
		h.killSup = noop
	}
	if h.killAllTribe == nil {
		h.killAllTribe = noop
	}
	if h.killAttTribe == nil {
		h.killAttTribe = noop
	}
	if h.killDefTribe == nil {
		h.killDefTribe = noop
	}
	if h.getPlayers == nil {
		h.getPlayers = noop
	}
	if h.getTribes == nil {
		h.getTribes = noop
	}
	if h.getVillages == nil {
		h.getVillages = noop
	}
}

func prepareTestServer(h *handlers) *httptest.Server {
	if h == nil {
		h = &handlers{}
	}
	h.init()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.URL.Path {
		case EndpointGetServers:
			h.getServers(w, r)
			return
		case EndpointKillAll:
			h.killAll(w, r)
			return
		case EndpointKillAtt:
			h.killAtt(w, r)
			return
		case EndpointKillDef:
			h.killDef(w, r)
			return
		case EndpointKillSup:
			h.killSup(w, r)
			return
		case EndpointKillAllTribe:
			h.killAllTribe(w, r)
			return
		case EndpointKillAttTribe:
			h.killAttTribe(w, r)
			return
		case EndpointKillDefTribe:
			h.killDefTribe(w, r)
			return
		case EndpointPlayer:
			h.getPlayers(w, r)
			return
		case EndpointTribe:
			h.getTribes(w, r)
			return
		case EndpointVillage:
			h.getVillages(w, r)
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

func createWriteCompressedStringHandler(resp string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gzipWriter := gzip.NewWriter(w)
		defer gzipWriter.Close()
		_, err := gzipWriter.Write([]byte(resp))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := gzipWriter.Flush(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
