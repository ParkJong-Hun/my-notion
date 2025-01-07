# Kotlin Native Command

### Gradlew

./gradlew nativeBinaries

- Executable 폴더를 생성.
    - 내부에 kexe나 exe 파일 존재.

### Kotlinc-native

kotlinc-native

- -o {kexe or exe name}
    - kexe의 이름을 지정.
- -p {output}
    - output 파일을 지정.
        - program
            - 기본값.
            - kexe, exe 등을 생성.
        - static
            - 정적 라이브러리(.a, .lib) 생성.
        - dynamic
            - 공유 라이브러리(.so, .dll) 생성.
        - framework
            - iOS용 프레임워크 생성.
        - library
            - klib + 공유 라이브러리 생성.
        - bitcode
            - LVVM Bitcode 생성.
- -r {path}
    - 저장소 경로 지정.
    - 기본적으로는 ~/.konan
- -l {path}
    - 링크할 klib 지정.