package main

import (
	"math/rand"
	"os"
	"time"

	"k8s.io/apiserver/pkg/util/logs"

	"github.com/glennswest/winsdncontroller/pkg/cmd/flagtypes"
	"github.com/glennswest/winsdncontroller/pkg/cmd/openshift-sdn"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	rand.Seed(time.Now().UTC().UnixNano())

	cmd := openshift_sdn.NewOpenShiftSDNCommand("openshift-sdn", os.Stderr)
	flagtypes.GLog(cmd.PersistentFlags())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
