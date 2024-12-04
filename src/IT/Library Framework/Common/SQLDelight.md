# SQLDelight

<aside>
💡 Kotlin 전용 라이브러리.
Kotlin DSL을 사용해 데이터베이스 스키마를 정의.
SQL을 직접 작성하고 이를 Kotlin 코드와 통합해 사용.
컴파일러를 사용해 Kotlin 코드를 생성하므로 정적 타입 검사가 가능.
SQL 쿼리가 컴파일 타임에 확인되며, IDE에서 코드 완성과 오류 검사를 제공.
SQL 스키마를 기반으로 Kotlin 클래스 생성.
복잡한 데이터베이스 작업을 처리하는 데 더 적합.

</aside>

**SqlDriver**

- DB 연결 관리 인터페이스.
- 각 플랫폼에 따라 구현클래스가 다름.

**SqlScheme**

- Scheme 정의.
- DB 초기화.

**Transacter**

- 트랜잭션 처리.
- DB 작업 그룹화.

**Transaction**

- 트랜잭션 정의.
- 시작, 커밋, 롤백 작업 정의.

**Query**

- SQL 쿼리 실행, 결과 반환.

**sq**

- SQL 스키마, 쿼리 정의 파일.

**sqm**

- 마이그레이션, 스키마 버전 정보를 포함하는 메타 데이터 파일.

**db**

- 실제 DB 파일.