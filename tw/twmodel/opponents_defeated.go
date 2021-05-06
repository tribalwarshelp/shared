package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
	"github.com/go-pg/pg/v10/orm"
)

type OpponentsDefeated struct {
	RankAtt    int `json:"rankAtt" pg:",use_zero" gqlgen:"rankAtt"`
	ScoreAtt   int `pg:",use_zero" json:"scoreAtt" gqlgen:"scoreAtt"`
	RankDef    int `pg:",use_zero" json:"rankDef" gqlgen:"rankDef"`
	ScoreDef   int `pg:",use_zero" json:"scoreDef" gqlgen:"scoreDef"`
	RankSup    int `pg:",use_zero" json:"rankSup,omitempty" gqlgen:"rankSup"`
	ScoreSup   int `pg:",use_zero" json:"scoreSup,omitempty" gqlgen:"scoreSup"`
	RankTotal  int `pg:",use_zero" json:"rankTotal" gqlgen:"rankTotal"`
	ScoreTotal int `pg:",use_zero" json:"scoreTotal" gqlgen:"scoreTotal"`
}

type OpponentsDefeatedFilter struct {
	RankAtt     int `json:"rankAtt" gqlgen:"rankAtt"`
	RankAttGT   int `json:"rankAttGT" gqlgen:"rankAttGT"`
	RankAttGTE  int `json:"rankAttGTE" gqlgen:"rankAttGTE"`
	RankAttLT   int `json:"rankAttLT" gqlgen:"rankAttLT"`
	RankAttLTE  int `json:"rankAttLTE" gqlgen:"rankAttLTE"`
	ScoreAtt    int `json:"scoreAtt" gqlgen:"scoreAtt"`
	ScoreAttGT  int `json:"scoreAttGT" gqlgen:"scoreAttGT"`
	ScoreAttGTE int `json:"scoreAttGTE" gqlgen:"scoreAttGTE"`
	ScoreAttLT  int `json:"scoreAttLT" gqlgen:"scoreAttLT"`
	ScoreAttLTE int `json:"scoreAttLTE" gqlgen:"scoreAttLTE"`

	RankDef     int `json:"rankDef" gqlgen:"rankDef"`
	RankDefGT   int `json:"rankDefGT" gqlgen:"rankDefGT"`
	RankDefGTE  int `json:"rankDefGTE" gqlgen:"rankDefGTE"`
	RankDefLT   int `json:"rankDefLT" gqlgen:"rankDefLT"`
	RankDefLTE  int `json:"rankDefLTE" gqlgen:"rankDefLTE"`
	ScoreDef    int `json:"scoreDef" gqlgen:"scoreDef"`
	ScoreDefGT  int `json:"scoreDefGT" gqlgen:"scoreDefGT"`
	ScoreDefGTE int `json:"scoreDefGTE" gqlgen:"scoreDefGTE"`
	ScoreDefLT  int `json:"scoreDefLT" gqlgen:"scoreDefLT"`
	ScoreDefLTE int `json:"scoreDefLTE" gqlgen:"scoreDefLTE"`

	RankSup     int `json:"rankSup,omitempty" gqlgen:"rankSup"`
	RankSupGT   int `json:"rankSupGT,omitempty" gqlgen:"rankSupGT"`
	RankSupGTE  int `json:"rankSupGTE,omitempty" gqlgen:"rankSupGTE"`
	RankSupLT   int `json:"rankSupLT,omitempty" gqlgen:"rankSupLT"`
	RankSupLTE  int `json:"rankSupLTE,omitempty" gqlgen:"rankSupLTE"`
	ScoreSup    int `json:"scoreSup,omitempty" gqlgen:"scoreSup"`
	ScoreSupGT  int `json:"scoreSupGT,omitempty" gqlgen:"scoreSupGT"`
	ScoreSupGTE int `json:"scoreSupGTE,omitempty" gqlgen:"scoreSupGTE"`
	ScoreSupLT  int `json:"scoreSupLT,omitempty" gqlgen:"scoreSupLT"`
	ScoreSupLTE int `json:"scoreSupLTE,omitempty" gqlgen:"scoreSupLTE"`

	RankTotal     int `json:"rankTotal" gqlgen:"rankTotal"`
	RankTotalGT   int `json:"rankTotalGT" gqlgen:"rankTotalGT"`
	RankTotalGTE  int `json:"rankTotalGTE" gqlgen:"rankTotalGTE"`
	RankTotalLT   int `json:"rankTotalLT" gqlgen:"rankTotalLT"`
	RankTotalLTE  int `json:"rankTotalLTE" gqlgen:"rankTotalLTE"`
	ScoreTotal    int `json:"scoreTotal" gqlgen:"scoreTotal"`
	ScoreTotalGT  int `json:"scoreTotalGT" gqlgen:"scoreTotalGT"`
	ScoreTotalGTE int `json:"scoreTotalGTE" gqlgen:"scoreTotalGTE"`
	ScoreTotalLT  int `json:"scoreTotalLT" gqlgen:"scoreTotalLT"`
	ScoreTotalLTE int `json:"scoreTotalLTE" gqlgen:"scoreTotalLTE"`
}

