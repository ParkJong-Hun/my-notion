# ./gradlew

<aside>
💡 Gradle 빌드 시스템을 실행하는 데 사용되는 스크립트.

</aside>

**build**

프로젝트 빌드. 소스 코드를 컴파일하고 테스트를 실행해 앱 생성.

**clean**

빌드된 파일 및 디렉터리 제거.

**tasks**

사용 가능한 모든 태스크 나열.

**dependencies**

프로젝트의 종속성 나열.

**—info, —debug, —stacktrace**

빌드 과정 정보, 디버그 정보, 스택 트레이스 등 추가 정보 표시.

**—Dproperty=value**

시스템 속성 설정. -`Dorg.gradle.daemon=false`는 Gradle 데몬을 사용하지 않도록 함.

**—refresh-dependencies**

종속성 캐시를 강제로 업데이트 해 외부 종속성을 새로 다운로드.

**—offline**

오프라인 모드에서 빌드 수행.

**-Dkotlin.compiler.execution.strategy**

Kotlin 컴파일러 실행 전략 지정. 외부 컴파일러에서 실행하면 일반적으로 안정성과 성능 면에서 좋지만, 대규모 프로젝트나 특정 환경에서는 오버헤드가 발생하지만, in-process 전략은 Gradle 빌드 프로세스 내에서 직접 실행해 일부 빌드 성능 향상을 기대할 수 있지만 안정성 문제가 발생할 수 있음.

**-Dorg.gradle.jvmargs**

Gradle을 실행하는 동안 JVM에 대한 옵션 지정. `-Xmx`는 JVM의 최대 힙 메모리 크기 지정. `-XX`는 어떤 특정 에러가 발생할 때 힙 덤프를 생성하도록 지시하는 옵션.