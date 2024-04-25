# Junit

<aside>
💡 Java 언어 기반 테스트 프레임워크.
Android 앱의 유닛 테스트 작성에 가장 기본적으로 사용.
테스트 케이스를 정의하고 실행할 수 있음.

</aside>

**@Test**

테스트 메서드 정의. 테스트 실행 시 자동 호출.

**@Before**

각 테스트 메서드가 실행되기 전 실행될 메서드 정의.

**@After**

각 테스트 메서드가 실행된 후 실행될 메서드 정의.

**@BeforeClass**

테스트 클래스 전체의 실행 전 실행될 메서드 정의.

**@AfterClass**

테스트 클래스 전체의 실행 후 실행될 메서드 정의.

**assertEquals, assertSame, assertNotEquals**

테스트 결과 값을 검증하는 데 사용.

**assertTrue, assertFalse**

주어진 조건이 참인지 거짓인지 검증하는 데 사용.

**assertNull, assertNotNull**

주어진 값이 null인지 아닌지 검증하는 데 사용.

**assertThrows**

특정 예외가 발생하는지 확인하는 데 사용.

**assertTimeout, assertTimeoutPreemptively**

테스트 메서드의 실행 시간을 제한하는 데 사용.

**@Ignore**

특정 테스트 메서드를 실행에서 제외하고 싶을 때 사용.

테스트를 일시적으로 비활성화.