func (f *OpponentsDefeatedFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	if !isZero(f.RankAtt) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("rank_att", alias), f.RankAtt)
	}
	if !isZero(f.RankAttGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("rank_att", alias), f.RankAttGT)
	}
	if !isZero(f.RankAttGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("rank_att", alias), f.RankAttGTE)
	}
	if !isZero(f.RankAttLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("rank_att", alias), f.RankAttLT)
	}
	if !isZero(f.RankAttLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("rank_att", alias), f.RankAttLTE)
	}
	if !isZero(f.ScoreAtt) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("score_att", alias), f.ScoreAtt)
	}
	if !isZero(f.ScoreAttGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("score_att", alias), f.ScoreAttGT)
	}
	if !isZero(f.ScoreAttGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("score_att", alias), f.ScoreAttGTE)
	}
	if !isZero(f.ScoreAttLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("score_att", alias), f.ScoreAttLT)
	}
	if !isZero(f.ScoreAttLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("score_att", alias), f.ScoreAttLTE)
	}

	if !isZero(f.RankDef) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("rank_def", alias), f.RankDef)
	}
	if !isZero(f.RankDefGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("rank_def", alias), f.RankDefGT)
	}
	if !isZero(f.RankDefGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("rank_def", alias), f.RankDefGTE)
	}
	if !isZero(f.RankDefLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("rank_def", alias), f.RankDefLT)
	}
	if !isZero(f.RankDefLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("rank_def", alias), f.RankDefLTE)
	}
	if !isZero(f.ScoreDef) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("score_def", alias), f.ScoreDef)
	}
	if !isZero(f.ScoreDefGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("score_def", alias), f.ScoreDefGT)
	}
	if !isZero(f.ScoreDefGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("score_def", alias), f.ScoreDefGTE)
	}
	if !isZero(f.ScoreDefLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("score_def", alias), f.ScoreDefLT)
	}
	if !isZero(f.ScoreDefLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("score_def", alias), f.ScoreDefLTE)
	}

	if !isZero(f.RankSup) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("rank_sup", alias), f.RankSup)
	}
	if !isZero(f.RankSupGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("rank_sup", alias), f.RankSupGT)
	}
	if !isZero(f.RankSupGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("rank_sup", alias), f.RankSupGTE)
	}
	if !isZero(f.RankSupLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("rank_sup", alias), f.RankSupLT)
	}
	if !isZero(f.RankSupLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("rank_sup", alias), f.RankSupLTE)
	}
	if !isZero(f.ScoreSup) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("score_sup", alias), f.ScoreSup)
	}
	if !isZero(f.ScoreSupGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("score_sup", alias), f.ScoreSupGT)
	}
	if !isZero(f.ScoreSupGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("score_sup", alias), f.ScoreSupGTE)
	}
	if !isZero(f.ScoreSupLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("score_sup", alias), f.ScoreSupLT)
	}
	if !isZero(f.ScoreSupLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("score_sup", alias), f.ScoreSupLTE)
	}

	if !isZero(f.RankTotal) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("rank_total", alias), f.RankTotal)
	}
	if !isZero(f.RankTotalGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("rank_total", alias), f.RankTotalGT)
	}
	if !isZero(f.RankTotalGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("rank_total", alias), f.RankTotalGTE)
	}
	if !isZero(f.RankTotalLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("rank_total", alias), f.RankTotalLT)
	}
	if !isZero(f.RankTotalLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("rank_total", alias), f.RankTotalLTE)
	}
	if !isZero(f.ScoreTotal) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("score_total", alias), f.ScoreTotal)
	}
	if !isZero(f.ScoreTotalGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("score_total", alias), f.ScoreTotalGT)
	}
	if !isZero(f.ScoreTotalGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("score_total", alias), f.ScoreTotalGTE)
	}
	if !isZero(f.ScoreTotalLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("score_total", alias), f.ScoreTotalLT)
	}
	if !isZero(f.ScoreTotalLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("score_total", alias), f.ScoreTotalLTE)
	}

	return q, nil
}
