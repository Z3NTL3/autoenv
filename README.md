<!-- header -->

<div align="center"> 
    <img src="./img/autoenv_nobg.png" width="280" style="border-radius: 4px;"><br>  
</div>

# AutoEnv

AutoEnv is a lightweight Viper plugin that handles variable interpolation and substitution, while Viper manages the configuration files.

Implementation is very minimal, Viper is the only dependency in this package. It leverages [os.Expand](https://pkg.go.dev/os#Expand) together with Viper managed environment and configuration variables, enabling data format–agnostic variable interpolation

Some known supported formats are TOML, JSON and YAML. Check further with [Viper](https://github.com/spf13/viper) for more details.

#### Features

- Ease of use
- Deep interpolation/substitution of variables
- Plugin for [Viper](https://github.com/spf13/viper)
- Elegant
- Minimal implementation

### Var interpolation/substitution forms

Either $var or ${var} where var is case sensitive, therefore ${var} or $var and ${VAR} or $VAR are different.

More documentation [here](https://pkg.go.dev/os#Expand)

#### Installation

> `go get z3ntl3.com/autoenv/v2`

#### After installation

You can just import it in your project and use it.

```go
package main

import "z3ntl3.com/autoenv/v2"
```

### Example

Let's consider the following as a `.env` file:

```env
NAME="z3ntl3"
DOMAIN="com"

WEBSITE="${NAME}.${DOMAIN}"
```

And this code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/v2/pkg/autoenv"
)

func main(){
    viper.AddConfigPath("../tests")

	// viper.AutomaticEnv()
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.Provision() // interpolates all the placeholder variables
	fmt.Println(viper.GetString("WEBSITE")) // see the change
}
```

### Output

`WEBSITE="${NAME}.${DOMAIN}"` turns into `z3ntl3.com`.

#### License

[GNU GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)
