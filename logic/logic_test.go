package logic

import (
	"fmt"
	"testing"
)

type mockStore struct {
	err error
}

func (m mockStore) InsertObjectInList(objectID string, listID string) error {
	return m.err
}

type mockNotificationService struct {
	spyMsg string
	err    error
}

func (m *mockNotificationService) Send(userID string, msg string) error {
	m.spyMsg = msg
	return m.err
}

func TestService_AddToListAndNotifyUser(t *testing.T) {
	type args struct {
		userID   string
		listID   string
		itemInfo ItemInfo
	}
	tests := []struct {
		name                         string
		mockStoreError               error
		mockNotificationServiceError error
		args                         args
		wantMsg                      string
		wantErr                      bool
	}{
		{
			name: "success",
			args: args{
				userID: "user1",
				listID: "list1",
				itemInfo: ItemInfo{
					ID:          "item1",
					Name:        "Item 1",
					Description: "Description 1",
				},
			},
			wantMsg: "You have inserted object Item 1 in your list!",
			wantErr: false,
		},
		{
			name:           "store error",
			mockStoreError: fmt.Errorf("store error"),
			args: args{
				userID: "user1",
				listID: "list1",
				itemInfo: ItemInfo{
					ID:          "item1",
					Name:        "Item 1",
					Description: "Description 1",
				},
			},
			wantErr: true,
		},
		{
			name:                         "notification service error",
			mockNotificationServiceError: fmt.Errorf("notification service error"),
			args: args{
				userID: "user1",
				listID: "list1",
				itemInfo: ItemInfo{
					ID:          "item1",
					Name:        "Item 1",
					Description: "Description 1",
				},
			},
			wantMsg: "You have inserted object Item 1 in your list!",
			wantErr: true,
		},
		{
			name:                         "store and notification service error",
			mockStoreError:               fmt.Errorf("store error"),
			mockNotificationServiceError: fmt.Errorf("notification service error"),
			args: args{
				userID: "user1",
				listID: "list1",
				itemInfo: ItemInfo{
					ID:          "item1",
					Name:        "Item 1",
					Description: "Description 1",
				},
			},
			wantErr: true,
		},
		{
			name: "empty item info",
			args: args{
				userID:   "user1",
				listID:   "list1",
				itemInfo: ItemInfo{},
			},
			wantMsg: "You have inserted object  in your list!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mStore := mockStore{
				err: tt.mockStoreError,
			}
			mNotificationService := mockNotificationService{
				err: tt.mockNotificationServiceError,
			}

			s := Service{
				store:               &mStore,
				notificationService: &mNotificationService,
			}
			if err := s.AddToListAndNotifyUser(tt.args.userID, tt.args.listID, tt.args.itemInfo); (err != nil) != tt.wantErr {
				t.Errorf("AddToListAndNotifyUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if mNotificationService.spyMsg != tt.wantMsg {
				t.Errorf("AddToListAndNotifyUser() msg = %v, wantMsg %v", mNotificationService.spyMsg, tt.wantMsg)
			}
		})
	}
}
