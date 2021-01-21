package ftc

import (
	"context"
	"fmt"
	"github.com/vardius/gollback"
)

// 测试 gollback
func gollbackTest() {
	a,v := gollback.All(context.TODO(),
		func(ctx context.Context) (interface{}, error) {
			return 1,nil
		},
		func(ctx context.Context) (interface{}, error) {
			return 2,nil
		},
		func(ctx context.Context) (interface{}, error) {
			return 3,nil
		},
	)
	fmt.Println(a,v)
}
