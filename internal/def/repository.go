package def

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DBHost         = os.Getenv("MSRV_DB_HOST")
	DBPort         = intGetenv("MSRV_DB_PORT", 5432)
	DBUser         = os.Getenv("MSRV_DB_USER")
	DBPass         = os.Getenv("MSRV_DB_PASS")
	DBName         = os.Getenv("MSRV_DB_NAME")
	DBSSLModeIsReq = boolGetenv("MSRV_DB_SSL_MODE_IS_REQUIRE")
	GooseDir       = "./migration"
	ResetDB        = boolGetenv("MSRV_RESET_DB")
)
