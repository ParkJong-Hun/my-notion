# SurfaceView vs TextureView

<aside>
💡 SurfaceView가 전력 효율이 더 높음.
프레임 타이밍이 더 정확해 동영상 재생이 원활.
지원 기기에서 고화질 HDR 동영상 출력 지원.
DRM으로 보호된 콘텐츠 재생 시 보안 출력 지원.
TextureView를 사용하면 일부 기기에서 동영상 재생 중 총 전력 소모량이 최대 30%까지 증가.
Android 7.0 이전에서는 부드러운 애니메이션, 동영상 노출 영역 스크롤을 SurfaceView가 지원하지 않기 때문에 사용.(현재로서 Android 24 이하를 지원하는 앱은 거의 없음.)

</aside>