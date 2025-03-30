## 함수를 이용한 반복
* 1.23 버전 부터 iterator 함수를 이용하여 반복문을 작성할 수 있습니다. 
  * https://tip.golang.org/doc/go1.23
* 사용할 수 있는 함수는 다음의 타입입니다. 
```go
func(func() bool)
func(func(K) bool)
func(func(K, V) bool)
```
## 사용법
* Go 언어는 표준라이브러리에서 Set을 제공하지 않지만 다음과 같이 구현해 보겠습니다. 
* 기본 제공하는 map을 이용하고 Add와 Contains 메소드만 제공합니다. 
* All() 함수를 작성하여 iter.Seq 함수를 반환합니다. iter.Seq 는 func(func(K) bool) 함수의 명칭입니다.

```go
package set

import "iter"

type Set[E comparable] struct {
	m map[E]struct{}
}

func New[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
```

## iter.Seq 를 반복문에서 사용
* 반복문의 range 에서 set.All() 메소드를 활용하면 함수를 이용한 반복문을 작성할 수 있습니다.  
```go
package main

import (
	"fmt"
	"go_for_spring_developer/01-go-basic/e6-1/set"
)

func main() {
	set1 := set.New[string]()
	set1.Add("1")
	set1.Add("2")
	set1.Add("3")
	set1.Add("4")
	set1.Add("5")
	set1.Add("6")

	for value1 := range set1.All() {
		fmt.Println(value1)
	}
}

```
## 왜 필요한가?
* 컬렉션 타입의 객체는 때에 따라 전체 요소를 순회하거나 특정 조건을 만족하는 요소를 찾아야 할 때가 있습니다.
* 이때 내부 요소를 노출시키지 않고 순회하도록 하려면 함수를 이용한 방법을 써야 합니다. 
* 이러한 방법을 사용하면 컬렉션 타입의 객체를 노출시키지 않고도 순회할 수 있습니다.
* Set 객체의 합집합을 만드는 예를 봅시다. 
```go
// S1 과 S2의 합집합을 만드는 함수
func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
    r := New[E]()
    for v := range s1.m {
        r.Add(v)
    }
    for v := range s2.m {
        r.Add(v)
    }
    return r
}
```
* 여기서는 s1 과 s2 의 내부 요소에 접근 가능하기 때문에 처리 할 수 있습니다. 하지만 Union 외 여러가지 용도의 요소 순회를 하려면 어떻게 할까요?

### Push Set Elements
* 접근법 하나를 보자면 함수를 인자로 받는 함수를 제공하면 됩니다. 이 함수는 모든 요소를 순회하며 실행됩니다. 
```go
func (s *Set[E]) Push(f func(E) bool) {
    for v := range s.m {
        if !f(v) {
            return
        }
    }
}
```
* Go 표준라이브러리에서는 sync.Map.Range, flag.Visit, filepath.Walk 등에서 이러한 방법을 사용합니다.
* 이제 이 함수를 이용하여 Set 객체의 모든 요소를 출력하는 함수를 만들어 봅시다. 
```go
func PrintAll[E comparable](s *Set[E]) {
    s.Push(func(v E) bool {
        fmt.Println(v)
        return true
    })
}
```

### Pull Set Elements
* 또 다른 전체 요소를 순회하는 방법은 함수를 반환하는 함수를 제공하는 방법입니다.  
* 반환하는 첫번째 함수는 호출될 때 마다 함수는 Set 의 요소를 반환합니다. 모든 요소를 순회하고 나면 bool 값은 false 가 됩니다.
* 두번째 함수는 순회를 중단하는 함수입니다.

```go
func (s *Set[E]) Pull() (func () (E, bool), func ()) {
    ch := make(chan E)
    stopCh := make(chan bool)

    go func () {
        defer close(ch)
        for v := range s.m {
            select {
            case ch <- v:
            case <-stopCh:
                return
            }
        }
    }()

    next := func () (E, bool) {
        v, ok := <-ch
        return v, ok
    }

    stop := func () {
        close(stopCh)
    }

    return next, stop
}
```
* 이 함수를 이용하여 Set 객체의 모든 요소를 출력하는 함수를 만들어 봅시다. 
```go
func PrintAllElementsPull[E comparable](s *Set[E]) {
    next, stop := s.Pull()
    defer stop()
    for v, ok := next(); ok; v, ok = next() {
        fmt.Println(v)
    }
}
```
## 표준 접근법
* 지금 까지 살펴본 것과 같이 표준화된 전체 요소 순회 방법이 없기 때문에 Go 개발자들은 새로운 컨테이너 패키지를 사용할때 마다 새로운 순회 메커니즘을 공부해야 합니다.
* 이러한 이유로 Go 1.23 버전에서는 이러한 문제를 해결하기 위해 함수를 이용한 반복문을 제공합니다.

### Iterator
*


