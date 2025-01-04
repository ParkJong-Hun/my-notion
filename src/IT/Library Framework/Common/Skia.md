# Skia

<aside>
💡 구글에서 개발하고 있는 C++ 크로스 플랫폼 2D 렌더링 라이브러리.
HTML의 Canvas API 같은 형태로 화면 상에 좌표 베이스로 화면 요소를 렌더링.
Flutter의 많은 라이브러리에서 사용됨.

</aside>

### 객체

**SkSurface**

- 렌더링 작업 대상.
- SkCanvas를 생성.
- CPU나 GPU 백엔드에 연결.

**SkCanvas**

- 모든 드로잉 작업이 이루어지는 표면.
- 그림을 그리는 대상(화면, 비트맵 등)에 대한 IF 제공.
- 그릴 수 있는 메서드 제공(`drawOOO`).

**SkPaint**

- 드로잉 스타일과 속성.
- 색상, 선, 두께, 투명도, 그라데이션, 쉐이더 설정.
- `SkCanvas`와 사용되어 드로잉 모양 결정.

**SkPath**

- 복잡한 도형.
- 선, 곡선, 닫힌 영역 등을 정의.
- 여러 도형을 결합해 변형.

**SkTypeface**

- 폰트 정보

**SkBitmap**

- 비트맵 데이터.
- 픽셀 데이터 직접 관리.
- 이미지를 가져오거나 내보낼 때도 사용.

**SkImage**

- 읽기 전용 이미지.
- `SkBitmap`과 다르게 수정 불가능.

**SkShader**

- 그라데이션이나 패턴.
- `SkPaint`와 결합해 렌더링 적용.

**SkMatrix**

- 좌표 변환(이동, 회전, 스케일링, 기울기).
- `SkCanvas`와 사용되어 드로잉 결과 변형.

[Skija%20157f37315c4480feac0bffb9aaa16707](Skija%20157f37315c4480feac0bffb9aaa16707)

[Skiko%2023a84f5beb954c09bed7f04f62d263e8](Skiko%2023a84f5beb954c09bed7f04f62d263e8)

[CanvasKit%20157f37315c448047ac6cee6eb023db40](CanvasKit%20157f37315c448047ac6cee6eb023db40)

[SkiaSharp%20157f37315c44803da112db169f6237b7](SkiaSharp%20157f37315c44803da112db169f6237b7)

[SkSL%20170f37315c4480888d48edc7e0b9ea98](SkSL%20170f37315c4480888d48edc7e0b9ea98)