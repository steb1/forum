package models

import (
	"database/sql"
	"fmt"
	"forum/lib"
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
	Read bool
}

func (nr *NotificationRepository) GetNotificationByID(NotificationID string) (*Notification, error) {
	var notification Notification
	row := nr.db.QueryRow("SELECT id, authorID, postID, postOwnerID, notif_type, time, readed FROM notification WHERE id = ?", NotificationID)
	err := row.Scan(&notification.ID, &notification.AuthorID, &notification.PostID, &notification.PostOwnerID, &notification.Notif_type, &notification.Time, &notification.Read)
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
	rows, err := nr.db.Query("SELECT id, authorID, postID, postOwnerID, notif_type, time, readed FROM notification")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, err
	}
	for rows.Next() {
		var notif Notification
		err := rows.Scan(&notif.ID, &notif.AuthorID, &notif.PostID, &notif.PostOwnerID, &notif.Notif_type, &notif.Time, &notif.Read)
		if err != nil {
			fmt.Println("2", err)
			return nil, err
		}
		notifications = append(notifications, &notif)
	}
	return notifications, nil
}

func (nr *NotificationRepository) GetAllNotifsByUser(userID string) ([]Notification, error) {
	var notifications []Notification
	rows, err := nr.db.Query("SELECT n.id, n.authorID, n.postID, n.postOwnerID, n.notif_type, n.time, n.readed FROM notification n JOIN user u ON n.authorid=u.id JOIN post p ON n.postid=p.id WHERE n.postownerid = ?", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("1", err)
		return nil, err
	}
	for rows.Next() {
		var notif Notification
		err := rows.Scan(&notif.ID, &notif.AuthorID, &notif.PostID, &notif.PostOwnerID, &notif.Notif_type, &notif.Time, &notif.Read)
		if err != nil {
			return nil, err
		}
		
		notifications = append([]Notification{notif}, notifications...)

	}
	return notifications, nil
}

func (nr *NotificationRepository) CreateNotification(notification *Notification) error {

	_, err := nr.db.Exec("INSERT INTO notification (id, authorID, postID, postOwnerID, notif_type, time, readed) VALUES (? ,? , ?, ?, ?, ?, ?)",
		notification.ID, notification.AuthorID, notification.PostID, notification.PostOwnerID, notification.Notif_type, notification.Time, notification.Read)
	return err
}
func FormatNotifications(Notifications []Notification) []string {
	var FormatedNotif []string
	for _, notification := range Notifications {
		var motif string
		if notification.Notif_type == "like" {
			motif = "have liked your post"
		}
		if notification.Notif_type == "Comment_like" {
			motif = "have liked your comment"
		}
		if notification.Notif_type == "dislike" {
			motif = "have disliked your post"
		}
		if notification.Notif_type == "Comment_dislike" {
			motif = "have disliked your comment"
		}
		if notification.Notif_type == "Comment" {
			motif = "have commented your post"
		}
		author, _ := UserRepo.GetUserByID(notification.AuthorID)
		
		timeago := lib.FormatDateDB(notification.Time)
		notif := fmt.Sprintf("%s %s   %s", author.Username, motif, timeago)
		FormatedNotif = append(FormatedNotif, notif)
	}
	return FormatedNotif
}
func ListNotifications(notifications []string) (s string) {
	for i := len(notifications) - 1; i >= 0; i-- {
		s += notifications[i]
		s += "\\n"
	}
	return s
}
