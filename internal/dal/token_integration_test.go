//go:build integration
// +build integration

package dal

import (
	"testing"

	"github.com/powerman/check"
)

func TestGetTokenSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	getted, err := testRepo.GetToken(testToken1.ID, isolatedEntityID)
	t.Nil(err)
	t.DeepEqual(getted, testToken1)

	t.Nil(testRepo.truncate())
}
func TestAddTokenSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestDeleteTokenSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
	t.Nil(testRepo.DeleteToken(testToken1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
