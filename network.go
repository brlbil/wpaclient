package wpaclient

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Network represents network data returned from "LIST_NETWORK" command
type Network struct {
	ID    int
	SSID  string
	BSSID string
	Flags []string
}

func parseNetwork(b []byte) ([]Network, error) {
	i := bytes.Index(b, []byte("\n"))
	if i > 0 {
		b = b[i:]
	}

	r := csv.NewReader(bytes.NewReader(b))
	r.Comma = '\t'
	r.FieldsPerRecord = 4

	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	nts := []Network{}
	for _, rec := range recs {
		id, err := strconv.Atoi(rec[0])
		if err != nil {
			return nil, fmt.Errorf("parse id: %w", err)
		}

		nts = append(nts, Network{
			ID:    id,
			SSID:  rec[1],
			BSSID: rec[2],
			Flags: parseFlags(rec[3]),
		})
	}

	return nts, nil
}

// AP represents Access Point data returned from "SCAN_RESULTS" commad
type AP struct {
	BSSID          net.HardwareAddr
	SSID           string
	Frequency      int
	SignalStrength int
	Flags          []string
}

func parseAP(b []byte) ([]AP, error) {
	i := bytes.Index(b, []byte("\n"))
	if i > 0 {
		b = b[i:]
	}

	r := csv.NewReader(bytes.NewReader(b))
	r.Comma = '\t'
	r.FieldsPerRecord = 5

	recs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	aps := []AP{}
	for _, rec := range recs {
		bssid, err := net.ParseMAC(rec[0])
		if err != nil {
			return nil, fmt.Errorf("parse mac: %w", err)
		}

		fr, err := strconv.Atoi(rec[1])
		if err != nil {
			return nil, fmt.Errorf("parse frequency: %w", err)
		}

		ss, err := strconv.Atoi(rec[2])
		if err != nil {
			return nil, fmt.Errorf("parse signal strength: %w", err)
		}

		aps = append(aps, AP{
			BSSID:          bssid,
			SSID:           rec[4],
			Frequency:      fr,
			SignalStrength: ss,
			Flags:          parseFlags(rec[3]),
		})
	}

	return aps, nil
}

func parseFlags(s string) []string {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	return strings.Split(s, "][")
}
