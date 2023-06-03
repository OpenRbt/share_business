package app

import (
	"encoding/json"
	"errors"
	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"go.uber.org/zap"
	"time"
)

type RabbitClient interface {
	SendMessage(msg interface{}, service rabbit_vo.Service, routingKey rabbit_vo.RoutingKey, messageType rabbit_vo.MessageType) error
}

type Repo interface {
	SavePendingMessage(msg PendingMessage) error
	GetPendingMessages(lastMessageID int64) ([]PendingMessage, error)
	MarkMessageAsSent(messageID int64) error
}

type Worker struct {
	rabbitClient RabbitClient
	repo         Repo
	routingKey   string
	l            *zap.SugaredLogger
}

func NewWorker(l *zap.SugaredLogger, repo Repo, rabbitClient RabbitClient) *Worker {
	return &Worker{
		l:            l,
		rabbitClient: rabbitClient,
		repo:         repo,
	}
}

func (w *Worker) PrepareMessage(messageType rabbit_vo.MessageType, payload interface{}) error {
	if payload != nil {
		bytes, err := json.Marshal(payload)
		if err != nil {
			w.l.Errorw("failed to prepare message", "messageType", messageType)
			return ErrPrepareMessage
		}

		err = w.repo.SavePendingMessage(PendingMessage{
			MessageType: messageType,
			Payload:     bytes,
		})

		return err
	}
	return ErrPrepareMessage
}

func (w *Worker) ProcessMessages() {
	var lastMessageID int64

	for {
		lastMessageID = 0

		for {
			messages, err := w.repo.GetPendingMessages(lastMessageID)
			if err != nil && !errors.Is(err, ErrNotFound) {
				w.l.Errorw("failed to retrieve pending messages", "error", err)
			}

			if len(messages) == 0 {
				break
			}

			lastMessageID = int64(messages[len(messages)-1].ID)

			for _, message := range messages {
				err := w.rabbitClient.SendMessage(message.Payload, rabbit_vo.WashAdminService, rabbit_vo.WashAdminServesEventsRoutingKey, message.MessageType)
				if err != nil {
					w.l.Warnw("failed to send pending message", "messageID", message.ID, "error", err)
				} else {
					err = w.repo.MarkMessageAsSent(int64(message.ID))
					if err != nil {
						w.l.Warnw("failed to mark pending message as sent", "messageID", message.ID, "error", err)
					}
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
