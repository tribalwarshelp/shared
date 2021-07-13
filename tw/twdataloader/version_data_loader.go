package twdataloader

import (
	phpserialize "github.com/Kichiyaki/go-php-serialize"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Server struct {
	Key string
	URL string
}

type VersionDataLoaderConfig struct {
	Host   string
	Client *http.Client
}

func (cfg *VersionDataLoaderConfig) init() {
	if cfg.Client == nil {
		cfg.Client = getDefaultHTTPClient()
	}
}

type VersionDataLoader struct {
	host   string
	client *http.Client
}

func NewVersionDataLoader(cfg *VersionDataLoaderConfig) *VersionDataLoader {
	if cfg == nil {
		cfg = &VersionDataLoaderConfig{}
	}
	cfg.init()
	return &VersionDataLoader{
		host:   cfg.Host,
		client: cfg.Client,
	}
}

func (dl *VersionDataLoader) LoadServers() ([]*Server, error) {
	resp, err := dl.client.Get(dl.host + EndpointGetServers)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load servers")
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read the response body")
	}
	body, err := phpserialize.Decode(string(bodyBytes))
	if err != nil || body == nil {
		fmtedErr := errors.New("couldn't decode the response body into a go value")
		if err != nil {
			fmtedErr = errors.Wrap(err, fmtedErr.Error())
		}
		return nil, fmtedErr
	}
	bodyMap, ok := body.(map[interface{}]interface{})
	if !ok {
		return nil, errors.Errorf("expected map, got %T", body)
	}

	var servers []*Server
	for serverKey, url := range bodyMap {
		serverKeyStr, ok := serverKey.(string)
		if !ok {
			return nil, errors.Errorf("expected string as the key of the map, got %T", serverKey)
		}
		urlStr, ok := url.(string)
		if !ok {
			return nil, errors.Errorf("expected string as the value of the map, got %T", url)
		}
		if serverKeyStr != "" && urlStr != "" {
			servers = append(servers, &Server{
				Key: serverKeyStr,
				URL: urlStr,
			})
		}
	}
	return servers, nil
}
