package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"washbonus/internal/app"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/swag"
	"go.uber.org/zap"
)

type errorSetter interface {
	SetPayload(payload *models.Error)
	SetStatusCode(code int)
}

func errorStatusCode(err error) int {
	switch err {
	case entities.ErrBadRequest:
		return http.StatusBadRequest
	case entities.ErrForbidden:
		return http.StatusForbidden
	case entities.ErrNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

var errorMapping = map[error]int{
	entities.ErrBadRequest: http.StatusBadRequest,
	entities.ErrForbidden:  http.StatusForbidden,
	entities.ErrNotFound:   http.StatusNotFound,
}

func setAPIError(l *zap.SugaredLogger, op string, err error, responder interface{}) {
	r, ok := responder.(errorSetter)
	if !ok {
		return
	}

	statusCode, exists := getStatusCodeForError(err)

	msg := err.Error()
	if !exists {
		statusCode = http.StatusInternalServerError
		msg = "internal server error"

		l.Errorln(op, err)
	}

	r.SetPayload(&models.Error{Code: swag.Int32(int32(statusCode)), Message: swag.String(msg)})
	r.SetStatusCode(statusCode)
}

func getStatusCodeForError(err error) (int, bool) {
	for knownErr, code := range errorMapping {
		if errors.Is(err, knownErr) {
			return code, true
		}
	}
	code, exists := errorMapping[err]
	return code, exists
}

func splitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}

func createCtxWithUserID(req *http.Request, auth *app.AdminAuth) context.Context {
	return context.WithValue(req.Context(), "userID", auth.User.ID)
}
