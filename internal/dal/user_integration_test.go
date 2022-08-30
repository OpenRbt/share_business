//go:build integration
// +build integration

package dal

import (
	"testing"

	"github.com/powerman/check"
)

func TestGetUserSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	getted, err := testRepo.GetUser(testUser1.ID, isolatedEntityID)
	t.Nil(err)
	t.DeepEqual(getted, testUser1)

	t.Nil(testRepo.truncate())
}
func TestAddUserSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditUserSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.EditUser(testUser1.ID, isolatedEntityID, testUser2))

	t.Nil(testRepo.truncate())
}
func TestDeleteUserSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.DeleteUser(testUser1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListUserSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	list, _, err := testRepo.ListUser(isolatedEntityID, listParams)
	t.Nil(err)
	t.DeepEqual(list, testUsers)

	t.Nil(testRepo.truncate())
}
