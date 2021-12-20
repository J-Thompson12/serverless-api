package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMapData(t *testing.T) {
	testData := `OrgAbuseEmail:  cvadmin@directcom.com
	OrgAbuseRef:    https://rdap.arin.net/registry/entity/NETWO5129-ARIN
	
	
	#
	# ARIN WHOIS data and services are subject to the Terms of Use
	# available at: https://www.arin.net/resources/registry/whois/tou/
	#
	# If you see inaccuracies in the results, please report at
	# https://www.arin.net/resources/registry/whois/inaccuracy_reporting/
	#
	# Copyright 1997-2021, American Registry for Internet Numbers, Ltd.
	#`

	result := createMapData(testData)
	require.Equal(t, 2, len(result))

}

func TestBuildIP(t *testing.T) {
	domainMap := map[string]string{"Domain Name": "test domain", "Registrant City": "LA"}
	orgMap := map[string]string{"OrgName": "test org", "City": "San Diego"}

	cases := []struct {
		input    map[string]string
		expected IPData
	}{
		{domainMap, IPData{Name: "test domain", City: "LA", StateProv: "", Email: ""}},
		{orgMap, IPData{Name: "test org", City: "San Diego", StateProv: "", Email: ""}},
	}
	for _, c := range cases {
		result := buildIPData(c.input)
		require.Equal(t, c.expected, result)
	}

}

func TestIsDomain(t *testing.T) {
	domainMap := map[string]string{"Domain Name": "test"}
	noDomainMap := map[string]string{"OrgName": "test"}
	cases := []struct {
		input    map[string]string
		expected bool
	}{
		{domainMap, true},
		{noDomainMap, false},
	}
	for _, c := range cases {
		result := isDomain(c.input)
		require.Equal(t, c.expected, result)
	}
}
