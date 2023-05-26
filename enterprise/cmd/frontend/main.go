// Command frontend is the enterprise frontend program.
package main

import (
	"os"

	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/shared"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/sanitycheck"
	"github.com/sourcegraph/sourcegraph/internal/service/svcmain"
	"github.com/sourcegraph/sourcegraph/ui/assets"

	_ "github.com/sourcegraph/sourcegraph/ui/assets/enterprise" // Select enterprise assets
)

func main() {
	sanitycheck.Pass()
	if os.Getenv("WEBPACK_DEV_SERVER") == "1" {
		assets.UseDevAssetsProvider()
	}
	svcmain.SingleServiceMainWithoutConf(shared.Service, svcmain.Config{}, svcmain.OutOfBandConfiguration{
		Logging: conf.NewStaticLogsSinksSource(log.SinksConfig{}),
	})
}
