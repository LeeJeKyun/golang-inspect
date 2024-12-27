package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	//WaitGroup 할당
	var wg sync.WaitGroup

	//Account 구조체 생성(계좌)
	//account := &Account{0}
	var account Account

	//WaitGroup에 실행횟수 설정
	wg.Add(10)

	//고루틴을 반복하여 실행(1000번)
	for i := 0; i < 1000; i++ {
		fmt.Println("new goroutine:", i)

		//
		go func() {
			for {
				DepositAndWithdraw(&account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
