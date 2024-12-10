# config.yml

<aside>
💡 자동화하기를 원하는 CI 처리들을 작성.

</aside>

## **version**

<aside>
💡 필수

</aside>

CI의 버전.

## **orbs**

재사용 가능한 구성 요소 패키지화 도구.

작업 단계, 실행 환경, 환경 변수, 명령어 등의 구성을 포함 가능.

(orbs는 구성 요소의 재사용성과 표준화에 중점을 두고, executors는 작업의 실행 환경과 리소스 관리에 중점을 둠)

## executors

개발자가 정의한 작업을 실행시킬 환경.

함께 실행되어야 하는 프로그램에 대한 이미지를 정의.

## machines

작업을 실행하는 데 사용되는 가산 머신 환경.

커스텀 가상 머신 이미지를 사용해 작업을 실행하고자 할 때 사용.

이를 통해 작업에 더 많은 제어권을 갖고, 필요한 시스템 패키지, 라이브러리, 도구 등을 설치하고 구성 가능.

**image**

사용할 가상 머신 이미지.

**environment**

가상 머신 환경에 대한 환경 변수 정의.

**save_cache**

캐시 저장. cache_key를 이용해 어떠한 디렉토리의 데이터 등을 저장해 다른 job에서 사용 가능하게 할 수 있음. 

**restore_cache**

save_cache로 저장한 데이터를 복원.

## **cache_keys**

캐시를 구성하고 관리하기 위해 사용되는 설정 옵션.

캐시에 고유의 식별자를 정의해, 어떤 조건에서 사용되는지 지정 가능.

**cache**

캐시를 설정할 때 사용되는 전체 설정 항목 이름.

- **key**
    - 식별하는 고유 값.
- **keys**
    - 위에서부터 순차적으로 캐시가 있는지 찾아서 사용하는 key.
- **paths**
    - 캐시할 디렉토리나 파일 경로 지정.

## **commands**

jobs의 steps에서 사용할 run의 예약어.

여러 곳에서 사용할 run의 내용을 변수화.

## jobs

<aside>
💡 필수

</aside>

자동화하고 싶은 작업의 정의.

정의한 것일 뿐 모든 job들은 어떤 방법으로 실행시킬지에 대해서는 workflow 항목에서 작성.

**executor**

executors에 정의해둔 개발 환경 중에서 해당 job 가동에 필요한 환경.

**steps**

job에서 실행할 내용들.

**checkout**

자동으로 저장소를 감지하고 해당 저장소의 최신 코드를 가져와 job 환경에 배치.

주로 맨 처음에 작성.

다른 job에서 이전 job의 결과물을 사용할 때 필요.

**restore_cache**

캐시를 복원.

job이 실행되기 전에, 다른 job에 save_cache를 통해 저장된 것이 있다면 이를 사용하기 위해 복원.

**save_cache**

캐시를 저장.

저장한 데이터는 다른 job 들에서 사용 가능.

**run**

명령을 실행.

## **workflows**

<aside>
💡 필수

</aside>

앞서 작성한 job들을 어떠한 조합으로 어떤 조건에에서 어떤 순서로 작동시킬건지에 대한 정의.

**context**

보안 정보와 인증 자격 증명을 저장하고 사용하는 데 사용되는 설정 요소.

**requires**

현재 job을 수행하기 이전에 선행되어야 할 job 목록.

**filter**

현재 job이 실행되기 위한 조건(예를 들어 branch).

- example

```
version: 2.1
jobs:
  build:
    docker:
      - image: circleci/python:3.8

    steps:
      - checkout

      - run:
          name: Install dependencies
          command: pip install -r requirements.txt

      - run:
          name: Run tests
          command: pytest

workflows:
  version: 2
  build_workflow:
    jobs:
      - build
```

**context**

여러 환경 변수를 그룹화 해 중앙 집중화 하는데 사용.