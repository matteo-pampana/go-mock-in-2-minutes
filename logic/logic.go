package logic

import "fmt"

type store interface {
	InsertObjectInList(objectID string, listID string) error
}

type notificationService interface {
	Send(userID string, msg string) error
}

type ItemInfo struct {
	ID          string
	Name        string
	Description string
}

type Service struct {
	store               store
	notificationService notificationService
}

func (s Service) AddToListAndNotifyUser(userID string, listID string, itemInfo ItemInfo) error {
	err := s.store.InsertObjectInList(itemInfo.ID, listID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("You have inserted object %s in your list!", itemInfo.Name)

	return s.notificationService.Send(userID, msg)
}
