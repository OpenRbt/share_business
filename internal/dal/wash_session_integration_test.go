//go:build integration
// +build integration

package dal

import (
	"testing"

	"github.com/powerman/check"
)

func TestGetWashSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	getted, err := testRepo.GetWashSession(testWashSession1.ID, isolatedEntityID)
	t.Nil(err)
	t.DeepEqual(getted, testWashSession1)

	t.Nil(testRepo.truncate())
}
func TestAddWashSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditWashSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.EditWashSession(testWashSession1.ID, isolatedEntityID, testWashSession2))

	t.Nil(testRepo.truncate())
}
func TestDeleteWashSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.DeleteWashSession(testWashSession1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListWashSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	list, _, err := testRepo.ListWashSession(isolatedEntityID, listParams)
	t.Nil(err)
	t.DeepEqual(list, testWashSessions)

	t.Nil(testRepo.truncate())
}
