# JUnit5

<aside>
💡

2017년 출시된 JUnit4 확장 버전.

</aside>

### 주요 컴포넌트

- **JUnit Platform**
    - 다양한 테스트 프레임워크를 실행할 수 있는 플랫폼.
- **JUnit Jupiter**
    - JUnit5의 새 테스트 모델.
    - 최신 기능, 어노테이션 제공.
- **JUnit Vintage**
    - JUnit4와 같은 오래된 테스트도 실행 가능하도록 지원.

**@Test**

**@BeforeEach**

- 각 테스트 실행 전에 처리.

**@AfterEach**

- 각 테스트 실행 후에 처리.

**@BeforeAll**

- 모든 테스트 메서드 실행 전 한 번만 실행.
- 이전의 Before와 동일.

**@AfterAll**

- 모든 테스트 메서드 실행 후 한 번만 실행.
- 이전의 After와 동일.

**@EnabledIf**

- 조건이 충족된 경우에만 테스트 실행.

**@DisabledIf**

- 조건이 충족된 경우에 테스트 비활성화.

**@Tag**

- 테스트 그룹화.
- 태그로 특정 그룹의 테스트만 실행 가능.

**@TestInstance**

- 클래스 수준에서 단 한번만 인스턴스를 생성 가능.
- 인스턴스 생성 방식 설정.

**@ParamterizedTest**

- 매개변수화된 테스트 실행에 사용.
- 동일한 입력값에 대해 여러 번 실행 가능.
- 여러 소스 제공.
    - @ValueSource
    - @CsvSource
    - @MethodSource