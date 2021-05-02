package twmodel

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const CoordsSeparator = "|"

type Coords struct {
	X int
	Y int
}

func ParseCoords(coords string) (*Coords, error) {
	splitted := strings.Split(coords, CoordsSeparator)
	if len(splitted) != 2 {
		return nil, errors.Errorf("%s: invalid format (should be number|number)", coords)
	}
	x, err := strconv.Atoi(splitted[0])
	if err != nil {
		return nil, errors.Wrapf(err, "%s: the part before | isn't a number", coords)
	}
	y, err := strconv.Atoi(splitted[1])
	if err != nil {
		return nil, errors.Wrapf(err, "%s: the part after | isn't a number", coords)
	}
	return &Coords{x, y}, nil
}
