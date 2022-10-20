package def

import (
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var (
	WashServerRSAKeyFilePath = pathGetenv("WASH_SERVER_RSA_KEYFILE_PATH", "~/wash_server_rsa_keyfile")

	TestTimeFactor = floatGetenv("GO_TEST_TIME_FACTOR", 1.0)
	TestSecond     = time.Duration(float64(time.Second) * TestTimeFactor)
)
