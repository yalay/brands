## Mobile Phone Model to Brand Conversion
## Supported mobile brand
 - Xiaomi
 - vivo
 - Samsung
 - Realme
 - oppo
 - OnePlus
 - MEIZU
 - HUAWEI
 - Apple
 - 360

#### A basic example:
```
package main

import (
	"fmt"

	"github.com/yalay/mm"
)

func main() {
	fmt.Println(mm.Model2Brand("21121119SC")) // Xiaomi
	fmt.Println(mm.Model2Brand("ADT-AN00"))   // HUAWEI
	fmt.Println(mm.Model2Brand("SM-S7110"))   // Samsung
	fmt.Println(mm.Model2Brand("V1824BA"))    // vivo
}

```