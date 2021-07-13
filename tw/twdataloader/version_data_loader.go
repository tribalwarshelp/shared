package twdataloader

import (
	"fmt"
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
	resp, err := dl.client.Get(fmt.Sprintf("https://%s%s", dl.host, EndpointGetServers))
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

	var servers []*Server
	for serverKey, url := range body.(map[interface{}]interface{}) {
		serverKeyStr := serverKey.(string)
		urlStr := url.(string)
		if serverKeyStr != "" && urlStr != "" {
			servers = append(servers, &Server{
				Key: serverKeyStr,
				URL: urlStr,
			})
		}
	}
	return servers, nil
}
