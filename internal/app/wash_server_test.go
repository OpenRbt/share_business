// Code generated by mtgroup-generator.
package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func TestGetWashServer(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	mockRepo.EXPECT().GetWashServer(gomock.Any(), gomock.Any()).Return(testWashServer1, nil)
	b, err := a.GetWashServer(profile, testWashServer1.ID)
	t.Nil(err)
	t.DeepEqual(testWashServer1, b)
}

func TestAddWashServer(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	mockRepo.EXPECT().AddWashServer(gomock.Any(), gomock.Any(), gomock.Any()).Return(testWashServer1, nil)
	b, err := a.AddWashServer(profile, testWashServer1)
	t.Nil(err)
	t.DeepEqual(testWashServer1, b)
}

func TestEditWashServer(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	mockRepo.EXPECT().EditWashServer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.EditWashServer(profile, testWashServer1.ID, testWashServer1)
	t.Nil(err)
}

func TestDeleteWashServer(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	mockRepo.EXPECT().DeleteWashServer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.DeleteWashServer(profile, testWashServer1.ID)
	t.Nil(err)
}

func TestListWashServer(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	mockRepo.EXPECT().ListWashServer(gomock.Any(), gomock.Any()).Return(testWashServers, []string{}, nil)
	b, _, err := a.ListWashServer(profile, listParams)
	t.Nil(err)
	t.DeepEqual(testWashServers, b)
}
