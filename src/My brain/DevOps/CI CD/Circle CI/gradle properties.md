# gradle.properties

<aside>
💡 Gradle은 빌드를 실행하는 데 사용할 Java 프로세스를 쉽게 구성할 수 있는 몇 가지 옵션을 제공.
`GRADLE_OPTS` 혹은 `JAVA_OPTS`를 통해 로컬 환경을 구성할 수 있고, 전체 팀이 일관된 환경에서 작업할 수 있도록 JVM 메모리 구성, Java home 위치와 같은 특정 설정을 버전 제어에 저장할 수 있다.

</aside>

<aside>
⚠️ p옵션이 여러 위치에서 구성된 경우 다음 위치 중 하나에서 가장 먼저 발견된 옵션이 우선된다.
1. `-D` 사용을 설정한 커맨드 라인
2. `GRADLE_USER_HOME` 디렉토리의 `gradle.properties`
3. 프로젝트 디렉토리의 `gradle.properties`, 그 다음 상위 프로젝트의 디렉토리에서 빌드의 루트 디렉토리까지
4. Gradle 설치 디렉토리의 `gradle.properties`

</aside>

**org.gradle.caching=(true, false)**

[Build Cache](https://docs.gradle.org/current/userguide/build_cache.html#build_cache)

true로 설정하면 Gradle은 가능한 경우 이전 빌드의 작업 출력을 재사용하므로 빌드 속도가 훨씬 빨라짐. 기본적으로는 활성화되어 있지 않음.

**org.gradle.caching.debug=(true,false)**

[Build Cache](https://docs.gradle.org/current/userguide/build_cache.html#sec:task_output_caching)

true로 설정하면 각 작업에 대한 개별 입력 속성 해시와 빌드 캐시 키가 콘솔에 기록됨. 기본값은 false.

**org.gradle.configureondemand=(true,false)**

[Configuration time and execution time](https://docs.gradle.org/current/userguide/multi_project_configuration_and_execution.html#sec:configuration_on_demand)

Gradle이 필요한 프로젝트만 구성하려고 시도하는 온디맨드 인큐베이팅 구성을 활성화함. 기본값은 거짓

**org.gradle.console=(auto,plain,rich,verbose)**

[Command-Line Interface](https://docs.gradle.org/current/userguide/command_line_interface.html#sec:command_line_logging)

콘솔 출력 색상 또는 자세한 정도를 사용지 지정. 기본값은 Gradle이 호출되는 방식에 따라 다름.

**org.gradle.continuous.quietperiod=(# of quiet period millis)**

[Command-Line Interface](https://docs.gradle.org/current/userguide/command_line_interface.html#sec:continuous_build)

연속 빌드를 사용하는 경우 Gradle은 다른 빌드를 트리거하기 전에 조용한 기간이 지날 때까지 기다림. 이 대기 기간 내의 추가 변경 사항은 대기 기간 동안 다시 시작됨. 기본값은 250ms.

**org.gradle.daemon=(true,false)**

[The Gradle Daemon](https://docs.gradle.org/current/userguide/gradle_daemon.html#gradle_daemon)

true로 설정하면 Gradle 데몬이 빌드를 실행하는 데 사용됨.

기본값은 true이며 데몬을 사용하여 빌드가 실행됨.

**org.gradle.daemon.idletimeout=(# of idle millis)**

Gradle 데몬이 지정된 유후 시간(밀리초) 후에 종료됨. 기본 값은 3시간.

**org.gradle.debug=(true,false)**

true로 설정하면 Gradle은 원격 디버깅이 활성화된 상태로 빌드를 실행하고 포트 5005에서 수신 대기. 이것은 JVM 명령줄에 `- agentlib:jdwp=transport=dt_socket,server=y,suspend=y,address=5005` 를 추가하는 것과 동일하며 디버거가 연결될 때까지 가상 머신을 일시 중단. 기본값은 거짓.

**org.gradle.debug.host=(host address)**

디버그가 활성화된 경우 수신하거나 연결할 호스트 주소를 지정. Java 9 이상의 서버 모드에서 호스트에 *를 전달하면 서버가 모든 네트워크 인터페이스에서 수신 대기. 기본적으로 호스트 주소는 JDWP에 전달되지 않으므로 Java 9 이상에서는 루프백 주소가 사용되는 반면 이전 버전은 모든 인터페이스에서 수신 대기.

**org.gradle.debug.port=(port number)**

디버그가 활성화된 경우 수신 대기할 포트 번호를 지정. 기본값은 5005.

**org.gradle.debug.server=(true,false)**

true로 설정하고 디버깅이 활성화되면 Gradle은 디버거의 소켓 연결 모드로 빌드를 실행. 그렇지 않으면 소켓 수신 모드가 사용됨. 기본값은 참.

**org.gradle.debug.suspend=(true,false)**

true로 설정되고 디버깅이 활성화되면 Gradle을 실행하는 JVM이 디버거가 연결될 때까지 일시 중단됨. 기본값은 참.

**org.gradle.java.home=(path to JDK home)**

[Build Environment](https://docs.gradle.org/current/userguide/build_environment.html#sec:gradle_environment_variables)

Gradle 빌드 프로세스에 대한 Java 홈을 지정. 값은 jdk 또는 jre 위치로 설정할 수 있지만 빌드가 수행하는 작업에 따라 JDK를 사용한느 것이 더 안전. 이는 Gradle 클라이언트 VM을 시작하는 데 사용되는 Java 버전에 영향을 주지 않음. 설정이 지정되지 않은 경우 적절한 기본값을 환경(JAVA_HOME 또는 java 경로)에서 파생.

**org.gradle.jvmargs=(JVM arguments)**

[Build Environment](https://docs.gradle.org/current/userguide/build_environment.html#sec:configuring_jvm_memory)

Gradle Daemon에 사용되는 JVM 인수를 지정. 이 설정은 빌드 성능을 위해 JVM 메모리 설정을 구성하는 데 특히 유용. 이는 Gradle 클라이언트 VM의 JVM 설정에 영향을 주지 않음. 기본값은 -Xmx512m “-XX:MaxMetaspaceSize=384m”.

**org.gradle.logging.level=(quiet,warn,lifecycle,info,debug)**

queit, warm, lifecycle, info 또는 debug로 설정하면 Gradle이 이 로그 수준을 사용. 값은 대소문자를 구분하지 않음. lifecycle이 기본값.

**org.gradle.logging.stacktrace=(internal,all,full)**

예외 발생 시 스택 추적을 빌드 결과의 일부로 표시할지 여부를 지정. `internal`로 설정하면 내부 예외가 발생한 경우에만 스택 추적이 출력에 표시됨. `all` 또는 `full`로 설정하면 모든 에외 및 빌드 실패에 대한 출력에 stacktrace가 표시됨. `full`을 사용하면 stacktrace가 잘리지 않으므로 훨씬 더 자세한 출력이 생성됨. 기본값은 `internal`.

**org.gradle.parallel=(true,false)**

[Improve the Performance of Gradle Builds](https://docs.gradle.org/current/userguide/performance.html#parallel_execution)

구성되면 Gradle은 최대 org.gradle.workers.max로 JVM을 분기해 프로젝트를 병렬로 실행. 기본값은 거짓.

**org.gradle.priority=(low,normal)**

[Command-Line Interface](https://docs.gradle.org/current/userguide/command_line_interface.html#sec:command_line_performance)

Gradle 데몬에서 실행되는 모든 프로세스의 스케줄링 우선순위를 지정. 기본값은 normal

**org.gradle.vfs.verbose=(true,false)**

파일 시스템을 감시할 때 상세 로깅을 구성. 기본값은 거짓.

**org.gradle.vfs.watch=(true,false)**

파일 시스템 보기를 토글. 활성화되면 Gradle은 빌드 간에 파일 시스템에 대해 수집한 정보를 재사용. Gradle이 이 기능을 지원하는 운영 체제에서는 기본적으로 활성화됨.

**org.gradle.warning.mode=(all,fail,summary,none)**

`all`, `summary`, `none`으로 설정하면 Gradle은 다른 경고 유형 표시를 사용.

**org.gradle.welcome=(never,once)**

Gradle이 환영 메시지를 인쇄할지 여부를 제어. `never`로 설정하면 환영 메시지가 표시되지 않음. `once`으로 설정하면 Gradle의 새 버전마다 메시지가 한번씩 인쇄됨. 기본 값은 `once`

**org.gradle.workers.max=(max # of worker processes)**

구성된 경우 Gradle은 지정된 작업자 수의 최대값을 사용.