package device

import (
	"reflect"
	"testing"

	"reconya-ai/models"
)

func TestSortDevicesByIP_InvalidIPs(t *testing.T) {
	devices := []models.Device{
		{IPv4: "invalid"},
		{IPv4: "10.0.0.2"},
		{IPv4: "10.0.0.1"},
	}

	sortDevicesByIP(devices)

	got := []string{devices[0].IPv4, devices[1].IPv4, devices[2].IPv4}
	want := []string{"10.0.0.1", "10.0.0.2", "invalid"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortDevicesByIP() = %v, want %v", got, want)
	}
}
