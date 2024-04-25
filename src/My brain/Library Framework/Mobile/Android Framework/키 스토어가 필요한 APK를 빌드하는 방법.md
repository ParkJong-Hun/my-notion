# 키 스토어가 필요한 APK를 빌드하는 방법

1. 키스토어 파일을 로컬의 프로젝트의 keystore 디렉토리에 추가
2. settings.gradle.kts에서 keystore를 참조하고 있는 곳에서 추가한 키스토어 파일을 지정
3. Build APK(s)를 클릭해서 빌드