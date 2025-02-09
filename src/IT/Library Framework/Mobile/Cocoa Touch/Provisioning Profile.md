# Provisioning Profile

<aside>
💡

Apple의 정책에 따라 iOS 앱을 특정 기기에서 실행할 수 있도록 하는 인증 파일.

유효 기간이 존재.

</aside>

### 내용

- **App ID**
- **Certificates**
- **Devices**
- **Entitlements**

### 종류

**Development Profile**

- 개발용.
- XCode에서 실행할 때 필요.
- 등록된 테스트 기기에서만 실행 가능.

**Ad Hoc Profile**

- 외부 테스트용.
- App Store에 올리지 않고 등록된 테스트 기기에서만 실행 가능.

**App Store Profile**

- 정식 배포용.
- App Store에 올릴 때 사용되며, 기기 제한 없이 실행 가능.

**Enterprise Profile**

- 기업 내부용.
- 기업 내부에 배포할 때 사용되며, 기기 제한 없이 실행 가능.