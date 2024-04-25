# CocoaPods 커맨드

[맥(Mac) - 코코아팟 제거하고 재 설치하기](https://mansu.tistory.com/23)

## Pods 제거

```
sudo gem install cocoapods-deintegrate cocoapods-clean
pod deintegrate
pod clean
rm Podfile
```

## Podfile 생성

```
pod init
```

## 설치(Podfile가 필요)

```
pod install
```

## 최신 **podspec 파일들로 업데이트**

```
pod repo update
```