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

func (cfg *VersionDataLoaderConfig) Init() {
	if cfg.Client == nil {
		cfg.Client = getDefaultHTTPClient()
	}
}

type VersionDataLoader interface {
	LoadServers() ([]*Server, error)
}

type versionDataLoader struct {
	host   string
	client *http.Client
}

func NewVersionDataLoader(cfg *VersionDataLoaderConfig) VersionDataLoader {
	if cfg == nil {
		cfg = &VersionDataLoaderConfig{}
	}
	cfg.Init()
	return &versionDataLoader{
		host:   cfg.Host,
		client: cfg.Client,
	}
}

func (d *versionDataLoader) LoadServers() ([]*Server, error) {
	resp, err := d.client.Get(fmt.Sprintf("https://%s%s", d.host, EndpointGetServers))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load servers")
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read the response body")
	}
	body, err := phpserialize.Decode(string(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't decode the response body into the go value")
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
