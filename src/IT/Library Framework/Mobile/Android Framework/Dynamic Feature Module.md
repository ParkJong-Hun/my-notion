# Dynamic Feature Module

### Android Manifest

**<dist>**

- Google Play의 Dynamic Delivery 기능 설정.
- **<dist:module>**
    - 모듈 이름 설정.
    - **dist:instant**
        - 앱을 설치 없이 즉시 실행할 수 있는 형태로 제공.
    - **dist:title**
        - 앱 번들을 배포할 때 사용되는 앱 대체 제목 설정 태그.
        - 앱 현지화에 사용.
- **<dist:delivery>**
    - 모듈이 앱에 제공되는 방식.
    - dist:install-time
        - 앱 설치 시 자동 설치.
    - dist:on-demand
        - 사용자가 필요할 때만 설치.
- **<dist:base-feature>**
    - 앱의 기본이 되는 모듈.
- **<dist:fusing>**
    - 여러 APK를 하나의 앱 번들로 병합해 필요한 리소스만 포함되도록 함.