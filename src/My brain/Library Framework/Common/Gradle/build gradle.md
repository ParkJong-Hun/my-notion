# build.gradle

<aside>
💡

Gradle 빌드 스크립트를 작성해 프로젝트 빌드를 구성.

각 프로젝트 디렉토리에 별도로 위치, 프로젝트의 컴파일, 의존성, 작업, 플러그인 적용.

KTS에서는 CommonExtension을 상속한 것을 사용.

</aside>

- **plugins**
    - Gradle 플러그인을 적용하고 관리.
    - 플러그인을 추가하고 구성.
- **android**
    - BaseExtension.
    - com.android.application으로 사용 가능.
    - Android 앱 빌드 설정 정의.
        - 앱 버전, 미니멈 SDK 버전, 타겟 SDK 버전, 서명 설정 등
    - **buildFeatures**
        - 앱 빌드 시 사용할 수 있는 특정 기능 구성.
            - dataBinding, viewBinding, compose, m1ModelBinding, buildConfig 등
    - **productFlavors**
        - 앱을 여러 버전 혹은 여러 파생 앱(Flavor)으로 구성할 때 사용.
        - 앱 이름, 앱 아이콘, 리소스 디렉토리, 빌드 버전 코드 등 구성 가능.
    - **buildTypes**
        - 빌드 구성에 대한 다양한 설정 정의.
            - ex. debug, release
- **kotlin**
    - Kotlin 설정과 의존성을 관리.
        - Kotlin 버전, 플러그인 구성 포함
- **java**
    - Java 설정과 의존성을 관리.
        - Java 버전, 플러그인 구성 포함
- **sourceSets**
    - 프로젝트의 소스 코드와 리소스 디렉토리 정의.
    - 빌드 프로세스에서 어떻게 처리되어야 하는지 정의.
- **tasks**
    - Gradle 작업을 정의하고 구성.
    - 사용자 지정 작업, 빌드 단계 정의.
    - **register**
        - 커스텀 Task 등록.
- **dependencies**
    - Gradle에서 작업 간 종속성을 관리.
    - 한 작업이 다른 작업에 의존하도록 지정 가능.
    - **implementation**
        - 컴파일, 실행 시 classpath가 추가됨.
        - module 내부에서만 사용 가능. 외부에 노출되지 않으며 다른 모듈에서는 참조 불가능.
    - **testImplementation**
        - 테스트 시에만 필요한 의존성
        - 주로 테스트 코드에서 사용됨.
        - 애플리케이션 빌드에는 포함하지 않음.
    - **kapt**
        - Kotlin 코드의 Annotation을 처리 해야하는 의존성을 설정할 때 사용.
    - **ksp**
        - kapt의 대안.
        - 더 좋음.
    - **compileOnly**
        - 컴파일 시에만 필요한 의존성.
        - 런타임에는 필요하지 않음.
        - 컴파일 시에만 클래스패스에 추가되고, 실제 실행에는 포함되지 않음.
    - **api**
        - 설정한 모듈을 참조하는 다른 모듈에서도 참조 가능한 의존성.
    - **testRuntimeOnly**
        - 테스트 실행 중에만 필요한 의존성
        - 테스트 실행 시 런타임 클래스패스에 추가됨
        - 애플리케이션 실행에는 포함되지 않음
    - **androidTestImplementation**
        - 안드로이드 애플리케이션의 테스트를 위해 필요한 의존성
        - 통합 테스트에 사용됨
    - **debugImplementation**
        - 디버그 빌드 구성에서만 필요한 의존성
        - 릴리스 빌드에 포함되지 않음
        - 주로 디버그용 라이브러리
    - **releaseImplementation**
        - 릴리스 빌드 구성에서만 필요한 의존성
        - 앱이 릴리스될 때는 사용되는 라이브러리를 추가할 때 사용
- **framework**
    - Objective-C 프레임워크를 생성하고 구성.
    - **export**
        - 데이터나 모듈을 쓰기
- **configurations**
    - implemenatation이나 api로 추가한 라이브러리들
- **KotlinMultiplatformExtension**
    - KMP 모듈에 사용
- **CommonExtension**
    - Android App, Library, Dynamic Feature에 사용하는 IF
- **targets.filterIsIstance<KotlinNativeTarget>()**
    - 플랫폼OS를 추출
- **publishLibraryVariants**
    - 라이브러리 variants를 정의
- **publishAndroidReleasePublicationToMavenRepository**
    - publishing{}에 정의한 곳으로 릴리스 라이브러리를 공개
- **publishAndroidDebugPublicationToMavenRepository**
    - publishing{}에 정의한 곳으로 디버그 라이브러리를 공개