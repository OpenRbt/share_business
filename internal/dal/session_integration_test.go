//go:build integration
// +build integration

package dal

import (
	"testing"

	"github.com/powerman/check"
)

func TestGetSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	getted, err := testRepo.GetSession(testSession1.ID, isolatedEntityID)
	t.Nil(err)
	t.DeepEqual(getted, testSession1)

	t.Nil(testRepo.truncate())
}
func TestAddSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.EditSession(testSession1.ID, isolatedEntityID, testSession2))

	t.Nil(testRepo.truncate())
}
func TestDeleteSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.DeleteSession(testSession1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestListSessionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	list, _, err := testRepo.ListSession(isolatedEntityID, listParams)
	t.Nil(err)
	t.DeepEqual(list, testSessions)

	t.Nil(testRepo.truncate())
}
