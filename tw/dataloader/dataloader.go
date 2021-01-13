package dataloader

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tribalwarshelp/shared/models"
)

const (
	EndpointConfig                 = "/interface.php?func=get_config"
	EndpointUnitConfig             = "/interface.php?func=get_unit_info"
	EndpointBuildingConfig         = "/interface.php?func=get_building_info"
	EndpointPlayer                 = "/map/player.txt.gz"
	EndpointPlayerNotGzipped       = "/map/player.txt"
	EndpointTribe                  = "/map/ally.txt.gz"
	EndpointTribeNotGzipped        = "/map/ally.txt"
	EndpointVillage                = "/map/village.txt.gz"
	EndpointVillageNotGzipped      = "/map/village.txt"
	EndpointKillAtt                = "/map/kill_att.txt.gz"
	EndpointKillAttNotGzipped      = "/map/kill_att.txt"
	EndpointKillDef                = "/map/kill_def.txt.gz"
	EndpointKillDefNotGzipped      = "/map/kill_def.txt"
	EndpointKillSup                = "/map/kill_sup.txt.gz"
	EndpointKillSupNotGzipped      = "/map/kill_sup.txt"
	EndpointKillAll                = "/map/kill_all.txt.gz"
	EndpointKillAllNotGzipped      = "/map/kill_all.txt"
	EndpointKillAttTribe           = "/map/kill_att_tribe.txt.gz"
	EndpointKillAttTribeNotGzipped = "/map/kill_att_tribe.txt"
	EndpointKillDefTribe           = "/map/kill_def_tribe.txt.gz"
	EndpointKillDefTribeNotGzipped = "/map/kill_def_tribe.txt"
	EndpointKillAllTribe           = "/map/kill_all_tribe.txt.gz"
	EndpointKillAllTribeNotGzipped = "/map/kill_all_tribe.txt"
	EndpointConquer                = "/map/conquer.txt.gz"
	EndpointConquerNotGzipped      = "/map/conquer.txt"
	EndpointGetConquer             = "/interface.php?func=get_conquer&since=%d"
)

type DataLoader interface {
	LoadOD(tribe bool) (map[int]*models.OpponentsDefeated, error)
	LoadPlayers() ([]*models.Player, error)
	LoadTribes() ([]*models.Tribe, error)
	LoadVillages() ([]*models.Village, error)
	LoadEnnoblements(cfg *LoadEnnoblementsConfig) ([]*models.Ennoblement, error)
	GetConfig() (*models.ServerConfig, error)
	GetBuildingConfig() (*models.BuildingConfig, error)
	GetUnitConfig() (*models.UnitConfig, error)
}

type Config struct {
	BaseURL string
	Client  *http.Client
}

type dataLoader struct {
	baseURL string
	client  *http.Client
}

