package hetzner

import (
	"time"
)

type MarketProduct struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Description     []string `json:"description"`
	Traffic         string   `json:"traffic"`
	Dist            []string `json:"dist"`
	Arch            []int    `json:"arch"`
	Lang            []string `json:"lang"`
	Cpu             string   `json:"cpu"`
	CpuBenchmark    int      `json:"cpu_benchmark"`
	MemorySize      int      `json:"memory_size"`
	HddSize         int      `json:"hdd_size"`
	HddText         string   `json:"hdd_text"`
	HddCount        int      `json:"hdd_count"`
	Datacenter      string   `json:"datacenter"`
	NetworkSpeed    string   `json:"network_speed"`
	Price           float64  `json:"price,string"`
	PriceSetup      float64  `json:"price_setup,string"`
	PriceVat        float64  `json:"price_vat,string"`
	PriceSetupVat   float64  `json:"price_setup_vat,string"`
	FixedPrice      bool     `json:"fixed_price"`
	NextReduce      int      `json:"next_reduce"`
	NextReduceDate  JSONTime `json:"next_reduce_date"`
	OrderableAddons []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Min    int    `json:"min"`
		Max    int    `json:"max"`
		Prices []struct {
			Location string `json:"location"`
			Price    struct {
				Net   float64 `json:"net,string"`
				Gross float64 `json:"gross,string"`
			} `json:"price"`
			PriceSetup struct {
				Net   float64 `json:"net,string"`
				Gross float64 `json:"gross,string"`
			} `json:"price_setup"`
		} `json:"prices"`
	} `json:"orderable_addons"`
}

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description []string `json:"description"`
	Traffic     string   `json:"traffic"`
	Dist        []string `json:"dist"`
	Arch        []int    `json:"arch"`
	Lang        []string `json:"lang"`
	Location    []string `json:"location"`
	Prices      []struct {
		Location string `json:"location"`
		Price    struct {
			Net   float64 `json:"net,string"`
			Gross float64 `json:"gross,string"`
		} `json:"price"`
		PriceSetup struct {
			Net   float64 `json:"net,string"`
			Gross float64 `json:"gross,string"`
		} `json:"price_setup"`
	} `json:"prices"`
	OrderableAddons []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Min    int    `json:"min"`
		Max    int    `json:"max"`
		Prices []struct {
			Location string `json:"location"`
			Price    struct {
				Net   float64 `json:"net,string"`
				Gross float64 `json:"gross,string"`
			} `json:"price"`
			PriceSetup struct {
				Net   float64 `json:"net,string"`
				Gross float64 `json:"gross,string"`
			} `json:"price_setup"`
		} `json:"prices"`
	} `json:"orderable_addons"`
}

type AuthorizedKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}

type HostKey struct {
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}

type Transaction struct {
	Id            string        `json:"id"`
	Date          time.Time     `json:"date"`
	Status        string        `json:"status"`
	ServerNumber  interface{}   `json:"server_number"`
	ServerIp      interface{}   `json:"server_ip"`
	AuthorizedKey []interface{} `json:"authorized_key"`
	HostKey       []interface{} `json:"host_key"`
	Comment       interface{}   `json:"comment"`
	Product       struct {
		Id          string      `json:"id"`
		Name        string      `json:"name"`
		Description []string    `json:"description"`
		Traffic     string      `json:"traffic"`
		Dist        string      `json:"dist"`
		Arch        string      `json:"arch"`
		Lang        string      `json:"lang"`
		Location    interface{} `json:"location"`
	} `json:"product"`
	Addons []string `json:"addons"`
}

type CreateTransactionRequest struct {
	ProductID     string   `url:"product_id"`
	AuthorizedKey []string `url:"authorized_key,brackets"`
	Password      string   `url:"password,omitempty"`
	Location      string   `url:"location"`
	Dist          string   `url:"dist,omitempty"`
	Arch          int      `url:"arch,omitempty"`
	Lang          string   `url:"lang,omitempty"`
	Comment       string   `url:"comment,omitempty"`
	Addon         []string `url:"addon,brackets"`
	Test          bool     `url:"test"`
}

type dataTransaction struct {
	Transaction *Transaction `json:"transaction"`
}

type dataProduct struct {
	Product *Product `json:"product"`
}
type dataMarketProduct struct {
	Product *MarketProduct `json:"product"`
}
