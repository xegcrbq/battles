package converter

import (
	"battles/internal/utils/logger"
	"math"
)

// Int64ToFloat64 преобразует из условных сатоши в биткоины
func Int64ToFloat64(i int64) (f float64) {
	f = float64(i%int64(math.Pow10(8)))*math.Pow10(-8) + float64(i/int64(int(math.Pow10(8))))
	return
}

// Float64ToInt64 преобразует из условных , биткоинов в сатоши
func Float64ToInt64(f float64) (i int64) {
	i = int64(math.Round((f-math.Floor(f))*math.Pow10(8))) + int64(math.Floor(f))*int64(math.Pow10(8))
	if i == 0 && f != 0 {
		logger.Get().Debug("Error Float64ToInt64 convert")
	}
	return
}
