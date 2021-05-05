package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"
	"reflect"
)

func isZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}

type filter interface {
	WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error)
}

type filterToAppend struct {
	filter       filter
	alias        string
	relationName string
}

func appendFilters(q *orm.Query, filtersToAppend ...filterToAppend) (*orm.Query, error) {
	tableModel := q.TableModel()
	var err error
	for _, f := range filtersToAppend {
		if f.relationName != "" && tableModel != nil {
			alias, err := gopgutil.BuildAliasFromRelationName(tableModel, f.relationName)
			if err != nil {
				return q, errors.Wrapf(err, "Couldn't build alias from relation name '%s'", f.relationName)
			}
			if join := tableModel.GetJoin(f.relationName); join == nil {
				q = q.Relation(f.relationName + "._")
			}
			q, err = f.filter.WhereWithAlias(q, alias)
			if err != nil {
				return q, errors.Wrapf(err, "Couldn't append filter for the relation '%s'", f.relationName)
			}
			continue
		}
		q, err = f.filter.WhereWithAlias(q, f.alias)
		if err != nil {
			return q, err
		}
	}

	return q, nil
}
