package autoenv_test

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/pkg/autoenv"
)

func Test_LoadEnv(t *testing.T) {
	viper.AddConfigPath(".")

	// viper.AutomaticEnv()
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.New().Execute()
	t.Logf("Evaluation of WEBSITE: %s", viper.GetString("website"))
}
/*
INPUT:

```env
NAME="z3ntl3"
DOMAIN="com"

WEBSITE="{NAME}.{DOMAIN}"
```

OUTPUT:


dyn_env_test.go:23: Evaluation of env variable 'WEBSITE': z3ntl3.com
*/