package hetzner

// Traffic contains information about traffic data
type Traffic struct {
	Type string                    `json:"type"`
	From string                    `json:"from"`
	To   string                    `json:"to"`
	Data map[string]TrafficDataRow `json:"data"`
}

// TrafficDataRow contains information about traffic data row
type TrafficDataRow struct {
	In  float64 `json:"in"`
	Out float64 `json:"out"`
	Sum float64 `json:"sum"`
}

// TrafficGroup contains information about traffic data
type TrafficGroup struct {
	Type string                               `json:"type"`
	From string                               `json:"from"`
	To   string                               `json:"to"`
	Data map[string]map[string]TrafficDataRow `json:"data"`
}

// TrafficRequest Requested information for TrafficService.Get and TrafficService.GetByGroup
type TrafficRequest struct {
	IP     []string `url:"ip,brackets"`     // One or more IP addresses
	Subnet []string `url:"subnet,brackets"` // One or more subnet addresses
	/*
		Date/Time from
		Query type "day": [YYYY]-[MM]-[DD]T[H], e.g. 2010-09-01T00
		Query type "month": [YYYY]-[MM]-[DD], e.g. 2010-09-01
		Query type "year": [YYYY]-[MM], e.g. 2010-09
	*/
	From string `url:"from"`
	/*
		Date/Time to
		Query type "day": [YYYY]-[MM]-[DD]T[H], e.g. 2010-09-01T00
		Query type "month": [YYYY]-[MM]-[DD], e.g. 2010-09-01
		Query type "year": [YYYY]-[MM], e.g. 2010-09
	*/
	To           string `url:"to"`
	Type         string `url:"type"`           // For type "day" the data is grouped by hours, for type "month" by days and for type "year" by months.
	SingleValues bool   `url:"single_values" ` // If set to "true" the traffic data is returned not as a sum over the whole interval but grouped by hour, day or month (optional)
}

type dataTraffic struct {
	Traffic *Traffic `json:"traffic"`
}

type dataTrafficGroup struct {
	Traffic *TrafficGroup `json:"traffic"`
}
