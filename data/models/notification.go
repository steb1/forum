package models

import (
	"database/sql"
	"fmt"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

type Notification struct {
	ID          string
	AuthorID    string
	PostID      string
	PostOwnerID string
	Notif_type  string
	Time        string
}

func (nr *NotificationRepository) GetNotificationByID(NotificationID string) (*Notification, error) {
	var notification Notification
	row := nr.db.QueryRow("SELECT id, authorID, postID, postOwnerID, notif_type FROM notification WHERE id = ?", NotificationID)
	err := row.Scan(&notification.ID, &notification.AuthorID, &notification.PostID, &notification.PostOwnerID, &notification.Notif_type, &notification.Time)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}
	return &notification, nil
}

func (nr *NotificationRepository) GetAllNotifs() ([]*Notification, error) {
	var notifications []*Notification
	rows, err := nr.db.Query("SELECT id, authorID, postID, postOwnerID, notif_type,time FROM notification")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		fmt.Println("1", err)
		return nil, err
	}
	for rows.Next() {
		var notif Notification
		err := rows.Scan(&notif.ID, &notif.AuthorID, &notif.PostID, &notif.PostOwnerID, &notif.Notif_type, &notif.Time)
		if err != nil {
			fmt.Println("2", err)
			return nil, err
		}
		notifications = append(notifications, &notif)
	}
	return notifications, nil
}
func (nr *NotificationRepository) CreateNotification(notification *Notification) error {

	_, err := nr.db.Exec("INSERT INTO notification (id, authorID, postID, postOwnerID, notif_type, time) VALUES (?, ?, ?, ?, ?, ?)",
		notification.ID, notification.AuthorID, notification.PostID, notification.PostOwnerID, notification.Notif_type, notification.Time)
	return err
}
