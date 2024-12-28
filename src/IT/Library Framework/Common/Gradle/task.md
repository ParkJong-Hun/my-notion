# task

**register**

- 동적으로 지연 Task 등록

**create**

- Task 즉시 등록
- 주로 초기 버전에서 사용

**by tasks.creating**

- Kotlin DSL에서 Task 등록

**named**

- 기존에 있던 Task에 추가 작업을 등록할 때 사용

**dependsOn**

- 한 Task가 다른 Task를 실행하도록 명시

**doLast**

- dependsOn에서 실행한 Task가 끝나면 실행할 것