<!-- header -->

<div align="center"> 
    <img src="image.png" width="300" style="border-radius: 4px;"><br>  
</div>


# AutoEnv

Dynamically interpolate variables in any config file. Made especially to be a plugin for [Viper](https://github.com/spf13/viper)

##### Installation
> ``go get z3ntl3.com/autoenv``

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

	autoenv.New().Execute()
}
```


### Output
``WEBSITE="{NAME}.{DOMAIN}"`` turns into ``z3ntl3.com``.

#### License
[GNU GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)