# google-services.json

<aside>
💡 Firebase, GCP의 설정 파일.
Android 앱이 Firebase나 GCP와 통신하기 위해 필요한 정보가 포함되어 있음.
Firebase 프로젝트를 생성하고 Android 앱을 프로젝트에 연결할 때 자동으로 생성됨.

</aside>

- 프로젝트 ID
- 앱 ID
- API 키

- 스토리지 버킷
    - Google Cloud Storage 버킷의 이름.
- 패키지 이름
    - Android 앱의 고유한 식별자.
    - Google play 스토어에 게시하고 안드로이드 단말에서 설치된 앱을 식별하는 데 사용.
    - 일반적으로 도메인 이름의 반대 형식.
        - ex) [your-project-id.appspot.com](http://your-project-id.appspot.com) → com.yourdomain.yourapp

- 번들 ID
    - iOS 앱의 고유한 식별자.
    - 개발 및 배포 과정에서 필요.
- 앱 스토어 ID
    - App Store에 게시하고 iOS 단말에서 설치된 앱을 식별하는 데 사용.