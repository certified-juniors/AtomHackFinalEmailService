package model

import (
	"time"
)

// Document представляет модель документа
type Document struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Title          string         `json:"title"`
	Owner          string         `json:"owner"`
	ReceivedTime   *time.Time     `json:"receivedTime"`
	SentTime       *time.Time     `json:"sentTime"`
	CreatedAt      time.Time      `json:"createdAt"`
	DeliveryStatus *DeliveryStatus `json:"deliveryStatus"`
	Status         Status         `json:"status"`
	Payload        string         `json:"payload" gorm:"type:text"`
	Files          []File         `json:"files"`
}

// func (d *Document) FormatMessage() string {
// 	message := "ID: " + strconv.Itoa(int(d.ID)) + "\n" +
// 		"Title: " + d.Title + "\n" +
// 		"Owner: " + d.Owner + "\n" +
// 		"Created At: " + d.CreatedAt.Format(time.RFC3339) + "\n" +
// 		"Payload: " + d.Payload

// 	if d.ReceivedTime != nil {
// 		message += "\nReceived Time: " + d.ReceivedTime.Format(time.RFC3339)
// 	}
// 	if d.SentTime != nil {
// 		message += "\nSent Time: " + d.SentTime.Format(time.RFC3339)
// 	}
// 	if d.DeliveryStatus != nil {
// 		message += "\nDelivery Status: " + *d.DeliveryStatus
// 	}

// 	return message
// }