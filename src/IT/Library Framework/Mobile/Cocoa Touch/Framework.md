# Framework

<aside>
💡

XCode에서 Framework란, XCode에서 사용하기 위해 만든 라이브러리 패키지를 의미.

Bundle를 가지고 있어서 UI에서 사용하는 이미지 등을 보관할 수 있음.

</aside>

<aside>
⚠️

Framework는 하나의 CPU 아키텍처만 대응하기 때문에 JVM의 JAR 같이 활용할 수는 없음.

</aside>

### 라이브러리 링크 타입

**Dynamic**

- 프레임워크를 별도로 빌드한 이후, 앱이 실행되는 런타임에 프레임워크를 불러오는 방식.
- `.dylib`

**Static**

- 프레임워크를 import한 곳의 일부로서 같이 빌드되는 방식.
- `.a`

**Embed**

- 앱이 패키징될 때 결과물에 포함됨.

[XCFramework%20101f37315c4480fc8101ccfab12f6876](XCFramework%20101f37315c4480fc8101ccfab12f6876)

[Umbrella%20Framework%20139f37315c4480b49e45e4d88869d67a](Umbrella%20Framework%20139f37315c4480b49e45e4d88869d67a)