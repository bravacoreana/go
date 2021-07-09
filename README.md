# 1.0 main package

- `main.go` 파일명은 프로젝트를 **컴파일** 하려면 반드시 가지고 있어야 함. 그게 아니라면 바꿔도 상관 x
- 왜냐하면 `main.go` 가 entry point 임 -> 컴파일러가 가장 먼저 찾는 패키지 이름

# 1.1 packages and imports

- fmt: formatting 을 위한 패키지

- 패키지가 어떻게 동작하나
- 왜 패키지 함수는 대문자로 시작하나
  - export 하고 싶으면 함수를 대문자로 시작해주면 됨
