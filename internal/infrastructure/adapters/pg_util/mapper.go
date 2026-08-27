package pg_util

import (
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func TimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func TimestamptzToTime(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}

	return &t.Time
}

func MeasurementToNumeric(value float32) (pgtype.Numeric, error) {
	value = float32(math.Round(float64(value)*1000) / 1000)

	var numeric pgtype.Numeric

	err := numeric.Scan(fmt.Sprintf("%.3f", value))
	if err != nil {
		return pgtype.Numeric{}, err
	}

	return numeric, nil
}

func NumericToMeasurement(value pgtype.Numeric) (float32, error) {
	if !value.Valid {
		return 0, nil
	}

	var result float64

	err := value.Scan(&result)
	if err != nil {
		return 0, err
	}

	return float32(result), nil
}

func NullableStringToText(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{
			Valid: false,
		}
	}

	return pgtype.Text{
		String: *value,
		Valid:  true,
	}
}

func TextToNullableString(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}
