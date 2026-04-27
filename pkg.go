/*
AutoEnv is a lightweight Viper plugin that handles variable interpolation and substitution, while Viper manages the configuration files.

Therefore AutoEnv is data format agnostic, known supported formats are TOML, JSON and YAML
*/
package autoenv

import (
	"os"

	"github.com/spf13/viper"
)

// Provision environment and configuration files load by Viper, in sense of performing variable expansion
// This is a custom implementation around [os.Expand] with Viper mappings via [github.com/spf13/viper.Set], [github.com/spf13/viper.AllKeys] to perform interpolation and variable substitution.
//
// Variable mappings should be used in the form of
//   - $var
//   - or ${VAR}
//
// where var is case-sensitive, therefore $var or ${var} and $VAR or ${VAR} are different.
func Provision() {
	keys := viper.AllKeys()

	for _, key := range keys {
		var_ := viper.GetString(key)
		// interpolation magic happens down here
		expanded := os.Expand(var_, func(v string) string {
			return viper.GetString(v)
		})
		viper.Set(key, expanded)
	}
}
