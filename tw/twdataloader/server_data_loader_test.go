package twdataloader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/tribalwarshelp/shared/tw/twmodel"
)

func TestLoadOD(t *testing.T) {
	type scenario struct {
		respKillAll      string
		respKillAtt      string
		respKillDef      string
		respKillSup      string
		respKillAllTribe string
		respKillAttTribe string
		respKillDefTribe string
		tribe            bool
		expectedResult   map[int]*twmodel.OpponentsDefeated
		expectedErrMsg   string
	}

	scenarios := []scenario{
		{
			respKillAll:    "1,1",
			expectedErrMsg: "invalid line format (should be rank,id,score)",
		},
		{
			respKillAllTribe: "1,1",
			expectedErrMsg:   "invalid line format (should be rank,id,score)",
			tribe:            true,
		},
		{
			respKillAll:    "1,1,1",
			respKillAtt:    "1,1,1",
			respKillDef:    "1,1,1",
			respKillSup:    "1,1",
			expectedErrMsg: "invalid line format (should be rank,id,score)",
		},
		{
			respKillAllTribe: "1,1,1",
			respKillAttTribe: "1,1,1",
			respKillDefTribe: "1,1",
			expectedErrMsg:   "invalid line format (should be rank,id,score)",
			tribe:            true,
		},
		{
			respKillAll:    "1,1,asd",
			expectedErrMsg: "parsedODLine.Score: strconv.Atoi: parsing \"asd\"",
		},
		{
			respKillAll:    "1,asd,1",
			expectedErrMsg: "parsedODLine.ID: strconv.Atoi: parsing \"asd\":",
		},
		{
			respKillAll:    "asd,1,1",
			expectedErrMsg: "parsedODLine.Rank: strconv.Atoi: parsing \"asd\":",
		},
		{
			respKillAllTribe: "1,1,asd",
			expectedErrMsg:   "parsedODLine.Score: strconv.Atoi: parsing \"asd\"",
			tribe:            true,
		},
		{
			respKillAllTribe: "1,asd,1",
			expectedErrMsg:   "parsedODLine.ID: strconv.Atoi: parsing \"asd\":",
			tribe:            true,
		},
		{
			respKillAllTribe: "asd,1,1",
			expectedErrMsg:   "parsedODLine.Rank: strconv.Atoi: parsing \"asd\":",
			tribe:            true,
		},
		{
			respKillAll: "1,1,1\n4,2,4\n2,3,2",
			respKillAtt: "2,1,2\n3,2,3\n1,3,1",
			respKillDef: "3,1,3\n2,2,2\n4,3,4",
			respKillSup: "4,1,4\n1,2,1\n3,3,3",
			expectedResult: map[int]*twmodel.OpponentsDefeated{
				1: {
					RankAtt:    2,
					ScoreAtt:   2,
					RankDef:    3,
					ScoreDef:   3,
					RankSup:    4,
					ScoreSup:   4,
					RankTotal:  1,
					ScoreTotal: 1,
				},
				2: {
					RankAtt:    3,
					ScoreAtt:   3,
					RankDef:    2,
					ScoreDef:   2,
					RankSup:    1,
					ScoreSup:   1,
					ScoreTotal: 4,
					RankTotal:  4,
				},
				3: {
					RankAtt:    1,
					ScoreAtt:   1,
					RankDef:    4,
					ScoreDef:   4,
					RankSup:    3,
					ScoreSup:   3,
					ScoreTotal: 2,
					RankTotal:  2,
				},
			},
		},
		{
			respKillAllTribe: "1,1,1\n2,2,2\n3,3,3",
			respKillAttTribe: "1,1,1\n2,2,2\n3,3,3",
			respKillDefTribe: "1,1,1\n2,2,2\n3,3,3",
			expectedResult: map[int]*twmodel.OpponentsDefeated{
				1: {
					RankAtt:    1,
					ScoreAtt:   1,
					RankDef:    1,
					ScoreDef:   1,
					ScoreTotal: 1,
					RankTotal:  1,
				},
				2: {
					RankAtt:    2,
					ScoreAtt:   2,
					RankDef:    2,
					ScoreDef:   2,
					ScoreTotal: 2,
					RankTotal:  2,
				},
				3: {
					RankAtt:    3,
					ScoreAtt:   3,
					RankDef:    3,
					ScoreDef:   3,
					ScoreTotal: 3,
					RankTotal:  3,
				},
			},
			tribe: true,
		},
	}

	for _, scenario := range scenarios {
		ts := prepareTestServer(&handlers{
			killAll:      createWriteCompressedStringHandler(scenario.respKillAll),
			killAtt:      createWriteCompressedStringHandler(scenario.respKillAtt),
			killDef:      createWriteCompressedStringHandler(scenario.respKillDef),
			killSup:      createWriteCompressedStringHandler(scenario.respKillSup),
			killAllTribe: createWriteCompressedStringHandler(scenario.respKillAllTribe),
			killAttTribe: createWriteCompressedStringHandler(scenario.respKillAttTribe),
			killDefTribe: createWriteCompressedStringHandler(scenario.respKillDefTribe),
		})

		dl := NewServerDataLoader(&ServerDataLoaderConfig{
			BaseURL: ts.URL,
			Client:  ts.Client(),
		})

		res, err := dl.LoadOD(scenario.tribe)
		if scenario.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}

		if scenario.expectedResult != nil {
			assert.Len(t, res, len(scenario.expectedResult))
			for id, singleResult := range res {
				expected, ok := scenario.expectedResult[id]
				assert.True(t, ok)
				assert.NotNil(t, expected)
				assert.EqualValues(t, expected, singleResult)
			}
		}

		ts.Close()
	}
}

