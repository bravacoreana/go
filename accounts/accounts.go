package accounts

// Account struct
type Account struct {
	owner   string // 이것도 첫글자가 소문자면 private
	balance int
}

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	// &account // {bilbo 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// receiver: 어떤 구조체를 전달받을 지 명시
	// receiver 만드는 규칙: struct의 첫글자를 따서 소문자로 지어야 함
	// (a Account): 메소드와 동등한 관계
	// fmt.Println("Gonna deposit", amount)
	a.balance += amount

}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
