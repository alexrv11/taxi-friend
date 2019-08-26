package utils

import "math"

const RadioEarth = 6371

func DistanceInKmBetweenEarthCoordinates(lat1, long1, lat2, long2 float64)  float64 {

	dLat := toRadians(lat2-lat1)
	dLon := toRadians(long2-long1)

	lat1 = toRadians(lat1)
	lat2 = toRadians(lat2)

	a := math.Sin(dLat/2) * math.Sin(dLat/2) + math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return RadioEarth * c
}

func toRadians(angle float64) float64 {

	return angle / (math.Pi/180)
}
