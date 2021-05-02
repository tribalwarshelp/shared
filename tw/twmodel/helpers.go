package twmodel

import (
	"time"
)

func isZero(v interface{}) bool {
	switch c := v.(type) {
	case string:
		return c == ""
	case *string,
		[]string,
		[]VersionCode,
		[]ServerStatus,
		*int:
		return c == nil
	case int:
		return c == 0
	case []int:
		return c == nil
	case float64:
		return c == 0
	case *float64:
		return c == nil
	case float32:
		return c == 0
	case *float32:
		return c == nil
	case bool:
		return !c
	case *bool:
		return c == nil
	case time.Time:
		return c.IsZero()
	case *time.Time:
		return c == nil
	default:
		return false
	}
}
