package attendance

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func ptrInt64ToPtrUint64(v *int64) *uint64 {
	if v == nil {
		return nil
	}
	u := uint64(*v)
	return &u
}

func nullTimeToPtr(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

func NullInt64ToUint64Ptr(ni sql.NullInt64) *uint64 {
	if !ni.Valid {
		return nil
	}
	// Ambil nilai Int64-nya, cast ke uint64, lalu simpan di variabel baru
	val := uint64(ni.Int64)
	return &val
}

func nullStringToPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case []byte:
		parsed, err := strconv.ParseFloat(string(val), 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse float: %w", err)
		}
		return parsed, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unexpected type for float: %T", v)
	}
}
