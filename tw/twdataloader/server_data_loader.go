package twdataloader

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/tribalwarshelp/shared/tw/twmodel"
)

type ServerDataLoaderConfig struct {
	BaseURL string
	Client  *http.Client
}

func (cfg *ServerDataLoaderConfig) init() {
	if cfg.Client == nil {
		cfg.Client = getDefaultHTTPClient()
	}
}

type ServerDataLoader struct {
	baseURL string
	client  *http.Client
}

func NewServerDataLoader(cfg *ServerDataLoaderConfig) *ServerDataLoader {
	if cfg == nil {
		cfg = &ServerDataLoaderConfig{}
	}
	cfg.init()
	return &ServerDataLoader{
		cfg.BaseURL,
		cfg.Client,
	}
}

type parsedODLine struct {
	ID    int
	Rank  int
	Score int
}

func (dl *ServerDataLoader) parseODLine(line []string) (*parsedODLine, error) {
	if len(line) != 3 {
		return nil, errors.New("invalid line format (should be rank,id,score)")
	}
	p := &parsedODLine{}
	var err error
	p.Rank, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, errors.Wrap(err, "parsedODLine.Rank")
	}
	p.ID, err = strconv.Atoi(line[1])
	if err != nil {
		return nil, errors.Wrap(err, "parsedODLine.ID")
	}
	p.Score, err = strconv.Atoi(line[2])
	if err != nil {
		return nil, errors.Wrap(err, "parsedODLine.Score")
	}
	return p, nil
}

func (dl *ServerDataLoader) LoadOD(tribe bool) (map[int]*twmodel.OpponentsDefeated, error) {
	m := make(map[int]*twmodel.OpponentsDefeated)
	formattedURLs := []string{
		buildURL(dl.baseURL, EndpointKillAll),
		buildURL(dl.baseURL, EndpointKillAtt),
		buildURL(dl.baseURL, EndpointKillDef),
		buildURL(dl.baseURL, EndpointKillSup),
	}
	if tribe {
		formattedURLs = []string{
			buildURL(dl.baseURL, EndpointKillAllTribe),
			buildURL(dl.baseURL, EndpointKillAttTribe),
			buildURL(dl.baseURL, EndpointKillDefTribe),
			"",
		}
	}
	for _, formattedURL := range formattedURLs {
		if formattedURL == "" {
			continue
		}
		lines, err := dl.getCSVData(formattedURL, true)
		if err != nil {
			lines, err = dl.getCSVData(strings.ReplaceAll(formattedURL, ".gz", ""), false)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't load data, formattedURL %s", formattedURL)
			}
		}
		for _, line := range lines {
			parsed, err := dl.parseODLine(line)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't parse the line, url %s, line %s", formattedURL, strings.Join(line, ","))
			}
			if _, ok := m[parsed.ID]; !ok {
				m[parsed.ID] = &twmodel.OpponentsDefeated{}
			}
			switch formattedURL {
			case formattedURLs[0]:
				m[parsed.ID].RankTotal = parsed.Rank
				m[parsed.ID].ScoreTotal = parsed.Score
			case formattedURLs[1]:
				m[parsed.ID].RankAtt = parsed.Rank
				m[parsed.ID].ScoreAtt = parsed.Score
			case formattedURLs[2]:
				m[parsed.ID].RankDef = parsed.Rank
				m[parsed.ID].ScoreDef = parsed.Score
			case formattedURLs[3]:
				m[parsed.ID].RankSup = parsed.Rank
				m[parsed.ID].ScoreSup = parsed.Score
			}
		}
	}
	return m, nil
}

func (dl *ServerDataLoader) parsePlayerLine(line []string) (*twmodel.Player, error) {
	if len(line) != 6 {
		return nil, errors.New("invalid line format (should be id,name,tribeid,villages,points,rank)")
	}

	var err error
	ex := true
	player := &twmodel.Player{
		Exists: &ex,
	}
	player.ID, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, errors.Wrap(err, "player.ID")
	}
	player.Name, err = url.QueryUnescape(line[1])
	if err != nil {
		return nil, errors.Wrap(err, "player.Name")
	}
	player.TribeID, err = strconv.Atoi(line[2])
	if err != nil {
		return nil, errors.Wrap(err, "player.TribeID")
	}
	player.TotalVillages, err = strconv.Atoi(line[3])
	if err != nil {
		return nil, errors.Wrap(err, "player.TotalVillages")
	}
	player.Points, err = strconv.Atoi(line[4])
	if err != nil {
		return nil, errors.Wrap(err, "player.Points")
	}
	player.Rank, err = strconv.Atoi(line[5])
	if err != nil {
		return nil, errors.Wrap(err, "player.Rank")
	}

	return player, nil
}

