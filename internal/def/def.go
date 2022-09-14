// Code generated by mtgroup-generator.
package def

import (
	"os"
	"strconv"
	"strings"
	"time"

	"wash-bonus/internal/api/restapi/restapi"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

// Log field names.
const (
	LogHost       = "host"
	LogPort       = "port"
	LogAddr       = "addr"
	LogRemote     = "remote" // aligned IPv4:Port "   192.168.0.42:1234 "
	LogFunc       = "func"   // RPC method name, REST resource path
	LogHTTPMethod = "httpMethod"
	LogHTTPStatus = "httpStatus"
	LogGRPCCode   = "grpcCode"
	LogUser       = "userID"
	LogEvent      = "ev"
	LogEventID    = "evID"

	ExtauthEndpointEnvName = "MSRV_EXTAUTH_ENDPOINT"
)

var (
	log = structlog.New()

	oapiHost, oapiPort, oapiBasePath = swaggerEndpoint()

	TestTimeFactor = floatGetenv("GO_TEST_TIME_FACTOR", 1.0)
	TestSecond     = time.Duration(float64(time.Second) * TestTimeFactor)

	DBHost         = os.Getenv("MSRV_DB_HOST")
	DBPort         = intGetenv("MSRV_DB_PORT", 5432)
	DBUser         = os.Getenv("MSRV_DB_USER")
	DBPass         = os.Getenv("MSRV_DB_PASS")
	DBName         = os.Getenv("MSRV_DB_NAME")
	DBSSLModeIsReq = boolGetenv("MSRV_DB_SSL_MODE_IS_REQUIRE")
	GooseDir       = "./migration"
	ResetDB        = boolGetenv("MSRV_RESET_DB")

	APIHost             = strGetenv("MSRV_HOST", oapiHost)
	APIPort             = intGetenv("MSRV_PORT", oapiPort)
	APIBasePath         = strGetenv("MSRV_BASEPATH", oapiBasePath)
	ExtauthEndpoint     = os.Getenv(ExtauthEndpointEnvName)
	CORSAllowedOrigins  = os.Getenv("MSRV_CORS_ALLOWED_ORIGINS")
	DisableCookieSecure = boolGetenv("MSRV_DISABLE_COOKIE_SECURE")

	FirebaseKeyFilePath = strGetenv("FIREBASE_KEYFILE_PATH", "../../firebase.json")
	FirebaseApiKey      = os.Getenv("FIREBASE_API_KEY")
)

func intGetenv(name string, def int) int {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Err(errors.Errorf("failed to parse %q=%q as int: %v", name, value, err))
		return def
	}
	return i
}

func floatGetenv(name string, def float64) float64 {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Err(errors.Errorf("failed to parse %q=%q as float: %v", name, value, err))
		return def
	}
	return v
}

func strGetenv(name, def string) string {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	return value
}

func boolGetenv(name string) bool {
	value := os.Getenv(name)
	if value == "true" {
		return true
	}
	return false
}

func swaggerEndpoint() (host string, port int, basePath string) {
	const portHTTP = 80
	const portHTTPS = 443

	spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return "", 0, ""
	}

	host, port, err = swag.SplitHostPort(spec.Host())
	switch {
	case err == nil:
		return host, port, spec.BasePath()
	case strings.Contains(err.Error(), "missing port"):
		schemes := spec.Spec().Schemes
		switch {
		case len(schemes) == 1 && schemes[0] == "http":
			return spec.Host(), portHTTP, spec.BasePath()
		case len(schemes) == 1 && schemes[0] == "https":
			return spec.Host(), portHTTPS, spec.BasePath()
		}
	}
	return spec.Host(), 0, spec.BasePath()
}
