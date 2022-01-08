package restore_ip_addresses

import (
	"reflect"
	"testing"
)

func TestRestoreIPAddresses1(t *testing.T) {
	ips := restoreIpAddresses("25525511135")
	expected := []string{"255.255.11.135", "255.255.111.35"}

	if !reflect.DeepEqual(expected, ips) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}

func TestRestoreIPAddresses2(t *testing.T) {
	ips := restoreIpAddresses("0000")
	expected := []string{"0.0.0.0"}

	if !reflect.DeepEqual(expected, ips) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}

func TestRestoreIPAddresses3(t *testing.T) {
	ips := restoreIpAddresses("101023")
	expected := []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}

	if !reflect.DeepEqual(expected, ips) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}

func TestRestoreIPAddresses4(t *testing.T) {
	ips := restoreIpAddresses("000")
	var expected []string

	if !reflect.DeepEqual(expected, ips) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}

func TestRestoreIPAddresses5(t *testing.T) {
	ips := restoreIpAddresses("010010")
	expected := []string{"0.10.0.10", "0.100.1.0"}

	if !reflect.DeepEqual(expected, ips) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}
