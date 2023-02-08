package main

import (
	"fmt"
)

type RecordCount struct{
	HostedZoneCount int `json:"HostedZoneCount,omitempty"`
}

type Route53Records struct {
	HostedZones []HostedZones `json:"HostedZones,omitempty"`
}

type HostedZones struct {
	Id                     string  `json:"Id,omitempty"`
	Name                   string  `json:"Name,omitempty"`
	CallerReference        string  `json:"CallerReference,omitempty"`
	Config                 *Config `json:"Config,omitempty"`
	ResourceRecordSetCount int     `json:"ResourceRecordSetCount,omitempty"`
}

type Config struct {
	PrivateZone bool `json:"PrivateZone,omitempty"`
}

type CheckDomainAvailability struct{
	Availability string `json:"Availability,omitempty"`
}

//----------------------------- START ---------------------------
func main() {

	fmt.Println("Welcome to Domain-checker")
	fmt.Println("")

	GetRoute53RecordDetails()
}