# Git

**git reset {COMMIT} {path/file}**

해당 파일을 해당 커밋으로 되돌림.

**git reflog**

깃에서 입력한 커맨드를 볼 수 있음. (reset 포함)

[Git%20%E1%84%8B%E1%85%A5%E1%86%AB%E1%84%8B%E1%85%A5%E1%84%80%E1%85%A1%20%E1%84%8C%E1%85%AE%E1%86%BC%E1%84%80%E1%85%AE%E1%86%A8%E1%84%8B%E1%85%A5%E1%84%8B%E1%85%B5%E1%86%AF%20%E1%84%84%E1%85%A2%20e04c6833e1e842eab7e0393047559a97](Git%20%E1%84%8B%E1%85%A5%E1%86%AB%E1%84%8B%E1%85%A5%E1%84%80%E1%85%A1%20%E1%84%8C%E1%85%AE%E1%86%BC%E1%84%80%E1%85%AE%E1%86%A8%E1%84%8B%E1%85%A5%E1%84%8B%E1%85%B5%E1%86%AF%20%E1%84%84%E1%85%A2%20e04c6833e1e842eab7e0393047559a97)

**git checkout -b {local-branch-name} origin/{remote-branch-name}**

트래킹 브랜치 작성

**git branch --set-upstream-to=origin/{remote-branch-name} {local-branch-name}**

트래킹 브랜치 설정

**git push -u origin {local-branch-name}**

트래킹 브랜치 설정 + 푸쉬

**git config --global push.default**

- simple
    - 현재 로컬 브랜치와 이름이 같은 리모트 브랜치가 트래킹 브랜치로 설정되어 있을 때만 push
    - 안전한 기본값으로, 실수로 다른 브랜치에 push하는 것을 방지.
    - Git 2.0+ 기본값
- current
    - 현재 로컬 브랜치와 이름이 같은 리모트 브랜치로 push.
    - 트래킹 브랜치 설정 여부와 관계없이 동일한 이름의 리모트 브랜치로 push를 시도.
    - 이전 Git 기본값
- upstream
    - 현재 로컬 브랜치가 트래킹하는 리모트 브랜치로 push.
- matching
    - 로컬 저장소에 있는 모든 브랜치 중에서 이름이 같은 리모트 브랜치가 있으면 모두 push.
    - 예상치 못한 변경을 일으킬 수 있어 현재는 권장되지 않음.
    - Git 2.0 이전 기본값