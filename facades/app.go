package facades

import (
	foundationcontract "github.com/wesleysnt/framework/contracts/foundation"
	"github.com/wesleysnt/framework/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
