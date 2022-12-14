// Code generated by mtgroup-generator.
package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func TestGetRole(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	rulesSet := NewMockRulesSet(ctrl)

	a := New(mockRepo, rulesSet)

	rulesSet.EXPECT().GetRoleAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().GetRole(gomock.Any(), gomock.Any()).Return(testRole1, nil)
	b, err := a.GetRole(profile, testRole1.ID)
	t.Nil(err)
	t.DeepEqual(testRole1, b)
}

func TestAddRole(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	rulesSet := NewMockRulesSet(ctrl)

	a := New(mockRepo, rulesSet)

	rulesSet.EXPECT().AddRoleAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().AddRole(gomock.Any(), gomock.Any(), gomock.Any()).Return(testRole1, nil)
	b, err := a.AddRole(profile, testRole1)
	t.Nil(err)
	t.DeepEqual(testRole1, b)
}

func TestEditRole(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	rulesSet := NewMockRulesSet(ctrl)

	a := New(mockRepo, rulesSet)

	rulesSet.EXPECT().EditRoleAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().EditRole(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.EditRole(profile, testRole1.ID, testRole1)
	t.Nil(err)
}

func TestDeleteRole(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	rulesSet := NewMockRulesSet(ctrl)

	a := New(mockRepo, rulesSet)

	rulesSet.EXPECT().DeleteRoleAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().DeleteRole(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.DeleteRole(profile, testRole1.ID)
	t.Nil(err)
}

func TestListRole(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)
	rulesSet := NewMockRulesSet(ctrl)

	a := New(mockRepo, rulesSet)

	rulesSet.EXPECT().ListRoleAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().ListRole(gomock.Any(), gomock.Any()).Return(testRoles, []string{}, nil)
	b, _, err := a.ListRole(profile, listParams)
	t.Nil(err)
	t.DeepEqual(testRoles, b)
}
