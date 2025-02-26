# Normalization

<aside>
💡

정규화.
이상현상이 있는 릴레이션을 분해해 이상현상을 없애는 과정.

</aside>

### 장점

- 이상 현상 제거.
- 새 데이터 타입 추가 시, 구조를 변경하지 않아도 됨.
- DB와 연결된 앱이 최소한의 영향만을 미침.

### 단점

- 릴레이션 간 JOIN 연산 증가.
- 경우에 따라 응답 시간이 느려질 수 있음.
- JOIN이 많이 발생하면 반정규화를 적용할 수 있음.

### 제 1 정규화(1NF)

- 테이블의 컬럼이 원자값을 갖도록 테이블을 분해.

### 제 2 정규화(2NF)

- 완전 함수 종속(기본키의 부분집합이 결정자가 되어선 안됨)을 만족하도록 테이블을 분해.

### 제 3 정규화(3NF)

- 이행적 종속을 없애도록 테이블을 분해.

### BCNF(Boyce-codd Normal Form)

- 모든 결정저가 후보키 집합에 속한 정규형.