func New(cfg *Config) DataLoader {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.Client == nil {
		cfg.Client = &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	return &dataLoader{
		cfg.BaseURL,
		cfg.Client,
	}
}

type parsedODLine struct {
	ID    int
	Rank  int
	Score int
}

func (d *dataLoader) parseODLine(line []string) (*parsedODLine, error) {
	if len(line) != 3 {
		return nil, fmt.Errorf("Invalid line format (should be rank,id,score)")
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

func (d *dataLoader) LoadOD(tribe bool) (map[int]*models.OpponentsDefeated, error) {
	m := make(map[int]*models.OpponentsDefeated)
	urls := []string{
		fmt.Sprintf("%s%s", d.baseURL, EndpointKillAll),
		fmt.Sprintf("%s%s", d.baseURL, EndpointKillAtt),
		fmt.Sprintf("%s%s", d.baseURL, EndpointKillDef),
		fmt.Sprintf("%s%s", d.baseURL, EndpointKillSup),
	}
	if tribe {
		urls = []string{
			fmt.Sprintf("%s%s", d.baseURL, EndpointKillAllTribe),
			fmt.Sprintf("%s%s", d.baseURL, EndpointKillAttTribe),
			fmt.Sprintf("%s%s", d.baseURL, EndpointKillDefTribe),
			"",
		}
	}
	for _, url := range urls {
		if url == "" {
			continue
		}
		lines, err := d.getCSVData(url, true)
		if err != nil {
			//fallback to not gzipped file
			lines, err = d.getCSVData(strings.ReplaceAll(url, ".gz", ""), false)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot get data, url %s", url)
			}
		}
		for _, line := range lines {
			parsed, err := d.parseODLine(line)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot parse line, url %s, line %s", url, strings.Join(line, ","))
			}
			if _, ok := m[parsed.ID]; !ok {
				m[parsed.ID] = &models.OpponentsDefeated{}
			}
			switch url {
			case urls[0]:
				m[parsed.ID].RankTotal = parsed.Rank
				m[parsed.ID].ScoreTotal = parsed.Score
			case urls[1]:
				m[parsed.ID].RankAtt = parsed.Rank
				m[parsed.ID].ScoreAtt = parsed.Score
			case urls[2]:
				m[parsed.ID].RankDef = parsed.Rank
				m[parsed.ID].ScoreDef = parsed.Score
			case urls[3]:
				m[parsed.ID].RankSup = parsed.Rank
				m[parsed.ID].ScoreSup = parsed.Score
			}
		}
	}
	return m, nil
}

func (d *dataLoader) parsePlayerLine(line []string) (*models.Player, error) {
	if len(line) != 6 {
		return nil, fmt.Errorf("Invalid line format (should be id,name,tribeid,villages,points,rank)")
	}

	var err error
	ex := true
	player := &models.Player{
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

func (d *dataLoader) LoadPlayers() ([]*models.Player, error) {
	url := d.baseURL + EndpointPlayer
	lines, err := d.getCSVData(url, true)
	if err != nil {
		lines, err = d.getCSVData(d.baseURL+EndpointPlayerNotGzipped, false)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get data, url %s", url)
		}
	}

	players := []*models.Player{}
	for _, line := range lines {
		player, err := d.parsePlayerLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse line, url %s, line %s", url, strings.Join(line, ","))
		}
		players = append(players, player)
	}

	return players, nil
}

func (d *dataLoader) parseTribeLine(line []string) (*models.Tribe, error) {
	if len(line) != 8 {
		return nil, fmt.Errorf("Invalid line format (should be id,name,tag,members,villages,points,allpoints,rank)")
	}

	var err error
	ex := true
	tribe := &models.Tribe{
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

func (d *dataLoader) LoadTribes() ([]*models.Tribe, error) {
	url := d.baseURL + EndpointTribe
	lines, err := d.getCSVData(url, true)
	if err != nil {
		lines, err = d.getCSVData(d.baseURL+EndpointTribeNotGzipped, false)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot to get data, url %s", url)
		}
	}
	tribes := []*models.Tribe{}
	for _, line := range lines {
		tribe, err := d.parseTribeLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse line, url %s, line %s", url, strings.Join(line, ","))
		}
		tribes = append(tribes, tribe)
	}
	return tribes, nil
}

func (d *dataLoader) parseVillageLine(line []string) (*models.Village, error) {
	if len(line) != 7 {
		return nil, fmt.Errorf("Invalid line format (should be id,name,x,y,playerID,points,bonus)")
	}
	var err error
	village := &models.Village{}
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

func (d *dataLoader) LoadVillages() ([]*models.Village, error) {
	url := d.baseURL + EndpointVillage
	lines, err := d.getCSVData(url, true)
	if err != nil {
		lines, err = d.getCSVData(d.baseURL+EndpointVillageNotGzipped, false)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get data, url %s", url)
		}
	}
	villages := []*models.Village{}
	for _, line := range lines {
		village, err := d.parseVillageLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse line, url %s, line %s", url, strings.Join(line, ","))
		}
		villages = append(villages, village)
	}
	return villages, nil
}

func (d *dataLoader) parseEnnoblementLine(line []string) (*models.Ennoblement, error) {
	if len(line) != 4 {
		return nil, fmt.Errorf("Invalid line format (should be village_id,timestamp,new_owner_id,old_owner_id)")
	}
	var err error
	ennoblement := &models.Ennoblement{}
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
	EnnobledAtGTE time.Time
}

func (d *dataLoader) LoadEnnoblements(cfg *LoadEnnoblementsConfig) ([]*models.Ennoblement, error) {
	if cfg == nil {
		cfg = &LoadEnnoblementsConfig{}
	}
	yesterdayDate := time.Now().Add(-23 * time.Hour)
	url := d.baseURL + EndpointConquer
	compressed := true
	if cfg.EnnobledAtGTE.After(yesterdayDate) || cfg.EnnobledAtGTE.Equal(yesterdayDate) {
		compressed = false
		url = d.baseURL + fmt.Sprintf(EndpointGetConquer, cfg.EnnobledAtGTE.Unix())
	}
	lines, err := d.getCSVData(url, compressed)
	if err != nil && compressed {
		lines, err = d.getCSVData(d.baseURL+EndpointConquerNotGzipped, false)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get data, url %s", url)
	}

	ennoblements := []*models.Ennoblement{}
	for _, line := range lines {
		ennoblement, err := d.parseEnnoblementLine(line)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot parse line, url %s, line %s", url, strings.Join(line, ","))
		}
		if ennoblement.EnnobledAt.After(cfg.EnnobledAtGTE) {
			ennoblements = append(ennoblements, ennoblement)
		}
	}
	return ennoblements, nil
}

func (d *dataLoader) GetConfig() (*models.ServerConfig, error) {
	url := d.baseURL + EndpointConfig
	cfg := &models.ServerConfig{}
	err := d.getXML(url, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getConfig")
	}
	return cfg, nil
}

func (d *dataLoader) GetBuildingConfig() (*models.BuildingConfig, error) {
	url := d.baseURL + EndpointBuildingConfig
	cfg := &models.BuildingConfig{}
	err := d.getXML(url, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getBuildingConfig")
	}
	return cfg, nil
}

func (d *dataLoader) GetUnitConfig() (*models.UnitConfig, error) {
	url := d.baseURL + EndpointUnitConfig
	cfg := &models.UnitConfig{}
	err := d.getXML(url, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "getUnitConfig")
	}
	return cfg, nil
}
