package commons

import "fmt"

type Async struct {
	Result interface{}
	Error  error
}

func DeferErrorAsync(asyncChan chan Async) {
	defer close(asyncChan)
	if x := recover(); x != nil {
		fmt.Println(x.(error).Error())
		asyncChan <- Async{
			Error: x.(error),
		}
	}
}

func GetAsyncResult(asyncChan chan Async) (interface{}) {
	async := <-asyncChan
	if async.Error != nil {
		fmt.Println(async.Error.Error())
		panic(async.Error)
	} else {
		return async.Result
	}
}

func ClearGoRoutine(asyncChan chan Async) () {
	defer func() {
		if x := recover(); x != nil {
		}
	}()
	<-asyncChan
}
