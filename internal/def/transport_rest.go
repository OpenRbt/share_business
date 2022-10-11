package def

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	oapiHost, oapiPort, oapiBasePath = swaggerEndpoint()

	APIHost     = strGetenv("MSRV_HOST", oapiHost)
	APIPort     = intGetenv("MSRV_PORT", oapiPort)
	APIBasePath = strGetenv("MSRV_BASEPATH", oapiBasePath)

	FirebaseKeyFilePath = pathGetenv("FIREBASE_KEYFILE_PATH", "~/firebase_keyfile.json")
	DisableCookieSecure = boolGetenv("MSRV_DISABLE_COOKIE_SECURE")

	CORSAllowedOrigins = os.Getenv("MSRV_CORS_ALLOWED_ORIGINS")
)
