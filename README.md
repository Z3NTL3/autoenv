<!-- header -->

<div align="center"> 
    <img src="./img/autoenv_nobg.png" width="280" style="border-radius: 4px;"><br>  
</div>


# AutoEnv

Dynamically interpolate placeholder variables in any config file. Made especially to be a plugin for [Viper](https://github.com/spf13/viper).

#### Features
- Very easy to use
- Deep interpolation/substitution of variables
- Works with [Viper](https://github.com/spf13/viper)
- Elegant
- Ability to change the left/right delimiters for convenience [How?](#how-to-change-delimiters)

#### Installation
> ``go get z3ntl3.com/autoenv``

#### After that

You can just import it in your project and use it.
```go
package main

import "z3ntl3.com/autoenv"
```

### Example

Let's consider the following as a `.env` file:

```env
NAME="z3ntl3"
DOMAIN="com"

WEBSITE="{NAME}.{DOMAIN}"
```

And this code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/pkg/autoenv"
)

func main(){
    viper.AddConfigPath("../tests")

	// viper.AutomaticEnv()
	viper.SetConfigName("test")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.New().Execute() // interpolates all the placeholder variables
	fmt.Println(viper.GetString("WEBSITE")) // see the change
}
```
### Output
``WEBSITE="{NAME}.{DOMAIN}"`` turns into ``z3ntl3.com``.


### How to change delimiters
Have this example ENV

```env
NAME="z3ntl3"
DOMAIN="com"

WEBSITE="<NAME>.<DOMAIN>"
```

And this code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"z3ntl3.com/autoenv/pkg/autoenv"
)

func main() {
	viper.AddConfigPath("../../tests")

	// viper.AutomaticEnv()
	viper.SetConfigName("test2")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	autoenv.New().SetDelim("<", ">").Execute() // interpolates all the placeholder variables with the delims
	fmt.Println(viper.GetString("WEBSITE")) // see the change
}
```

### Output
``WEBSITE="<NAME>.<DOMAIN>"`` turns into ``z3ntl3.com``.

#### License
[GNU GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)