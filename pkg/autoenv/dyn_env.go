package autoenv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var vMatcher = regexp.MustCompile(`\{([^}]+)\}`) // Matches placeholder variables like {NAME}
type AutoEnv struct{}

// Creates a new instance of [autoenv.AutoEnv].
func New() *AutoEnv { return &AutoEnv{} }

// Execute iterates over all the keys in the viper configuration and dynamically updates all
// placeholder variables
func (env *AutoEnv) Execute() {
	keys := viper.AllKeys()

	for _, key := range keys {
		value := viper.GetString(key)

		changed := interpolate(value)
		if value != changed {
			viper.Set(key, changed)
		}
	}
}

func interpolate(value string) string {
	delims := vMatcher.FindAllStringSubmatch(value, -1)

	for _, delim := range delims {
		placeholder := delim[0]
		variable := delim[1]

		var_value := viper.GetString(variable)
		if strings.ContainsAny(var_value, "{}") {

			if strings.Contains(var_value, fmt.Sprintf("{%s}", variable)){
				var_value = strings.Split(var_value, placeholder)[0]
			} 

			var_value = interpolate(var_value)
		}

		value = strings.ReplaceAll(value, placeholder, var_value)
	}
	return value
}