package tfe

import (
	"reflect"
	"testing"
)

func testResourceTfeWorkspaceStateDataV0() map[string]interface{} {
	return map[string]interface{}{
		"id":          "hashicorp/test",
		"external_id": "ws-123",
	}
}

func testResourceTfeWorkspaceStateDataV1() map[string]interface{} {
	v0 := testResourceTfeWorkspaceStateDataV0()
	return map[string]interface{}{
		"id":          v0["external_id"],
		"external_id": v0["external_id"],
	}
}

func TestResourceTfeWorkspaceStateUpgradeV0(t *testing.T) {
	expected := testResourceTfeWorkspaceStateDataV1()
	actual, err := resourceTfeWorkspaceStateUpgradeV0(testResourceTfeWorkspaceStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
