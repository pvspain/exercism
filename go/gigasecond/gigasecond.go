package gigasecond

import . "time"

func AddGigasecond(t Time) Time {
  return t.Add(Second * 1000*1000*1000)
}
var location, error = LoadLocation("Local")
var Birthday = Date(1964, April, 10, 18, 0, 0, 0, location)