func (dl *ServerDataLoader) LoadPlayers() ([]*twmodel.Player, error) {
	formattedURL := buildURL(dl.baseURL, EndpointPlayer)
	lines, err := dl.getCSVData(formattedURL, true)
	if err != nil {
		lines, err = dl.getCSVData(buildURL(dl.baseURL, EndpointPlayerNotGzipped), false)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load data, formattedURL %s", formattedURL)
		}
	}

	var players []*twmodel.Player
	for _, line := range lines {
		player, err := dl.parsePlayerLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't parse the line, url %s, line %s", formattedURL, strings.Join(line, ","))
		}
		players = append(players, player)
	}

	return players, nil
}

func (dl *ServerDataLoader) parseTribeLine(line []string) (*twmodel.Tribe, error) {
	if len(line) != 8 {
		return nil, errors.New("invalid line format (should be id,name,tag,members,villages,points,allpoints,rank)")
	}

	var err error
	ex := true
	tribe := &twmodel.Tribe{
		Exists: &ex,
	}
	tribe.ID, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.ID")
	}
	tribe.Name, err = url.QueryUnescape(line[1])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.Name")
	}
	tribe.Tag, err = url.QueryUnescape(line[2])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.Tag")
	}
	tribe.TotalMembers, err = strconv.Atoi(line[3])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.TotalMembers")
	}
	tribe.TotalVillages, err = strconv.Atoi(line[4])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.TotalVillages")
	}
	tribe.Points, err = strconv.Atoi(line[5])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.Points")
	}
	tribe.AllPoints, err = strconv.Atoi(line[6])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.AllPoints")
	}
	tribe.Rank, err = strconv.Atoi(line[7])
	if err != nil {
		return nil, errors.Wrap(err, "tribe.Rank")
	}

	return tribe, nil
}

func (dl *ServerDataLoader) LoadTribes() ([]*twmodel.Tribe, error) {
	formattedURL := buildURL(dl.baseURL, EndpointTribe)
	lines, err := dl.getCSVData(formattedURL, true)
	if err != nil {
		lines, err = dl.getCSVData(buildURL(dl.baseURL, EndpointTribeNotGzipped), false)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load data, url %s", formattedURL)
		}
	}
	var tribes []*twmodel.Tribe
	for _, line := range lines {
		tribe, err := dl.parseTribeLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't parse the line, url %s, line %s", formattedURL, strings.Join(line, ","))
		}
		tribes = append(tribes, tribe)
	}
	return tribes, nil
}

func (dl *ServerDataLoader) parseVillageLine(line []string) (*twmodel.Village, error) {
	if len(line) != 7 {
		return nil, errors.New("invalid line format (should be id,name,x,y,playerID,points,bonus)")
	}
	var err error
	village := &twmodel.Village{}
	village.ID, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, errors.Wrap(err, "village.ID")
	}
	village.Name, err = url.QueryUnescape(line[1])
	if err != nil {
		return nil, errors.Wrap(err, "village.Name")
	}
	village.X, err = strconv.Atoi(line[2])
	if err != nil {
		return nil, errors.Wrap(err, "village.X")
	}
	village.Y, err = strconv.Atoi(line[3])
	if err != nil {
		return nil, errors.Wrap(err, "village.Y")
	}
	village.PlayerID, err = strconv.Atoi(line[4])
	if err != nil {
		return nil, errors.Wrap(err, "village.PlayerID")
	}
	village.Points, err = strconv.Atoi(line[5])
	if err != nil {
		return nil, errors.Wrap(err, "village.Points")
	}
	village.Bonus, err = strconv.Atoi(line[6])
	if err != nil {
		return nil, errors.Wrap(err, "village.Bonus")
	}
	return village, nil
}

