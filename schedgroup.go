package ftc

import (
	"context"
	"fmt"
	sgp "github.com/mdlayher/schedgroup"
	"time"
)
// 此包，内部使用了 heap将时间排列起来，因为我们使用time.Timer也行，但是它不能排顺序，这个包自动排序了。更节省时间。
func schedgroup_test(){
	g := sgp.New(context.TODO())
	for i:= 0;i < 100;i++ {
		n := i +1
		g.Delay(time.Duration(n)*100*time.Millisecond, func() {
			fmt.Println(n)
		})
	}
	g.Wait()
}

