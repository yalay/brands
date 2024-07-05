## Mobile Phone Model to Brand Conversion
### Supported mobile brand
 - Xiaomi
 - Redmi
 - SHARK
 - coolpad
 - google
 - vivo
 - Samsung
 - realme
 - OPPO
 - OnePlus
 - MEIZU
 - HUAWEI
 - HONOR
 - Apple
 - 360
 - nubia
 - nokia
 - nothing
 - Lenovo
 - ZTE
 - ASUS
 - Letv
 - Motorola
 - iQOO
 - Smartisan
 - SONY

### A basic example
```
package main

import (
	"fmt"

	"github.com/yalay/brands"
)

func main() {
	fmt.Println(brands.Model2Brand("21121119SC")) // Xiaomi
	fmt.Println(brands.Model2Brand("ADT-AN00"))   // HUAWEI
	fmt.Println(brands.Model2Brand("SM-S7110"))   // Samsung
	fmt.Println(brands.Model2Brand("V1824BA"))    // vivo
}
```
