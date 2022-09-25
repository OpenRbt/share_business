//go:build integration
// +build integration

package dal

import (
	"testing"

	"github.com/powerman/check"
)

func TestGetWashServerSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	getted, err := testRepo.GetWashServer(testWashServer1.ID, isolatedEntityID)
	t.Nil(err)
	t.DeepEqual(getted, testWashServer1)

	t.Nil(testRepo.truncate())
}
func TestAddWashServerSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditWashServerSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.EditWashServer(testWashServer1.ID, isolatedEntityID, testWashServer2))

	t.Nil(testRepo.truncate())
}
func TestDeleteWashServerSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.DeleteWashServer(testWashServer1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListWashServerSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	list, _, err := testRepo.ListWashServer(isolatedEntityID, listParams)
	t.Nil(err)
	t.DeepEqual(list, testWashServers)

	t.Nil(testRepo.truncate())
}
