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
	parts := strings.Split(coords, CoordsSeparator)
	if len(parts) != 2 {
		return nil, errors.Errorf("%s: invalid format (should be number|number)", coords)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.Wrapf(err, "%s: the part before | isn't a number", coords)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.Wrapf(err, "%s: the part after | isn't a number", coords)
	}
	return &Coords{x, y}, nil
}
