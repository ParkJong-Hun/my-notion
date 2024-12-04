# Kotlin Multiplatfrom Gradle plugin

### 최상위 블록

**<targetName>**

- Targets 중에서 프로젝트에서 사용할 타겟 선언.

**targets**

- 프로젝트에서 선언된 모든 대상 리스트.
- **Targets**
    - 지원되는 플랫폼 중 목표로 하는 것.
    - jvm, wasmJs, wasmWasi, js, macosX64, macosArm64, iosSimulatorArm64, iosX64, iosArm64, androidTarget이 있음.

**sourceSets**

- 프로젝트에 선언할 소스 세트 선언.
- Source sets
    - 기본적으로 commonMain, commonTest, <targetName><compliationName>가 선언됨.

**compilerOptions**

- 모든 대상과 공유 소스 세트에 대한 기본값으로 사용되는 일반 확장 수준 컴파일러 옵션 지정.

### 공통 대상 구성

**platformType**

- 타겟을 위한 코틀린 플랫폼.
- jvm, androidJvm, js, wasm, native, common이 있음.

**artifactsTaskName**

- 타겟의 결과 아티팩트를 빌드하는 Task 이름.

components

- Gradle publications를 설정하기 위해 사용되는 컴포넌트들.

compilerOptions

- 타겟을 위한 컴파일러 옵션.