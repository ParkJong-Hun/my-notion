# @Stable과 @Immutable

# **Stable로 간주되는 타입**

- 원시 타입.
- 람다식으로 표현되는 함수 타입.
- class의 public 프로퍼티가 모두 불변이거나 stable.
- class에 @Stable @Immutable 등을 사용해 명시적으로 stable 하다고 표기된 경우.

# Unstable로 간주되는 타입

- 컴파일 시 구현체를 예측할 수 없는 Any 유형 같은 추상 클래스, List, Map, 등을 포함한 모든 인터페이스.
- class의 public 프로퍼티 중 최소 하나 이상이 가변적이거나 unstable한 경우.

# class

**@Stable**

- 데이터가 안정적임(입력값이 같으면, 결과값도 같음)을 나타냄.
- 해당 어노테이션이 적용된 class는 Compose compiler가 내용을 관찰해 recomposition을 행함.
- 단, 해당 인스턴스의 내용이 자주 변경된다면 오히려 퍼포먼스에 악영향이 있을 수 있음.

<aside>
⚠️ 클래스 자체나 인터페이스의 구현체가 안정적이지 않을 수 있는 경우 사용.

</aside>

**@Immutable**

- 데이터가 불변성(Immutable)을 가지고 있음을 나타냄.
- data class에 적용되며, 클래스가 불변성을 보장하는 데이터를 보유하고 있는 경우 사용.
- 해당 어노테이션이 적용된 data class는 동일한 데이터가 입력되면 항상 동일한 출력을 생성하는 것을 보장.
- 데이터가 변경되면 Compose는 변경된 데이터를 기반으로 변경된 부분만 다시 렌더링.
- recomposition 수를 줄일 수 있음.
- Immutable이 될 수 있는 클래스는 예를 들어, 사용자, 위치, 이벤트 등이 있음.
- 기본적으로 var가 존재하는 class, map, list, set은 Mutable로 판단함.
- Kotlinx immutable collections를 사용하는 것이 대안이 될 수 있음.

<aside>
⚠️ 모든 public 프로퍼티에 대하여 val을 사용해 불변임을 확인. 커스텀 setter 사용을 피함.

</aside>

# fun

**Skippable**

- recomposition이 호출될 때, 모든 파라미터가 이전의 값과 같다면 skip 가능한 함수.

Restartable

- recomposition이 시작될 수 있는 엔트리 포인트 함수.