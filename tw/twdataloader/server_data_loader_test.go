package twdataloader

import (
	"github.com/stretchr/testify/assert"
	"testing"

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
			respKillAll: "1,1,1\n2,2,2\n3,3,3",
			respKillAtt: "1,1,1\n2,2,2\n3,3,3",
			respKillDef: "1,1,1\n2,2,2\n3,3,3",
			respKillSup: "1,1,1\n2,2,2\n3,3,3",
			expectedResult: map[int]*twmodel.OpponentsDefeated{
				1: {
					RankAtt:    1,
					ScoreAtt:   1,
					RankDef:    1,
					ScoreDef:   1,
					RankSup:    1,
					ScoreSup:   1,
					RankTotal:  1,
					ScoreTotal: 1,
				},
				2: {
					RankAtt:    2,
					ScoreAtt:   2,
					RankDef:    2,
					ScoreDef:   2,
					RankSup:    2,
					ScoreSup:   2,
					ScoreTotal: 2,
					RankTotal:  2,
				},
				3: {
					RankAtt:    3,
					ScoreAtt:   3,
					RankDef:    3,
					ScoreDef:   3,
					RankSup:    3,
					ScoreSup:   3,
					ScoreTotal: 3,
					RankTotal:  3,
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
