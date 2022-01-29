package hetzner

import (
	"testing"
	"time"
)

func TestTrafficGetDay(t *testing.T) {
	traffic, _, err := client.Traffic.Get(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   time.Now().Format("2006-01-02") + "T00",
			To:     time.Now().Format("2006-01-02") + "T24",
			Type:   "day",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data))
}

func TestTrafficGetByGroupDay(t *testing.T) {
	traffic, _, err := client.Traffic.GetByGroup(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   time.Now().Format("2006-01-02") + "T00",
			To:     time.Now().Format("2006-01-02") + "T24",
			Type:   "day",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data[testingServer.ServerIP]))
}

func TestTrafficGetMonth(t *testing.T) {
	traffic, _, err := client.Traffic.Get(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   "2022-01-01",
			To:     time.Now().Format("2006-01-02"),
			Type:   "month",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data))
}

func TestTrafficGetByGroupMonth(t *testing.T) {
	traffic, _, err := client.Traffic.GetByGroup(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   "2022-01-01",
			To:     time.Now().Format("2006-01-02"),
			Type:   "month",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data[testingServer.ServerIP]))
}

func TestTrafficGetYear(t *testing.T) {
	traffic, _, err := client.Traffic.Get(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   "2021-01",
			To:     "2021-12",
			Type:   "year",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data))
}

func TestTrafficGetByGroupYear(t *testing.T) {
	traffic, _, err := client.Traffic.GetByGroup(
		&TrafficRequest{
			IP:     []string{testingServer.ServerIP},
			Subnet: nil,
			From:   "2021-01",
			To:     "2021-12",
			Type:   "year",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Traffic Data: %d", len(traffic.Data[testingServer.ServerIP]))
}
