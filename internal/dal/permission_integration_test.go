// +build integration  

package dal

import (
	"testing"

	"github.com/powerman/check"
)
func TestGetPermissionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		getted, err := testRepo.GetPermission(testPermission1.ID, isolatedEntityID)
		t.Nil(err)
		t.DeepEqual(getted, testPermission1)

	t.Nil(testRepo.truncate())
}
func TestAddPermissionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditPermissionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.EditPermission(testPermission1.ID, isolatedEntityID, testPermission2))

	t.Nil(testRepo.truncate())
}
func TestDeletePermissionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.DeletePermission(testPermission1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListPermissionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		list, _,  err := testRepo.ListPermission(isolatedEntityID, listParams)
		t.Nil(err)
		t.DeepEqual(list, testPermissions)

	t.Nil(testRepo.truncate())
}