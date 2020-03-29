package vultr

import (
	"context"
	"reflect"
	"testing"

	cloudprovider "k8s.io/cloud-provider"
)

func TestZones_GetZoneByNodeName(t *testing.T) {
	client := newFakeClient()
	zone := newZones(client, "1")

	expected := cloudprovider.Zone{Region: "1"}
	actual, err := zone.GetZoneByNodeName(context.TODO(), "ccm-test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v got %+v", expected, actual)
	}

}

func TestZones_GetZoneByProviderID(t *testing.T) {
	client := newFakeClient()
	zone := newZones(client, "1")

	expected := cloudprovider.Zone{Region: "1"}

	actual, err := zone.GetZoneByProviderID(context.Background(), "vultr://576965")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v got %+v", expected, actual)
	}
}