func TestLoadPlayers(t *testing.T) {
	type scenario struct {
		resp           string
		expectedResult []*twmodel.Player
		expectedErrMsg string
	}

	exists := true
	scenarios := []scenario{
		{
			resp:           "1,1,1,1",
			expectedErrMsg: "invalid line format (should be id,name,tribeid,villages,points,rank)",
		},
		{
			resp:           "1,name,1,500,500",
			expectedErrMsg: "invalid line format (should be id,name,tribeid,villages,points,rank)",
		},
		{
			resp:           "asd,name,1,500,500,500",
			expectedErrMsg: "player.ID: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "1,name,asd,500,500,500",
			expectedErrMsg: "player.TribeID: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "1,name,123,asd,500,500",
			expectedErrMsg: "player.TotalVillages: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "1,name,123,500,asd,500",
			expectedErrMsg: "player.Points: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "1,name,123,500,123,asd",
			expectedErrMsg: "player.Rank: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "1,name,123,500,500,asd",
			expectedErrMsg: "player.Rank: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp: "1,name,123,124,125,126\n2,name2,256,257,258,259",
			expectedResult: []*twmodel.Player{
				{
					ID:            1,
					Name:          "name",
					TribeID:       123,
					TotalVillages: 124,
					Points:        125,
					Rank:          126,
					Exists:        &exists,
				},
				{
					ID:            2,
					Name:          "name2",
					TribeID:       256,
					TotalVillages: 257,
					Points:        258,
					Rank:          259,
					Exists:        &exists,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		ts := prepareTestServer(&handlers{
			player: createWriteCompressedStringHandler(scenario.resp),
		})

		dl := NewServerDataLoader(&ServerDataLoaderConfig{
			BaseURL: ts.URL,
			Client:  ts.Client(),
		})

		res, err := dl.LoadPlayers()
		if scenario.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}

		if scenario.expectedResult != nil {
			assert.Len(t, res, len(scenario.expectedResult))
			for _, singleResult := range res {
				found := false
				var player *twmodel.Player
				for _, expected := range scenario.expectedResult {
					if expected.ID == singleResult.ID {
						found = true
						player = expected
						break
					}
				}
				assert.True(t, found)
				assert.NotNil(t, player)
				assert.EqualValues(t, player, singleResult)
			}
		}

		ts.Close()
	}
}

func TestLoadTribes(t *testing.T) {
	type scenario struct {
		resp           string
		expectedResult []*twmodel.Tribe
		expectedErrMsg string
	}

	ex := true
	scenarios := []scenario{
		{
			resp:           "1,1,1,1",
			expectedErrMsg: "invalid line format (should be id,name,tag,members,villages,points,allpoints,rank)",
		},
		{
			resp:           "1,name,1,500,500",
			expectedErrMsg: "invalid line format (should be id,name,tag,members,villages,points,allpoints,rank)",
		},
		{
			resp:           "asd,name,tag,500,500,500,500,500",
			expectedErrMsg: "tribe.ID: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,tag,asd,500,500,500,500",
			expectedErrMsg: "tribe.TotalMembers: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,tag,500,asd,500,500,500",
			expectedErrMsg: "tribe.TotalVillages: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,tag,500,500,asd,500,500",
			expectedErrMsg: "tribe.Points: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,tag,500,500,500,asd,500",
			expectedErrMsg: "tribe.AllPoints: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,tag,500,500,500,500,asd",
			expectedErrMsg: "tribe.Rank: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp: "123,name,tag,500,501,502,503,504\n1234,name2,tag2,5000,5001,5002,5003,5004",
			expectedResult: []*twmodel.Tribe{
				{
					ID:            123,
					Name:          "name",
					Tag:           "tag",
					TotalMembers:  500,
					TotalVillages: 501,
					Points:        502,
					AllPoints:     503,
					Rank:          504,
					Exists:        &ex,
				},
				{
					ID:            1234,
					Name:          "name2",
					Tag:           "tag2",
					TotalMembers:  5000,
					TotalVillages: 5001,
					Points:        5002,
					AllPoints:     5003,
					Rank:          5004,
					Exists:        &ex,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		ts := prepareTestServer(&handlers{
			tribe: createWriteCompressedStringHandler(scenario.resp),
		})

		dl := NewServerDataLoader(&ServerDataLoaderConfig{
			BaseURL: ts.URL,
			Client:  ts.Client(),
		})

		res, err := dl.LoadTribes()
		if scenario.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}

		if scenario.expectedResult != nil {
			assert.Len(t, res, len(scenario.expectedResult))
			for _, singleResult := range res {
				found := false
				var tribe *twmodel.Tribe
				for _, expected := range scenario.expectedResult {
					if expected.ID == singleResult.ID {
						found = true
						tribe = expected
						break
					}
				}
				assert.True(t, found)
				assert.NotNil(t, tribe)
				assert.EqualValues(t, tribe, singleResult)
			}
		}

		ts.Close()
	}
}

func TestLoadVillages(t *testing.T) {
	type scenario struct {
		resp           string
		expectedResult []*twmodel.Village
		expectedErrMsg string
	}

	scenarios := []scenario{
		{
			resp:           "1,1,1,1",
			expectedErrMsg: "invalid line format (should be id,name,x,y,playerID,points,bonus)",
		},
		{
			resp:           "1,name,1,500,500",
			expectedErrMsg: "invalid line format (should be id,name,x,y,playerID,points,bonus)",
		},
		{
			resp:           "asd,name,500,500,500,500,0",
			expectedErrMsg: "village.ID: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,asd,500,500,500,0",
			expectedErrMsg: "village.X: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,500,asd,500,500,0",
			expectedErrMsg: "village.Y: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,500,500,asd,500,0",
			expectedErrMsg: "village.PlayerID: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,500,500,500,asd,0",
			expectedErrMsg: "village.Points: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp:           "123,name,500,500,500,500,asd",
			expectedErrMsg: "village.Bonus: strconv.Atoi: parsing \"asd\"",
		},
		{
			resp: "123,village 1,500,501,502,503,504\n124,village 2,100,101,102,103,104",
			expectedResult: []*twmodel.Village{
				{
					ID:       123,
					Name:     "village 1",
					X:        500,
					Y:        501,
					PlayerID: 502,
					Points:   503,
					Bonus:    504,
				},
				{
					ID:       124,
					Name:     "village 2",
					X:        100,
					Y:        101,
					PlayerID: 102,
					Points:   103,
					Bonus:    104,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		ts := prepareTestServer(&handlers{
			village: createWriteCompressedStringHandler(scenario.resp),
		})

		dl := NewServerDataLoader(&ServerDataLoaderConfig{
			BaseURL: ts.URL,
			Client:  ts.Client(),
		})

		res, err := dl.LoadVillages()
		if scenario.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}

		if scenario.expectedResult != nil {
			assert.Len(t, res, len(scenario.expectedResult))
			for _, singleResult := range res {
				found := false
				var village *twmodel.Village
				for _, expected := range scenario.expectedResult {
					if expected.ID == singleResult.ID {
						found = true
						village = expected
						break
					}
				}
				assert.True(t, found)
				assert.NotNil(t, village)
				assert.EqualValues(t, village, singleResult)
			}
		}

		ts.Close()
	}
}

func TestLoadEnnoblements(t *testing.T) {
	type scenario struct {
		respConquer    string
		respGetConquer string
		cfg            *LoadEnnoblementsConfig
		expectedResult []*twmodel.Ennoblement
		expectedErrMsg string
	}

	nowMinus5h := time.Now().Add(-5 * time.Hour).Unix()
	nowMinus6h := time.Now().Add(-6 * time.Hour).Unix()
	scenarios := []scenario{
		{
			respConquer:    "1,1,1",
			expectedErrMsg: "invalid line format (should be village_id,timestamp,new_owner_id,old_owner_id)",
		},
		{
			respConquer:    "1,name,1,500,500",
			expectedErrMsg: "invalid line format (should be village_id,timestamp,new_owner_id,old_owner_id)",
		},
		{
			respConquer:    "asd,123,123,123",
			expectedErrMsg: "ennoblement.VillageID: strconv.Atoi: parsing \"asd\"",
		},
		{
			respConquer:    "123,asd,123,123",
			expectedErrMsg: "timestamp: strconv.Atoi: parsing \"asd\"",
		},
		{
			respConquer:    "123,123,asd,123",
			expectedErrMsg: "ennoblement.NewOwnerID: strconv.Atoi: parsing \"asd\"",
		},
		{
			respConquer:    "123,123,123,asd",
			expectedErrMsg: "ennoblement.OldOwnerID: strconv.Atoi: parsing \"asd\"",
		},
		{
			respConquer: "123,124,125,126\n127,128,129,130",
			expectedResult: []*twmodel.Ennoblement{
				{
					VillageID:  123,
					EnnobledAt: time.Unix(124, 0),
					NewOwnerID: 125,
					OldOwnerID: 126,
				},
				{
					VillageID:  127,
					EnnobledAt: time.Unix(128, 0),
					NewOwnerID: 129,
					OldOwnerID: 130,
				},
			},
		},
		{
			respGetConquer: "123,124,125,126\n127,128,129,130",
			cfg: &LoadEnnoblementsConfig{
				EnnobledAtGT: time.Now().Add(5 * time.Hour),
			},
			expectedResult: []*twmodel.Ennoblement{},
		},
		{
			respConquer: fmt.Sprintf("123,%d,125,126\n127,%d,129,130", nowMinus5h, nowMinus6h),
			expectedResult: []*twmodel.Ennoblement{
				{
					VillageID:  123,
					EnnobledAt: time.Unix(nowMinus5h, 0),
					NewOwnerID: 125,
					OldOwnerID: 126,
				},
				{
					VillageID:  127,
					EnnobledAt: time.Unix(nowMinus6h, 0),
					NewOwnerID: 129,
					OldOwnerID: 130,
				},
			},
		},
	}

	for _, scenario := range scenarios {
		ts := prepareTestServer(&handlers{
			conquer:    createWriteCompressedStringHandler(scenario.respConquer),
			getConquer: createWriteStringHandler(scenario.respGetConquer),
		})

		dl := NewServerDataLoader(&ServerDataLoaderConfig{
			BaseURL: ts.URL,
			Client:  ts.Client(),
		})

		res, err := dl.LoadEnnoblements(scenario.cfg)
		if scenario.expectedErrMsg != "" {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), scenario.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}

		if scenario.expectedResult != nil {
			assert.Len(t, res, len(scenario.expectedResult))
			for _, singleResult := range res {
				found := false
				var ennoblement *twmodel.Ennoblement
				for _, expected := range scenario.expectedResult {
					if expected.VillageID == singleResult.VillageID {
						found = true
						ennoblement = expected
						break
					}
				}
				assert.True(t, found)
				assert.NotNil(t, ennoblement)
				assert.EqualValues(t, ennoblement, singleResult)
			}
		}

		ts.Close()
	}
}
