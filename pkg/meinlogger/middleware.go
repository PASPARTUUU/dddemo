package meinlogger

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LogrusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		req := ectx.Request()
		// resp := c.Response()

		logrus.WithFields(map[string]interface{}{
			"uri":    req.RequestURI,
			"method": req.Method,
			// "status": resp.Status,
		}).Info("Handled request")

		if err := next(ectx); err != nil {
			ectx.Error(err)
		}

		return nil
	}
}
