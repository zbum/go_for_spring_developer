## Go 언어의 특징 (What makes Go great?)
* Go로 작성된 코드는 읽기가 쉽다.
  * 마술과 같은 언어가 아니다.
* 어떤것도 숨어있지 않다.
  * 어떤 함수가 어떤 함수에 의해서 어떤 순서로 실행되는지 찾기가 쉽다.
* 네이티브 바이너리를 생산한다.
  * VM 이나 공유 라이브러리가 설치되었는지를 걱정할 필요가 없음

## Bean
* 저런 특징 때문에 프레임워크가 자동으로 처리하는 부분이 극히 없음
* 따라서 Bean 을 작성하려면 직접 객체를 생성하고 의존성을 부여해야 합니다.
* bean scope 속성도 개발자가 직접 제어
<br />

## 스프링 프레임워크의 Bean
* 스프링 Bean 은 스프링 프레임워크가 관리하는 재사용 가능한 컴포넌트입니다. 
* Java Configuration 이나 @Component, @ComponentScan 과 같은 애너테이션을 기반으로 IoC 와 DI 를 수행합니다.
* 생성자에 주입받을 빈의 타입으로 인자를 선언하면 스프링 프레임워크가 자동으로 의존성을 주입합니다. 
* 다음의 사용자를 관리하는 스프링 빈을 선언해 봅시다. 
```java
@Service
public class ScoreService {

    private final ScoreRepository scoreRepository;

    public ScoreService(ScoreRepository scoreRepository) {
        this.scoreRepository = scoreRepository;
    }

    public List<Score> getScores() {
        return scoreRepository.findAll();
    }

    public List<Score> getScoresByStudentId(Long studentId) {
        return scoreRepository.findByStudentId(studentId);
    }

    @Transactional
    public Score addScoresByStudentId(Score score) {
        return scoreRepository.save(score);
    }
}
```
* ScoreService 클래스는 ScoreRepository 타입의 빈에 의존하고 있습니다. 
* 스프링 프레임워크의 생성자 타입의 의존성 주입을 활용해서 주입을 받고 있습니다. 
* getScores 메서드에서 ScoreRepository findAll 메서드를 사용합니다.

## Go 로 작성해보는 DI
* 이번에는 Go 언어로 ScoreService, ScoreRepository 를 개발해 봅시다.
* 아직 데이터베이스 접근과 ORM을 배우지 않았으니 고정값을 반환하도록 개발합니다.


### Score Repository 빈
```go
package bean

import "fmt"

type ScoreRepository struct {
}

func (u *ScoreRepository) FindAll() []Score {
    return []Score{
        {1, 1, Student{1, "Manty"}, 100},
        {2, 2, Student{1, "Manty"}, 99},
        {3, 3, Student{1, "Manty"}, 98},
    }
}

func NewScoreRepository() *ScoreRepository {
    fmt.Println("init UserRepository")
    return new(ScoreRepository)
}
```

* 구조체를 선언합니다. 
* FindAll() 이라는 타입 메소드를 선언합니다. 전 시간에 배운 것처럼 리시버를 사용해야 합니다. 
* 보통 생성자 대신 NewScoreRepository 형식의 팩토리 메소드를 작성하여 생성합니다.

## Dependency Injection
* ScoreRepository에 의존하는 ScoreService 객체를 작성해 보겠습니다. 
```go
package bean

type ScoreService struct {
    scoreRepository *ScoreRepository
}

func NewScoreService(scoreRepository *ScoreRepository) *ScoreService {
    return &ScoreService{scoreRepository}
}

func (s *ScoreService) GetScores() []Score {
    return s.scoreRepository.FindAll()
}

```

* ScoreService 스트럭트는 `scoreRepository *ScoreRepository` 필드를 가지고 있습니다. 
* 팩토리 메소드는 의존성 주입을 받기 위한 `*ScoreRepository` 를 인자로 가지고 있고 ScoreService 초기화 할때, `*ScoreRepository` 타입의 인자를 사용합니다.


> 10장에서 Dependency Injection 을 하는 외부 라이브러리를 확인해 보겠습니다. 
