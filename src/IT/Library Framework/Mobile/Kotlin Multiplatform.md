# Kotlin Multiplatform

<aside>
💡 KMP(KMM, MPP) 특히 KMM은 모바일 애플리케이션 개발을 위한 Kotlin의 기술.

Kotlin을 사용하여 안드로이드 및 iOS 플랫폼에서 동작하는 앱을 개발하기 위한 도구와 라이브러리 집합.

개발자가 Kotlin으로 공통 코드를 작성하고, 이를 여러 플랫폼에서 공유하면서 개발 생산성을 높일 수 있도록 도와줌. 

공통 코드는 비즈니스 로직, 데이터 모델, 네트워킹, 로깅 등과 같은 앱의 핵심 부분을 포함할 수 있음. 

플랫폼별 코드는 안드로이드 및 iOS에 대한 특정 기능 및 UI/UX를 처리하는 데 사용.

KMM을 사용하면 코드의 재사용성이 향상되고 개발 시간이 단축되며, 앱을 여러 플랫폼에 배포하기 쉬워짐. 

또한 Kotlin의 강력한 기능과 생산성 도구를 활용하여 개발할 수 있으므로 개발자들 사이에서 인기가 있는 기술.

</aside>

### **사용하는 프로그래밍 언어**

[Kotlin%20d3702579c4774777b0e880ce2ebbdae8](Kotlin%20d3702579c4774777b0e880ce2ebbdae8)

[KDoctor%20d307d53bf8494e4e9f2c4f2acf413f15](KDoctor%20d307d53bf8494e4e9f2c4f2acf413f15)

[Kodein%207ee8d6e8662a4e03b152387b95aec1b3](Kodein%207ee8d6e8662a4e03b152387b95aec1b3)

[BuildKonfig%20d74b6015badd4a8fae491ef1b9c240fc](BuildKonfig%20d74b6015badd4a8fae491ef1b9c240fc)

[CrashKiOS%2013ef37315c4480138299cef3c25b7a48](CrashKiOS%2013ef37315c4480138299cef3c25b7a48)

[Kotlin%20Multiplatfrom%20Gradle%20plugin%20151f37315c4480929fa0c79919f8b371](Kotlin%20Multiplatfrom%20Gradle%20plugin%20151f37315c4480929fa0c79919f8b371)

### iOS에 종속성 추가 방법

**직접 통합**

기본. 

iOS 앱 빌드에 새 실행 스크립트 단계를 추가해 직접 프레임워크를 연결.

**CocoaPods 통합**

로컬, 원격 종속성. 

Swift와 Objective-C 프로젝트에서 주로 사용되는 CococPods를 이용해 프레임워크를 연결.

**SPM 사용**

Swift 코드 배포를 관리하는 SPM으로 프레임워크를 연결.

XCFramework를 사용해 SPM에 종속성 설정 가능.

### 여러 모듈 공유

iOS의 경우 KMP 모듈에서 생성된 하나의 프레임워크에 의존.

여러 모듈을 사용하는 경우 모든 모듈에 따라 추가 모듈인 **엄브렐라 모듈**을 추가하고 모든 모듈을 포함하는 프레임워크인 **엄브렐라 프레임워크**를 구성해야 함.

### Module

**shared**

Android와 iOS 앱 모두에 공통적인 로직.

**androidMain**, **commonMain**, **iosMain**  소스 세트로 구성됨.

**expect**, **actual** 문법을 사용해 각 플랫폼 코드를 구현

**composeApp**

Android 앱에 빌드되는 Kotlin 모듈.

**iosApp**

iOS 앱에 빌드되는 Xcode 프로젝트.

[Swift%20Objective-C%20%E1%84%89%E1%85%A1%E1%86%BC%E1%84%92%E1%85%A9%20%E1%84%8B%E1%85%AE%E1%86%AB%E1%84%8B%E1%85%AD%E1%86%BC%E1%84%89%E1%85%A5%E1%86%BC%20f01a5f4a054f4502845ef07c47dfc8d2](Swift%20Objective-C%20%E1%84%89%E1%85%A1%E1%86%BC%E1%84%92%E1%85%A9%20%E1%84%8B%E1%85%AE%E1%86%AB%E1%84%8B%E1%85%AD%E1%86%BC%E1%84%89%E1%85%A5%E1%86%BC%20f01a5f4a054f4502845ef07c47dfc8d2)

[Compose%20Multiplatform%2024b683c77035484a9f9637e2995accad](Compose%20Multiplatform%2024b683c77035484a9f9637e2995accad)