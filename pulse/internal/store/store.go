package store

import (
	"sync"

	"pulse/internal/models"
)

type (
	UserNotifications map[string][]models.Notification

	NotificationStore struct {
		Data UserNotifications
		mu   sync.RWMutex
	}
)

func (ns *NotificationStore) Add(userID string,
	notification models.Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Data[userID] = append(ns.Data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []models.Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.Data[userID]
}
