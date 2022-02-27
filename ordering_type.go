package hetzner

import (
	"time"
)

// MarketProduct contains information about auction server.
type MarketProduct struct {
	ID              int               `json:"id"`
	Name            string            `json:"name"`
	Description     []string          `json:"description"`
	Traffic         string            `json:"traffic"`
	Dist            []string          `json:"dist"`
	Arch            []int             `json:"arch"`
	Lang            []string          `json:"lang"`
	Cpu             string            `json:"cpu"`
	CpuBenchmark    int               `json:"cpu_benchmark"`
	MemorySize      int               `json:"memory_size"`
	HddSize         int               `json:"hdd_size"`
	HddText         string            `json:"hdd_text"`
	HddCount        int               `json:"hdd_count"`
	Datacenter      string            `json:"datacenter"`
	NetworkSpeed    string            `json:"network_speed"`
	Price           float64           `json:"price,string"`
	PriceSetup      float64           `json:"price_setup,string"`
	PriceVat        float64           `json:"price_vat,string"`
	PriceSetupVat   float64           `json:"price_setup_vat,string"`
	FixedPrice      bool              `json:"fixed_price"`
	NextReduce      int               `json:"next_reduce"`
	NextReduceDate  JSONTime          `json:"next_reduce_date,omitempty"`
	OrderableAddons []orderableAddons `json:"orderable_addons"`
}

type orderableAddons struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Min    int        `json:"min"`
	Max    int        `json:"max"`
	Prices []priceRow `json:"prices"`
}

type priceRow struct {
	Location   string `json:"location"`
	Price      price  `json:"price"`
	PriceSetup price  `json:"price_setup"`
}

type price struct {
	Net   float64 `json:"net,string"`
	Gross float64 `json:"gross,string"`
}

// Product information about standard server product
type Product struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Description     []string          `json:"description"`
	Traffic         string            `json:"traffic"`
	Dist            []string          `json:"dist"`
	Arch            []int             `json:"arch"`
	Lang            []string          `json:"lang"`
	Location        []string          `json:"location"`
	Prices          []priceRow        `json:"prices"`
	OrderableAddons []orderableAddons `json:"orderable_addons"`
}

// AuthorizedKey information about AuthorizedKey
type AuthorizedKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}

// HostKey information about HostKey
type HostKey struct {
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}

type transactionMeta struct {
	ID            string    `json:"id"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	ServerNumber  string    `json:"server_number"`
	ServerIp      string    `json:"server_ip"`
	AuthorizedKey []struct {
		Key struct {
			Name        string `json:"name"`
			Fingerprint string `json:"fingerprint"`
			Type        string `json:"type"`
			Size        int    `json:"size"`
		} `json:"key"`
	} `json:"authorized_key"`
	HostKey []string `json:"host_key"`
	Comment string   `json:"comment"`
	Addons  []string `json:"addons"`
}

type productMeta struct {
	Name        string   `json:"name"`
	Description []string `json:"description"`
	Traffic     string   `json:"traffic"`
	Dist        string   `json:"dist"`
	Arch        int      `json:"arch"`
	Lang        string   `json:"lang"`
}

// MarketTransaction contains information about auction server order transaction
type MarketTransaction struct {
	transactionMeta
	Product struct {
		productMeta
		Cpu          string `json:"cpu"`
		CpuBenchmark int    `json:"cpu_benchmark"`
		MemorySize   int    `json:"memory_size"`
		HddSize      int    `json:"hdd_size"`
		HddText      string `json:"hdd_text"`
		HddCount     int    `json:"hdd_count"`
		Datacenter   string `json:"datacenter"`
		NetworkSpeed string `json:"network_speed"`
	} `json:"product"`
}

// Transaction contains information about standard server order transaction
type Transaction struct {
	transactionMeta
	Product struct {
		ID string `json:"id"`
		productMeta
		Location interface{} `json:"location"`
	} `json:"product"`
}

// CreateTransactionRequest Requested information for OrderingService.CreateMarketTransaction
type CreateTransactionRequest struct {
	// Product ID
	ProductID string `url:"product_id"`
	// One or more SSH key fingerprints (Optional, you can use either parameter "authorized_key" or parameter "password")
	AuthorizedKey []string `url:"authorized_key,brackets"`
	// Root password (Optional: you can use either parameter "authorized_key" or parameter "password")
	Password string `url:"password,omitempty"`
	// The desired location
	Location string `url:"location"`
	// Distribution name which should be preinstalled (optional)
	Dist string `url:"dist,omitempty"`
	// Architecture of preinstalled distribution (optional)
	Arch int `url:"arch,omitempty"`
	// Language of preinstalled distribution (optional)
	Lang string `url:"lang,omitempty"`
	// Order comment (optional); Please note that if a comment is supplied, the order will be processed manually.
	Comment string `url:"comment,omitempty"`
	// Array of addon IDs (optional)
	Addon []string `url:"addon,brackets"`
	// The order will not be processed if set to "true" (optional)
	Test bool `url:"test"`
}

type dataTransaction struct {
	Transaction *Transaction `json:"transaction"`
}

type dataMarketTransaction struct {
	Transaction *MarketTransaction `json:"transaction"`
}

type dataProduct struct {
	Product *Product `json:"product"`
}

type dataMarketProduct struct {
	Product *MarketProduct `json:"product"`
}
