# 2.0 account + new account

[에러 핸들링](#에러-핸들링)


## struct 생성

```go
// Account struct  // 코멘트가 반드시 필요
package banking

type Account struct {
	Owner string // 이것도 첫글자가 소문자면 private
	Balance int
}
```


```go
package main

import (
	"fmt"

	"github.com/lihakim/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "bilbo", Balance: 1000}
	fmt.Println(account) // {bilbo 1000}
}
```

**문제점:** 

- 다른 사람들이 account를 마음대로 조작할 수 있음 (오너 이름이나 balance를 마음대로 바꿀 수 있음)
- Java 같은 언어에서는 constructor로 조작을 불가능하게 할 수 있지만 go 에는 constructor 가 없음 
  - 함수로 construct 하거나 struct를 만들도록 함


```go
package accounts

// Account struct
type Account struct { // struct를 private으로 만들면 struct 자체에는 접근하지 못하게 함
	owner string // 이것도 첫글자가 소문자면 private
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account{ // Account를 copy
	account := Account {
		owner: owner,
		balance: 0,
	}
	return &account // 새로운 object을 생성하지 않고 복사본 자체를 리턴
}
```

```go
package main

import (
	"fmt"

	"github.com/lihakim/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bilbo")
	fmt.Println(account) // &{bilbo 0}
	// &은 카피가 아니라 새로운 object임
	// copy를 매번 하는 것이 아님에 주의
}
```

`&`은 카피가 아니라 새로운 object임(copy를 매번 하는 것이 아님에 주의)





## 에러 핸들링

### go.mod file을 찾을 수 없다고 하는 경우

```
go run main.go
main.go:6:2: no required module provides package github.com/username/learngo/banking: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

해결방법

```
go mod init github.com/username/learngo
```