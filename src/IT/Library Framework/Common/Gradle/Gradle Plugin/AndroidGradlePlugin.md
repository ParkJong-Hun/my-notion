# AndroidGradlePlugin

### 빌드 시 실행되는 작업

1. clean : 프로젝트 빌드 전 캐쉬 정리.
2. preBuild : 프로젝트 빌드 전 설정 처리.
3. processResources : 리소스 파일을 빌드 디렉토리로 복사.
4. generateBuildConfig : Android 빌드 구성 파일 생성.
5. processApplicationManifest : AndroidManifest.xml을 처리해 메타데이터를 포함.
6. generateRFile : res/를 처리해 [R](R) 혹은 R.kt 파일 생성.
7. linkResources : 리소스를 최종 APK에 포함.
8. stripDebugSymbols (release) : 디버그 심볼 제거.
9. minifyReleaseWithR8 (release) : 코드 최적화, 난독화.
10. compileJavaWithJavac : 네이티브 앱 코드 컴파일.
11. merge : 여러 리소스와 매니페스트 파일을 병합.
12. assemble : 최종 APK 빌드.
13. package : APK를 패키징해 설치 가능한 파일 생성.
14. buildOutputs : 빌드 후 APK, AAB 파일을 디렉터리에 작성.
15. signingReport : 앱 서명 정보 관련 보고서 생성.
16. install (디바이스 설치)