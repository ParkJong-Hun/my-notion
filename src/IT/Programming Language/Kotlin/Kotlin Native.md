# Kotlin Native

<aside>
💡 JVM 바이트 코드로 변환되지 않는 Kotlin.
네이티브 플랫폼에서 실행될 수 있는 Kotlin 코드. 
주로 iOS, macOS, Linux 등과 같은 네이티브 플랫폼에서 Kotlin을 사용하고자 할 때 유용. 
네이티브 환경에서 최상의 성능을 제공.
LLVM 프레임워크와 함께 작동하여 네이티브 코드로 컴파일.

</aside>

**Kotlin Native Interop**

- Kotlin/Native의 핵심.
- C, C++, Swift 등과 상호작용.
    - **cinterop**
        - C와 Kotlin 간의 상호작용을 위한 바인딩을 생성하고 헤더파일을 읽어서 Kotlin 코드에서 사용할 수 있도록 변환.

**platform.posix**

- POSIX 환경에서 제공되는 기능(파일 시스템, 네트워크, 프로세스 관리)을 Kotlin에서 사용할 수 있게 해주는 모듈.