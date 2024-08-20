# static vs final vs static final

**static**

- 객체 생성 없이 사용할 수 있는 필드와 메소드.
- this 키워드를 사용할 수 없음.
- **런타임**에 메모리가 할당됨.

<aside>
⚠️ Kotlin에서는 `@JvmStatic`을 사용함으로 사용할 수 있음.

</aside>

**final**

- 수정이 불가능한 변수(상수)라는 의미.
- 객체 생성 시 대입이 가능.

**static final**

- 상수라는 의미.
- 객체 생성 없이 사전 정의.
- **컴파일 타임**에 메모리가 할당됨.