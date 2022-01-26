package hetzner

type ServerSummary struct {
	ServerIP      string   `json:"server_ip"`
	ServerIpv6Net string   `json:"server_ipv6_net"`
	ServerNumber  int      `json:"server_number"`
	ServerName    string   `json:"server_name"`
	Product       string   `json:"product"`
	Dc            string   `json:"dc"`
	Traffic       string   `json:"traffic"`
	Status        string   `json:"status"`
	Cancelled     bool     `json:"cancelled"`
	PaidUntil     string   `json:"paid_until"`
	Ip            []string `json:"ip"`
	Subnet        []struct {
		Ip   string `json:"ip"`
		Mask string `json:"mask"`
	} `json:"subnet"`
}

type Server struct {
	ServerSummary
	Reset            bool `json:"reset"`
	Rescue           bool `json:"rescue"`
	Vnc              bool `json:"vnc"`
	Windows          bool `json:"windows"`
	Plesk            bool `json:"plesk"`
	Cpanel           bool `json:"cpanel"`
	Wol              bool `json:"wol"`
	HotSwap          bool `json:"hot_swap"`
	LinkedStoragebox int  `json:"linked_storagebox"`
}

type UpdateServerRequest struct {
	ServerNumber string // Server ID
	ServerName   string `url:"server_name"` // Server name
}

type Cancellation struct {
	ServerIp                 string   `json:"server_ip"`
	ServerIpv6Net            string   `json:"server_ipv6_net"`
	ServerNumber             int      `json:"server_number"`
	ServerName               string   `json:"server_name"`
	EarliestCancellationDate JSONTime `json:"earliest_cancellation_date"`
	Cancelled                bool     `json:"cancelled"`
	Reserved                 bool     `json:"reserved"`
	ReservationPossible      bool     `json:"reservation_possible"`
	CancellationDate         JSONTime `json:"cancellation_date"`
	CancellationReason       []string `json:"cancellation_reason"`
}

type CancelServerRequest struct {
	ServerNumber       string // Server ID
	CancellationDate   string `url:"cancellation_date"`             // Date to which the server should be cancelled
	CancellationReason string `url:"cancellation_reason,omitempty"` // Cancellation reason, optional
	ReserveLocation    bool   `url:"reserve_location"`              // Whether server location shall be reserved ('true' or 'false')
}

type WithdrawOrderRequest struct {
	ServerNumber   string // Server ID
	ReversalReason string `url:"reversal_reason"` // Reason for withdrawal, optional
}

type dataCancellation struct {
	Cancellation *Cancellation `json:"cancellation"`
}

type dataServer struct {
	Server *Server `json:"server"`
}

type dataServerSummary struct {
	Server *ServerSummary `json:"server"`
}
