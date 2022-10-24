package def

import (
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var (
	WashServerRSAPrivateKeyFilePath = pathGetenv("WASH_SERVER_RSA_PRIVATE_KEYFILE_PATH", "~/certs/wash-server-rsa-keyfile")
	WashServerRSAPublicKeyFilePath  = pathGetenv("WASH_SERVER_RSA_PUBLIC_KEYFILE_PATH", "~/certs/wash-server-rsa-keyfile.pub")

	TestTimeFactor = floatGetenv("GO_TEST_TIME_FACTOR", 1.0)
	TestSecond     = time.Duration(float64(time.Second) * TestTimeFactor)
)
