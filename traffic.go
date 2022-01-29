package hetzner

import (
	"net/http"
)

// API: https://robot.your-server.de/doc/webservice/en.html#subnet

// TrafficService represents a service to work with Traffic.
// Query traffic data of IPs and subnets. There are three query types: "day", "month" and "year". With "day" you can query hourly aggregated traffic data within a day. With "month" you can query daily aggregated data within a month. And with "year" it is possible to get monthly aggregated data within a year.
//
// Please note that the traffic data is only available once the specified hour, day or month has already passed.
//
// The interval is given with the parameters "from" and "to" with the following syntax:
//
//    Query type "day": [YYYY]-[MM]-[DD]T[H], e.g. 2010-09-01T00
//    Query type "month": [YYYY]-[MM]-[DD], e.g. 2010-09-01
//    Query type "year": [YYYY]-[MM], e.g. 2010-09
//
// When using the query type "day", you must specify a date combined with a time value (hour). The hour value is separated from the date with the letter "T". When using the query type "month", you must specify a date without the time value. With query type "year", you must additionally leave out the day value.
//
// IP addresses or subnets without traffic data are omitted in the response.
//
// Using the parameter "single_values" it is possible to get the traffic data grouped by hours, days or month over the specified interval. For type "day" the data is grouped by hours, for type "month" by days and for type "year" by months.
type TrafficService interface {
	// Get Query traffic data for multiple IPs or Subnet.
	// See: https://robot.your-server.de/doc/webservice/en.html#post-traffic
	Get(req *TrafficRequest) (*Traffic, *http.Response, error)
	// GetByGroup Query traffic data for multiple IPs or Subnet, traffic data is returned not as a sum over the whole
	// interval but grouped by hour, day or month
	// See: https://robot.your-server.de/doc/webservice/en.html#post-traffic
	GetByGroup(req *TrafficRequest) (*TrafficGroup, *http.Response, error)
}

type TrafficServiceImpl struct {
	client *Client
}

var _ TrafficService = &TrafficServiceImpl{}

func (s *TrafficServiceImpl) Get(req *TrafficRequest) (*Traffic, *http.Response, error) {
	path := "/traffic"

	req.SingleValues = false
	data := dataTraffic{}
	resp, err := s.client.Call(http.MethodGet, path, req, &data)

	return data.Traffic, resp, err
}

func (s *TrafficServiceImpl) GetByGroup(req *TrafficRequest) (*TrafficGroup, *http.Response, error) {
	path := "/traffic"

	req.SingleValues = true
	data := dataTrafficGroup{}
	resp, err := s.client.Call(http.MethodGet, path, req, &data)

	return data.Traffic, resp, err
}
