package main

import (
	"errors"
	"fmt"

	xContext "golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func GroupTask(fList []func() error) (taskErr error) {
	// 传入闭包列表
	ctx, cancel := xContext.WithCancel(xContext.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for _, f := range fList {
		task := f
		// 新建子协程
		group.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic recovered: %v", r)
				}
				cancel()
			}()
			//检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
			err = CheckGoroutineErr(errCtx)
			if err != nil {
				return err

			}
			err = task()
			if err != nil {
				return err
			}
			return nil
		})
	}

	// 捕获err
	err := group.Wait()
	if err == nil {
		return
	} else {
		taskErr = errors.New(err.Error())
	}
	return
}

//校验是否有协程已发生错误
func CheckGoroutineErr(errContext xContext.Context) error {
	select {
	case <-errContext.Done():
		return errContext.Err()
	default:
		return nil
	}
}
