# Bundler

[Ruby%209d04253dc0dd4c1984886719f85cb506](Ruby%209d04253dc0dd4c1984886719f85cb506)

<aside>
💡 Ruby 언어에서 사용되는 패키지 의존성 관리 도구.

</aside>

프로젝트에 필요한 외부 라이브러리인 Gem 설치와 관리를 용이하게 할 수 있음.

## 명령어

**bundle install**

Gemfile에 명시된 종속성 설치. 명령어를 실행하기 전에 Gemfile이 있는 디렉토리에 이동해야 함.

**bundle update**

Gemfile에 명시된 종속성을 최신 버전으로 업데이트.

**bundle exec**

Bundle를 통해 설치된 종속성으로 프로젝트의 Ruby 스크립트나 명령어를 실행함.

이 명령어를 사용하면 Bundler가 정확한 버전 종속성을 사용하도록 보장.

**bundle show**

현재 프로젝트에 사용 중인 종속성의 목록을 보여줌.

**bundle add**

Gemfile에 종속성을 추가.

예를 들어, bundle add <gem_name>과 같이 사용.

**bundle remove**

Gemfile에서 종속성 제거.

예를 들어 bundle remove <gem_name>과 같이 사용.

**bundle exec rspec**

RSpec과 같은 테스트 프레임워크를 실행.

프로젝트에서 테스트를 작성하고 실행할 때 자주 사용.

**bundle console**

프로젝트의 Ruby 콘솔을 실행.

이 명령어를 사용하면 프로젝트의 환경과 종속성이 로드된 Ruby REPL을 사용 간능.

**bundle —help**

Bundler 전체 명령어 목록 확인.