func (dl *ServerDataLoader) LoadVillages() ([]*twmodel.Village, error) {
	formattedURL := buildURL(dl.baseURL, EndpointVillage)
	lines, err := dl.getCSVData(formattedURL, true)
	if err != nil {
		lines, err = dl.getCSVData(buildURL(dl.baseURL, EndpointVillageNotGzipped), false)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load data, formattedURL %s", formattedURL)
		}
	}
	var villages []*twmodel.Village
	for _, line := range lines {
		village, err := dl.parseVillageLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't parse the line, formattedURL %s, line %s", formattedURL, strings.Join(line, ","))
		}
		villages = append(villages, village)
	}
	return villages, nil
}

func (dl *ServerDataLoader) parseEnnoblementLine(line []string) (*twmodel.Ennoblement, error) {
	if len(line) != 4 {
		return nil, errors.New("invalid line format (should be village_id,timestamp,new_owner_id,old_owner_id)")
	}
	var err error
	ennoblement := &twmodel.Ennoblement{}
	ennoblement.VillageID, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, errors.Wrap(err, "ennoblement.VillageID")
	}
	timestamp, err := strconv.Atoi(line[1])
	if err != nil {
		return nil, errors.Wrap(err, "timestamp")
	}
	ennoblement.EnnobledAt = time.Unix(int64(timestamp), 0)
	ennoblement.NewOwnerID, err = strconv.Atoi(line[2])
	if err != nil {
		return nil, errors.Wrap(err, "ennoblement.NewOwnerID")
	}
	ennoblement.OldOwnerID, err = strconv.Atoi(line[3])
	if err != nil {
		return nil, errors.Wrap(err, "ennoblement.OldOwnerID")
	}

	return ennoblement, nil
}

type LoadEnnoblementsConfig struct {
	EnnobledAtGT time.Time
}

func (dl *ServerDataLoader) LoadEnnoblements(cfg *LoadEnnoblementsConfig) ([]*twmodel.Ennoblement, error) {
	if cfg == nil {
		cfg = &LoadEnnoblementsConfig{}
	}
	yesterdaysDate := time.Now().Add(-23 * time.Hour)
	formattedURL := buildURL(dl.baseURL, EndpointConquer)
	compressed := true
	if cfg.EnnobledAtGT.After(yesterdaysDate) || cfg.EnnobledAtGT.Equal(yesterdaysDate) {
		formattedURL = buildURL(dl.baseURL, fmt.Sprintf(EndpointGetConquer, cfg.EnnobledAtGT.Unix()))
		compressed = false
	}
	lines, err := dl.getCSVData(formattedURL, compressed)
	if err != nil && compressed {
		lines, err = dl.getCSVData(buildURL(dl.baseURL, EndpointConquerNotGzipped), false)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't load data, formattedURL %s", formattedURL)
	}

	var ennoblements []*twmodel.Ennoblement
	for _, line := range lines {
		ennoblement, err := dl.parseEnnoblementLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't parse the line, formattedURL %s, line %s", formattedURL, strings.Join(line, ","))
		}
		if ennoblement.EnnobledAt.After(cfg.EnnobledAtGT) {
			ennoblements = append(ennoblements, ennoblement)
		}
	}
	return ennoblements, nil
}

func (dl *ServerDataLoader) GetConfig() (*twmodel.ServerConfig, error) {
	formattedURL := buildURL(dl.baseURL, EndpointConfig)
	cfg := &twmodel.ServerConfig{}
	err := dl.getXML(formattedURL, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getConfig")
	}
	return cfg, nil
}

func (dl *ServerDataLoader) GetBuildingConfig() (*twmodel.BuildingConfig, error) {
	formattedURL := buildURL(dl.baseURL, EndpointBuildingConfig)
	cfg := &twmodel.BuildingConfig{}
	err := dl.getXML(formattedURL, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getBuildingConfig")
	}
	return cfg, nil
}

func (dl *ServerDataLoader) GetUnitConfig() (*twmodel.UnitConfig, error) {
	formattedURL := buildURL(dl.baseURL, EndpointUnitConfig)
	cfg := &twmodel.UnitConfig{}
	err := dl.getXML(formattedURL, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getUnitConfig")
	}
	return cfg, nil
}

func (dl *ServerDataLoader) getCSVData(url string, compressed bool) ([][]string, error) {
	resp, err := dl.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !compressed {
		return csv.NewReader(resp.Body).ReadAll()
	}
	return uncompressAndReadCsvLines(resp.Body)
}

func (dl *ServerDataLoader) getXML(url string, decode interface{}) error {
	resp, err := dl.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, decode)
}
