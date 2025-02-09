# App Release Process

### Android

1. 앱 서명 준비
    1. keystore 파일 생성
    2. singingConfig에 등록
2. 앱 빌드
    1. 릴리스 빌드
    2. .apk, .aab 파일 생성
3. Google Play Console에 앱 업로드
    1. Google Play Console에 계정 등록
    2. 앱 패키지 등록
    3. .aab 혹은 .apk 파일 업로드
    4. 앱 심사 
4. 릴리스됨
    1. .apk 파일 다운로드 가능

### iOS

1. 앱 서명 준비
    1. App Developer 계정 로그인
    2. 계정에서 Provisioning Profile과 Certificate 생성
    3. Signing & Capabilities 설정
2. 앱 빌드
    1. 릴리스 빌드 시, App Store Connect 선택 후 앱 서명
    2. .ipa 파일 생성
3. App Store Console에 앱 업로드
    1. App Store Connect에 계정 등록
    2. 앱 등록
    3. 앱 심사
4. 릴리스됨
    1. .app 파일 다운로드 가능

<aside>
⚠️

생성한 Provisioning Profile이나 Certificate는 로컬에 저장할 수 있지만, Apple 서버에서 계정과의 검증을 거쳐야하는 경우가 많기 때문에 결국 필요할 수 있음.

</aside>