## 챕터별

| chapter |          title          |                 Link                  |
| :-----: | :---------------------: | :-----------------------------------: |
|   1.0   |      main package       |      [:link:](#10-main-package)       |
|   1.1   |  packages and imports   |  [:link:](#11-packages-and-imports)   |
|   1.2   | variables and constants | [:link:](#12-variables-and-constants) |
|   1.4   |        functions        |        [:link:](#14-functions)        |
|   1.5   |   for, range, ...args   |     [:link:](#15-for-range-args)      |
|   1.6   |     if with a twist     |     [:link:](#16-if-with-a-twist)     |
|   1.7   |         switch          |         [:link:](#17-switch)          |
|   1.8   |        pointers         |        [:link:](#18-pointers)         |
|   1.9   |    arrays and slices    |    [:link:](#19-arrays-and-slices)    |
|  1.10   |          maps           |          [:link:](#110-maps)          |
|  1.11   |         structs         |        [:link:](#111-structs)         |

---

# 1.0 main package

- `main.go` 파일명은 프로젝트를 **컴파일** 하려면 반드시 가지고 있어야 함. 그게 아니라면 바꿔도 상관 x
- 왜냐하면 `main.go` 가 entry point 임 -> 컴파일러가 가장 먼저 찾는 패키지 이름

# 1.1 Packages and Imports

- fmt: formatting 을 위한 패키지

- 패키지가 어떻게 동작하나
- 왜 패키지 함수는 대문자로 시작하나
  - export 하고 싶으면 함수를 대문자로 시작해주면 됨

```go
package something

func bye() { // not able to import from other files
	fmt.Println("bye")
}

func Hello() { // able to import automatically - Start with Capital!
	fmt.Println("hello")
}
```

```go
package main

import (
	"fmt"
	"github.com/lihakim/learngo/something" // 임포트 되었음
)

func main() {
	fmt.Println("hello world")
	something.Hello()
}
```

# 1.2 Variables and Constants

```go
package main

import "fmt"

func main() {
	name := "nico" // 이렇게 해도 되고
    var name string = "nico" // 이렇게 해도 됨

	fmt.Println(name)
}
```

```go
package main
import "fmt"

func multiply(a int,b int) int { // 1. 하나하나 다 타입 지정
	return a * b
}

func multiply2(a, b int) int {  // 2. 파라미터의 타입이 같으면 하나로 쓸 수 있음
	return a * b
}

func main() {
 	fmt.Println(multiply(3, 4))
}
```

## multiple value를 반환하는 함수

```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func main() {
 	totalLength, upperName := lenAndUpper("nico")
	 fmt.Println(totalLength, upperName)

    totalLength2, _ := lenAndUpper("bilbo") // 무시하고자 하는 값은 _ 로
     fmt.Println(totalLength2)
}
```

## 매개변수들이 모두 같은 타입일 때 한 번에 타입 지정하기

```go
func repeatMe(words ...string) {
	fmt.Println(words);
}
func main() {
 	repeatMe("bilbo", "nenya", "max") // [bilbo nenya max]
}
```

# 1.4 Functions

## naked return

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
    // 리턴값이 비어있어도 위에서 선언해줬기 때문에 에러없이 리턴값이 반환됨
}

func main() {
 totalLength, upperCase := lenAndUpper("bilbo")
 fmt.Println(totalLength, upperCase)
}
```

## defer

함수가 끝났을 때 추가적으로 무언가 동작하도록 할 수 있음

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!") // #3 I'm done 출력
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // #2 값 리턴
}

func main() {
 totalLength, upperCase := lenAndUpper("bilbo")  // #1
 fmt.Println(totalLength, upperCase) // #4 리턴값 출력
}
```

```
I'm done!
5 BILBO
```

# 1.5 For, range, ...args

```go
package main

import "fmt"

func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func main() {
	superAdd(1,2,3,4,5,6)
}
```

```go
func superAdd(numbers ...int) int {
	for i:=0; i< len(numbers); i++ {
		fmt.Println((numbers[i]))
	}
	return 1
}

func main() {
	superAdd(1,2,3,4,5,6)
}
```

```go

func superAdd(numbers ...int) int {
	total :=0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1,2,3,4,5,6)
	fmt.Println(result) // 21
}
```

# 1.6 If with a Twist

### variable expressions

if를 쓰면서 변수를 새로 생성할 수 있음

```go
func canIDrink(age int) bool {
	if koreanAge := age +2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16)) // false
}
```

물론 아래의 방법처럼 변수를 선언한 후 if문을 써도 되지만 위의 방식으로 코드를 짜면
`koreanAge`를 이 조건문을 위해서만 썼다는 것을 가시적으로 명확히 보여줄 수 있음

```go
koreanAge := age +2;

if koreanAge < 20 {
	return false
}
```

# 1.7 Switch

```go
func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age > 50:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(20))
}
```

```go
func canIDrink(age int) bool {
	switch koreanAge:= age+2; {
	case koreanAge < 20:
		return false
	case koreanAge > 80:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(20))
}
```

# 1.8 Pointers

메모리 주소 보는 법: 변수 앞에 `&`

```go
func main() {
    a := 2
    b := 5
    fmt.Println(&a,&b) //0xc0000b4008 0xc0000b4010
}
```

```go
func main() {
    a := 2
    b := &a
    fmt.Println(&a,b) //0xc0000140a8 0xc0000140a8
}
```

```go
func main() {
    a := 2
    b := &a // a의 메모리 주소를 저장하는 방법: 변수에 저장 -> 해당 변수가 a의 포인터가 됨
    fmt.Println(a, &a,b, *b) // 2 0xc0000140a8 0xc0000140a8 2
    *b = 20
    fmt.Println(a) // 20
}
```

- `a` 는 2
- `b` 는 a의 메모리 주소값
- `&a` 는 a의 메모리 주소값
- `*b` 는 b를 see through: b가 현재 a의 메모리 주소값이므로 그 메모리 주소값에 저장되어 있는 값인 2가 출력
- 따라서 `*b`의 값을 바꿔주면 b가 가지고 있는 메모리 주소 안의 값(a)이 바뀜-> a도 바뀜
- 여기서 b는 a의 포인터가 됨!

# 1.9 Arrays and Slices

## 갯수가 정해진 배열을 만들 때

```go
func main() {
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "bilbo"
	names[4] = "nenya"
	// names[5] = "no" // error
	fmt.Println(names)
}
```

## 갯수가 정해지지 않은 배열을 만들 때

```go
func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "nenya") // returns new slice
	fmt.Println(names)
}
```

- `[]` 에 요소를 추가하고자 할 때는 `배열이름 = append(배열이름, 추가하고자 하는 값)` 으로 써준다.
- append가 새로운 slice를 반환하기 때문에 다시 해당 배열에 넣어주는 작업이 필요

# 1.10 Maps

```go
func main() {
	nico := map[string]string{"name":"nico", "age":"12"}
	fmt.Println(nico) // map[age:12 name:nico]
}
```

`map[key type]value type{key:value, key:value ...}`

```go
func main() {
	nico := map[string]string{"name":"nico", "age":"12"}
	for key, value := range nico {
		fmt.Println(key , value)
	}
	// name nico
	// age 12
	for _, value := range nico {
		fmt.Println(value)
	}
	// nico
	// 12
	for key, _ := range nico {
		fmt.Println(key)
	}
	// name
	// age
}
```

# 1.11 Structs

- map에서는 기본적으로 한가지 타입의 key나 value를 넣을 수 있음
- 다양한 타입을 넣고자 한다면 structs 이용

```go
type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 13, favFood}
	fmt.Println(nico) // {nico 13 [kimchi ramen]}
	fmt.Println(nico.name) // nico
}
```

하지만 이 방법은 가시적으로 그다지 명확하지 않음

```go
func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person {
		name:"nico",
		age:18,
		favFood: favFood,
        }
	fmt.Println(nico) //{nico 18 [kimchi ramen]}
}
```

- 조금 번거롭더라도 key, value를 모두 써주는 것이 보기 좋음
- 단, 하나라도 key, value를 함께 쓰고자 한다면 나머지 요소들도 모두 key, value를 함께 써줘야 함
- 어떤건 안쓰고, 어떤건 쓰는 그런 자유도가 없음
