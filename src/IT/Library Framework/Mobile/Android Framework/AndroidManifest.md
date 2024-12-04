# AndroidManifest

<aside>
💡 Android 빌드 도구, Android OS, Google Play에 앱에 관한 필수 정보를 설명.
Manifest, Application을 반드시 명시.

</aside>

**Manifest**

- <manifest>

**Application**

- <application>

### Component

- <activity>
- <service>
- <receiver>
- <provider>
    - android:exported
        - 다른 앱의 컴포넌트에서 실행할 수 있는 컴포넌트인지

**Intent Filter**

- <intent-filter>
    - 인텐트를 시스템에 발행하면 시스템은 각 앱의 매니페스트 파일에 선언된 이 것을 보고 처리할 수 있는 컴포넌트를 찾음.
    - 일치하는 컴포넌트의 인스턴스를 시작해 해당 컴포넌트에 Intent 객체를 전달.
    - 두 개 이상의 앱이 인텐트 처리가 가능한 경우 사용자는 어느 앱을 사용할지 선택 가능.
    - 하위로 <action>, <category>, <data>를 추가 작성.

**Icon, Label**

**Permission**

- <uses-permission>

**Feature**

- <uses-feature>
    - 앱에 필요한 하드웨어, 소프트웨어 기능.
- <uses-sdk>
    - build.gradle을 대신 사용해 SDK 버전 지정.
    

**Other**

- <activity-alias>
    - Activity 별칭 선언.
- <compatible-screen>
    - 앱이 호환되는 각 화면 구성 지정.
- <grant-uri-permission>
    - 상위 Content provider에게 액세스 권한이 있는 앱 데이터의 하위 집합을 지정.
- <instrumentation>
    - 앱과 시스템의 상호작용을 모니터링할 수 있는 클래스 선언.
- <meta-data>
    - 상위 컴포넌트에게 제공될 수 있는 추가 임의 데이터 항목의 이름-값 쌍.
- <path-permission>
    - Content provider 내의 특정 데이터 하위 집합에 관해 경로와 필수 권한 정의.
- <permission>
    - 애 앱이나 다른 앱의 특정 컴포넌트 또는 기능에 대한 액세스를 제한하는 데 사용될 수 있는 보안 권한.
- <permission-group>
    - 관련 권한의 논리적인 그룹 이름.
- <permission-tree>
    - 권한 트리의 기본 이름.
- <provider>
    - Content provider 선언.
- <queries>
    - 앱에서 액세스하려는 다른 앱 집합 선언.
- <supports-gl-texture>
    - 앱에서 지원하는 단일 GL 텍스처 압축 형식 선언.
- <supports-screens>
    - 앱이 지원하는 화면 크기를 선언해 앱이 지원하는 것보다 큰 화면에 대해 화면 호환성 모드 활성화.
- <uses-configuration>
    - 앱이 요구하는 특정 입력 기능.
- <uses-library>
    - 앱이 연결되어야 하는 공유 라이브러리 지정.
- <uses-native-library>
    - 앱이 연결되어야 하는 공급업체 제공 네이티브 공유 라이브러리 지정.
- <uses-permission-sdk-23>
    - 앱이 특정 권한을 원한다는 것을 지정. Android 6.0 이상을 실행하는 앱이 설치된 경우만 해당.