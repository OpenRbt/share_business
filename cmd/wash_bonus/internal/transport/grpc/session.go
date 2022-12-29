package grpc

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/yeqown/go-qrcode/v2"
	qr_writers "github.com/yeqown/go-qrcode/writer/standard"
	"io"
	"net/http"
	"wash_bonus/intapi"
)

func (s *Service) Begin(ctx context.Context, request *intapi.BeginRequest) (response *intapi.BeginAnswer, err error) {
	var code int64
	var msg string

	defer func() {
		if err != nil {
			response = &intapi.BeginAnswer{
				Error: &intapi.ErrorMsg{
					Code: code,
					Msg:  msg,
				},
			}
			err = nil
		}
	}()

	if !s.isValidConnection(request.ServiceKey, request.ConnectionID) {
		code = http.StatusForbidden
		msg = "bad connection"
		err = fmt.Errorf(msg)
	}

	connId, err := uuid.FromString(request.ConnectionID)
	if err != nil {
		code = http.StatusBadRequest
		msg = "bad request"
		return
	}

	session, err := s.sessionsSvc.CreateSession(ctx, connId, request.PostID, request.ServiceKey)
	if err != nil {
		code = http.StatusInternalServerError
		msg = "failed to create session"
		return
	}

	url := s.createLink(session.ID)
	qr, err := s.createQR(url)
	if err != nil {
		code = http.StatusInternalServerError
		msg = "failed to generate qr"
		return
	}

	response = &intapi.BeginAnswer{
		SessionID: session.ID.String(),
		SessionQR: qr,
		Error:     nil,
	}

	return
}

func (s *Service) Refresh(ctx context.Context, request *intapi.RefreshRequest) (response *intapi.RefreshAnswer, err error) {
	var (
		code         int64
		msg          string
		responseData map[string]*intapi.SessionRefreshResponseData
	)

	defer func() {
		if err != nil {
			response = &intapi.RefreshAnswer{
				Error: &intapi.ErrorMsg{
					Code: code,
					Msg:  msg,
				},
			}
			err = nil
		}
	}()

	if !s.isValidConnection(request.ServiceKey, request.ConnectionID) {
		code = http.StatusForbidden
		msg = "bad connection"
		err = fmt.Errorf(msg)
	}

	for sessionID, data := range request.Sessions {
		id, err := uuid.FromString(sessionID)
		if err != nil {
			code = http.StatusBadRequest
			msg = "bad session id"
			err = fmt.Errorf(msg)
			return nil, err
		}

		balance := decimal.NewFromInt(data.PostBalance)

		session, err := s.sessionsSvc.RefreshSession(ctx, id, balance)
		if err != nil {
			return nil, err
		}

		data := &intapi.SessionRefreshResponseData{}

		if session.User != nil {
			data.UserAssigned = true
		}

		if session.AddAmount != decimal.Zero {
			amount := session.AddAmount.IntPart()
			data.AddAmount = amount

			err = s.sessionsSvc.ConsumeMoney(ctx, session.ID)
			if err != nil {
				return nil, err
			}
		}

		responseData[sessionID] = data
	}

	response = &intapi.RefreshAnswer{
		Data:  responseData,
		Error: nil,
	}

	return
}

func (s *Service) End(ctx context.Context, request *intapi.FinishRequest) (response *intapi.FinishAnswer, err error) {
	var code int64
	var msg string

	defer func() {
		if err != nil {
			response = &intapi.FinishAnswer{
				Error: &intapi.ErrorMsg{
					Code: code,
					Msg:  msg,
				},
			}
			err = nil
		}
	}()

	if !s.isValidConnection(request.ServiceKey, request.ConnectionID) {
		code = http.StatusForbidden
		msg = "bad connection"
		err = fmt.Errorf(msg)
	}

	id, err := uuid.FromString(request.SessionID)
	if err != nil {
		code = http.StatusBadRequest
		msg = "bad session id"
		err = fmt.Errorf(msg)
	}

	err = s.sessionsSvc.EndSession(ctx, id)

	if err != nil {
		return
	}

	response = &intapi.FinishAnswer{
		Error: nil,
	}

	return
}

func (s *Service) createLink(sessionID uuid.UUID) (url string) {
	url = fmt.Sprintf("%s/session/%s", s.basePath, sessionID)

	return
}

func (s *Service) createQR(url string) (qr []byte, err error) {
	qrc, err := qrcode.NewWith(url, []qrcode.EncodeOption{}...)
	if err != nil {
		return
	}

	var (
		w       io.WriteCloser
		options []qr_writers.ImageOption = []qr_writers.ImageOption{
			qr_writers.WithFgColorRGBHex("#ffffff"),
			qr_writers.WithBgColorRGBHex("#000000"),
		}
	)

	writer := qr_writers.NewWithWriter(w, options...)

	err = qrc.Save(writer)
	if err != nil {
		return
	}

	_, err = w.Write(qr)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) EnterMoney(ctx context.Context, request *intapi.EnterMoneyRequest) (response *intapi.EnterMoneyAnswer, err error) {
	return
}
