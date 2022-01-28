package hetzner

type serverMeta struct {
	ServerIP      string `json:"server_ip"`
	ServerIpv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
	ServerName    string `json:"server_name"`
}

// ServerSummary contains information about server summary
type ServerSummary struct {
	serverMeta
	Product   string   `json:"product"`
	Dc        string   `json:"dc"`
	Traffic   string   `json:"traffic"`
	Status    string   `json:"status"`
	Cancelled bool     `json:"cancelled"`
	PaidUntil string   `json:"paid_until"`
	Ip        []string `json:"ip"`
	Subnet    []struct {
		Ip   string `json:"ip"`
		Mask string `json:"mask"`
	} `json:"subnet"`
}

// Server contains information about server
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

// UpdateServerRequest Requested information for ServerService.Update
type UpdateServerRequest struct {
	ServerNumber int    // Server ID
	ServerName   string `url:"server_name"` // Server name
}

// Cancellation contains information about cancellation server options
type Cancellation struct {
	serverMeta
	EarliestCancellationDate JSONDate `json:"earliest_cancellation_date"`
	Cancelled                bool     `json:"cancelled"`
	Reserved                 bool     `json:"reserved"`
	ReservationPossible      bool     `json:"reservation_possible"`
	CancellationDate         JSONDate `json:"cancellation_date,omitempty"`
	CancellationReason       []string `json:"cancellation_reason"`
}

// CancellationSingle contains information about a cancelled server
type CancellationSingle struct {
	serverMeta
	EarliestCancellationDate JSONDate `json:"earliest_cancellation_date"`
	Cancelled                bool     `json:"cancelled"`
	Reserved                 bool     `json:"reserved"`
	ReservationPossible      bool     `json:"reservation_possible"`
	CancellationDate         JSONDate `json:"cancellation_date,omitempty"`
	CancellationReason       string   `json:"cancellation_reason"`
}

// CancelServerRequest Requested information for ServerService.CancelServer
type CancelServerRequest struct {
	ServerNumber       int    // Server ID
	CancellationDate   string `url:"cancellation_date"`             // Date to which the server should be cancelled
	CancellationReason string `url:"cancellation_reason,omitempty"` // Cancellation reason, optional
	ReserveLocation    bool   `url:"reserve_location"`              // Whether server location shall be reserved ('true' or 'false')
}

// WithdrawOrderRequest Requested information for ServerService.WithdrawCancellation
type WithdrawOrderRequest struct {
	ServerNumber   string // Server ID
	ReversalReason string `url:"reversal_reason"` // Reason for withdrawal, optional
}

type dataCancellation struct {
	Cancellation *Cancellation `json:"cancellation"`
}

type dataCancellationSingle struct {
	Cancellation *CancellationSingle `json:"cancellation"`
}

type dataServer struct {
	Server *Server `json:"server"`
}

type dataServerSummary struct {
	Server *ServerSummary `json:"server"`
}
