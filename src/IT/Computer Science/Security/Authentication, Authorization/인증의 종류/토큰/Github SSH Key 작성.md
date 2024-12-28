# Github SSH Key 작성

[generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent](generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

## 새 SSH 키 생성

```
$ ssh-keygen -t ed25519 -C "your_email@example.com"
```

## SSH-AGENT에 SSH 키 추가

1. 백그라운드에서 ssh-agent 시작.

```
$ eval "$(ssh-agent -s)"
> Agent pid 59566
```

1. macOS Sierra 10.12.2 이상을 사용하는 경우 ssh-agent에 키를 자동 로드하고 키 집합에 저장하도록 ~/.ssh/config을 수정해야함.
    - 먼저 `~/.ssh/config` 파일이 기본 위치에 있는지 확인
    
    ```
    $ open ~/.ssh/config
    > The file /Users/YOU/.ssh/config does not exist.
    ```
    
    - 파일이 없으면 파일을 생성
    
    ```
    $ touch ~/.ssh/config
    ```
    
    - `~/.ssh/config` 파일을 열고 다음 줄을 포함하도록 파일을 수정합니다. SSH 키 파일에 예제 코드와 다른 이름 또는 경로가 있는 경우 현재 설정과 일치하도록 파일 이름 또는 경로를 수정
    
    ```
    Host github.com
      AddKeysToAgent yes
      UseKeychain yes
      IdentityFile ~/.ssh/id_ed25519
    ```
    
    - Windows의 경우, `Start-Process -Verb runas powershell`로 관리자 모드 실행 이후 아래 커맨드 입력.
    
    ```jsx
    Get-Service ssh-agent | Set-Service -StartupType Automatic -PassThru | Start-Service -PassThru | Select-Object -Property Name, StartType, Status
    ```
    
2. ssh-agent에 SSH 프라이빗 키를 추가하고 키 집합에 암호를 저장. 다른 이름으로 키를 만들거나 이름이 다른 기존 키를 추가하는 경우 명령의 *id_ed25519* 를 프라이빗 키 파일의 이름으로 바꿈.

```
$ ssh-add --apple-use-keychain ~/.ssh/id_ed25519

윈도우즈의 경우 경로는 c:/Users/YOU/.ssh/id_ed25519
```

1. SSH 퍼블릭 키를 클립보드에 복사. 

```
$ pbcopy < ~/.ssh/id_ed25519.pub
  # Copies the contents of the id_ed25519.pub file to your clipboard
  
윈도우즈의 경우 cat c:/users/YOU/.ssh/id_ed25519.pub | clip
```

1. Github - Settings - 새 SSH 키 또는 SSH 키 추가에서 붙여넣기.