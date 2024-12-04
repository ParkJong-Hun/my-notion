# ProGuard

<aside>
💡 Android 앱의 코드 압축, 최적화, 난독화를 수행하는 도구.
ProGuard는 Java 클래스 파일을 분석해 사용되는 않는 코드를 제거하고, 변수와 메소드 이름을 난독화해, 클래스 파일 크기를 최소화해 애플리케이션 배포 크기를 줄일 수 있음.

</aside>

[R8%20d147c8a4f587405c98fd4e29b18d72f5](R8%20d147c8a4f587405c98fd4e29b18d72f5)

**proguard-rules.pro**

모듈과 그 모듈의 종속되어 있는 라이브러리의 proguard 와 연관되어 있는 규칙을 선언하는 파일.

**consumer-rules.pro**

소비자가 모듈에 적용할 수 있는 규칙을 선언하는 파일.

**minifyEnabled**

R8에 의한 코드 압축 활성화.

**-keep**

특정 클래스, 메소드, 필드 등을 보존하도록 지정. 보존하려는 항목의 패키지명, 클래스명, 시그니처 등을 지정 가능.

**-keepclassmembers**

클래스 멤버(메서드, 필드)를 보존하도록 지정.

**-dontwarn**

경고 메시지를 표시하지 않도록 지정. ProGuard는 애플리케이션을 난독화하거나 최적화할 때 경고 메시지를 발생시킬 수 있는데, 이를 무시하고 싶을 때 사용.

**-optimizationpasses**

최적화 패스 수를 지정.

**-keepattributes**

특정 속성을 보존하도록 지정.

**-printmapping**

난독화된 클래스와 멤버의 매핑 정보를 출력.

**-injars, -outjars**

입력 및 출력 JAR 파일을 지정.