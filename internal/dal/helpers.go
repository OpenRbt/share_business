package dal

import (
	"fmt"

	"go.uber.org/zap"
)

func LogOptionalError(l *zap.SugaredLogger, module string, err error, additionalInfo ...any) {
	if err != nil {
		l.Named("dal").Named(module).Errorw(fmt.Sprintf("failed to perform database action, error : %s", err), additionalInfo...)
	}
}
