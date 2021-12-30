package weather

import (
	"time"
)

// 24 phases per cycle, 1 cycle per week
func phase() int64 {
	return time.Now().Unix() / 25200;
}

// rough boilerplate for actual implementation
// assuming worldwrap & value wraparound works ok
func simulate(x, y, z uint8) uint8 {
	zone uint32 = x + y*256 + z*65536;
	forecast := getWeather(zone);
	// add function to decode conditions from forecast?

	humidity, temperature, wind uint8;
	// randomly modify one of the attributes (or not at all) to see next weather pattern?
}

