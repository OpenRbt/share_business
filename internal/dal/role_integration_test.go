// +build integration  

package dal

import (
	"testing"

	"github.com/powerman/check"
)
func TestGetRoleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		getted, err := testRepo.GetRole(testRole1.ID, isolatedEntityID)
		t.Nil(err)
		t.DeepEqual(getted, testRole1)

	t.Nil(testRepo.truncate())
}
func TestAddRoleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditRoleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.EditRole(testRole1.ID, isolatedEntityID, testRole2))

	t.Nil(testRepo.truncate())
}
func TestDeleteRoleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.DeleteRole(testRole1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListRoleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		list, _,  err := testRepo.ListRole(isolatedEntityID, listParams)
		t.Nil(err)
		t.DeepEqual(list, testRoles)

	t.Nil(testRepo.truncate())
}