# Rust

<aside>
💡

모질라 재단에서 개발한 저수준 시스템 프로그래밍 언어.
메모리 관리 안전성 강조.

</aside>

[Cargo%20125f37315c44800086a9c4fc48d7b453](Cargo%20125f37315c44800086a9c4fc48d7b453)

### Type

**u**

- unsigned integer.

**i**

- signed integer.
- 타입을 정의하지 않고 정수를 작성하면 기본적으로 i32.

**f**

- floating point.
- 웬만하면 계산 안정성을 위해 64가 좋음.

**&str**

- string slice(string literal).
- 문자열을 담는 불변형 타입.

**tuple**

- 여러 값을 하나의 단위로 묶어 표현할 수  있는 타입.

<aside>
⚠️

char는 무조건 4bytes.
u8으로 캐스팅 가능.

숫자 뒤에 u8 등 타입을 기입하면 해당 타입으로 인식됨.

</aside>

### Keyword

**let**

- 런타임 변수 선언.
- 기본적으로 immutable.

**const**

- 컴파일 타임 상수 선언.

**static**

- 전역 변수 선언.

**mut**

- mutable variable.

**&mut**

- mutable reference.
- 변경 가능하도록 참조를 전달.

**{}**

- display print.
- 문자열 안에 다른 타입의 값을 넣을 때.

**fn name(paramters) → returnType { code }**

- function.
- 내부 코드는 return이 없으면 마지막 작성한 내용이 return.

**()**

- empty tuple.
- unit type(void).

**{:?}**

- debug print.
- 문자열 안에 다른 타입의 type 이름이나 프로퍼티.

**unsafe**

- 플랫폼 의존적 코드나 메모리에 직접 접근 하는 코드의 wrapper.
- 생 포인터에 역참조.
- unsafe 함수, 메소드 호출.
- mutable static에 접근, 수정.

**trait**

- 인터페이스.
- impl로 다른 타입으로서 사용할 수 있도록 구현 가능.

**match**

- 스위치문.

**struct**

- 구조체.
- **unit struct**
    - 메모리 0인 아무것도 가지지 않는 struct

**enum**

- 열거형.

**loop**

- 반복문.

**impl**

- 구조체나 열거형, Trait 생성, 메서드 작성을 담당하는 것.
- 기존에 있던 것을 현재 구조체 등에서 override 가능.

**Self**

- 자기 자신.

**use**

- Enum을 사용할 때, 타입 이름 생략 가능하도록 선언.

**where**

- 복잡한 generics 선언.

**?**

- 매치 가능하면 매치, 아니면 return Error.

**from**

- 한 타입을 다른 타입으로 변환.

**|매개변수| 처리**

- 클로저.

**dbg!**

- debug quick print.
- 디버깅용 프린트.

**‘라이프타임이름**

- 라이프타임.
- 이미 free된 참조를 이용하려고 하는 after use 문제를 해결하기 위해 등장한 패러다임.
- 함수에서 선언한 라이프타임을 참조(&)에서 사용함.
- **‘static**
    - 프로그램 전체의 라이프타임.

**mod**

- module.
- mod이름::내부struct 식으로 사용 가능.

**pub**

- public.
- module 내의 로직을 외부에 노출.

### Trait

**AsRef<T>**

- 어떤 타입을 &T 참조로 변환.
- 소유권 변경 x

**Iterator**

- next를 실행할 수 있는 컬렉션.

### Struct

**String**

- 문자열.

**Vec<>**

- vector.
- 사이즈 변경 가능한 배열 타입.

**HashMap**

- map type.

**BTreeMap**

- 순서가 있는 map type.

**HashSet**

- set type.

**BinaryHeap**

- 2진 Heap Type.

**VecDeque**

- 양쪽 끝에서 항목 추가 가능.
- 기본적으로 Vec보다 성능이 안좋음.

**Cell**

- 내부에서 가지는 값을 변경할 수 있도록하는 구조체.
- 일반적인 Cell은 thread safe하지 않음.

**RefCell**

- 값을 직접 복사(get, set)하지 않고, 런타임에서 borrow를 체크(Ref/RefMut)하는 Cell.
- get할 때 Ref 혹은 RefMut으로만 접근 가능.
- set할 때 borrow_mut()을 사용해서 참조 생성 후 값을 변경해야 함.
    - borrow (&)은 여러 번 사용 가능.
        - 불변 참조이므로
    - borrow_mut(&mut)은 한 번만 가능.
        - 가변 참조이므로
        - 여러번 가변 참조를 생성하면 런타임 에러(패닉)이 발생.
        - borrow가 있는데, borrow_mut을 생성하면 런타임 패닉 발생.
- 주로 String이나 Vec<T> 같은 복사 불가능한 데이터를 내부적으로 변경할 때 사용.

<aside>
⚠️

빌림에 이런 규칙이 있는 이유는, 데이터 경쟁과 메모리 안전성을 보장하기 위해.

</aside>

**Rc**

- Reference counting.
- 스마트 포인터.
- 일반 clone과 다르게 값을 복사하는 것이 아닌, 참조 횟수만 증가해 성능에 영향 없음.
- 여러 Rc 인스턴스가 같은 데이터를 가리킬 수 있음(읽기 전용이므로).
- 내부적으로 참조 횟수를 추적해, 마지막 Rc가 드롭될 때 데이터가 해제됨.
- 단일 스레드에서만 사용 가능.

**Method**

**ok**

- Ok(값) 혹은 Error(값)을 반환

**ok_or**

- Ok(값) 혹은 Error(지정한 값)을 반환

**and_then**

- also 같은 것

**drop**

- Rc 객체를 하나 더 이상 사용하지 않음을 의미.