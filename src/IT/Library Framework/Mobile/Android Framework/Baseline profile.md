# Baseline profile

<aside>
💡 코드 경로의 해석과 JIT(Just-In-Time) 컴파일 단계를 피해 최초 실행 후 코드 실행 속도를 약 30% 개선.
ART이 AOT(ahead-of-time) 컴파일을 통해 포함된 코드 경로를 최적화해 모든 신규 앱 설치, 모든 앱 업데이트에 향상된 성능 제공.

</aside>

**Macrobenchmark**

UX, 시작, 애니메이션 같은 상호작용을 측정하기 위해 앱용 Baseline profile을 생성하기 위해 사용.

API 23 이상.

**Microbenchmark**

보다 세분화된 앱별 상황에 관한 성능 분석을 위해 앱용 Baseline profile을 생성하기 위해 사용.

API 14 이상.