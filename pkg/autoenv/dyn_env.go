package autoenv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// AutoEnv  struct defines a structure with two fields,
// LDelim and RDelim, which represent the left and right delimiters for placeholder variables in configuration files.
type AutoEnv struct{
	LDelim string
	RDelim string
	matcher *regexp.Regexp
}

// Creates a new instance of [autoenv.AutoEnv].
// This does also compile the regexp to be used later.
//
// Will panic if it cannot compile the regexp.
func New() *AutoEnv { 
	ld := regexp.QuoteMeta("{")
	rd := regexp.QuoteMeta("}")

	return &AutoEnv{
		LDelim: ld,
		RDelim: rd,
		matcher: regexp.MustCompile(fmt.Sprintf(`%s(\w+)%s`, ld, rd)),
	} 
}

// SetDelim sets the left and right delimiters for the AutoEnv struct.
//
// Example:
// # your env
// NAME="z3ntl3"
// HELLO="Hello {NAME}"
// 
// Here the '{' is the left delim and '}' is the right delim.	
//
// Parameters:
// - left_delim: the left delimiter to be set.
// - right_delim: the right delimiter to be set.
//
// Returns the instance for conveniently chaining methods 
// and will panic if it cannot compile the regexp.
func (env *AutoEnv) SetDelim(left_delim, right_delim string) *AutoEnv {
	env.LDelim  = regexp.QuoteMeta(left_delim)
	env.RDelim = regexp.QuoteMeta(right_delim)
	env.matcher = regexp.MustCompile(fmt.Sprintf(`%s(\w+)%s`, env.LDelim, env.RDelim))

	return env
}

// ResetDelims resets the left and right delimiters for the AutoEnv struct.
//
// Returns the instance for convenience it ill chain methods. Will panic if it cannot compile the regexp.
func (env *AutoEnv) ResetDelims() *AutoEnv {
	env.LDelim  = regexp.QuoteMeta("{")
	env.RDelim = regexp.QuoteMeta("}")
	env.matcher = regexp.MustCompile(fmt.Sprintf(`%s(\w+)%s`, env.LDelim, env.RDelim))

	return env
}

// Execute iterates over all the keys in the viper configuration and dynamically updates all
// placeholder variables
func (env *AutoEnv) Execute() {
	keys := viper.AllKeys()

	for _, key := range keys {
		value := viper.GetString(key)

		changed := env.interpolate(value)
		if value != changed {
			viper.Set(key, changed)
		}
	}
}

func (env *AutoEnv) interpolate(value string) string {
	delims := env.matcher.FindAllStringSubmatch(value, -1)
	for _, delim := range delims {
		placeholder := delim[0]
		variable := delim[1]

		var_value := viper.GetString(variable)
		if strings.ContainsAny(var_value, fmt.Sprintf("%s%s", env.LDelim, env.RDelim)){ 
			if strings.Contains(var_value, placeholder){
				var_value = strings.Split(var_value, placeholder)[0]
			} 

			var_value = env.interpolate(var_value)
		}

		value = strings.ReplaceAll(value, placeholder, var_value)
	}
	return value
}