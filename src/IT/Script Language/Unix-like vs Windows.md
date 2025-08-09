# Unix-like vs Windows

| function | Unix / Linux / macOS | Windows (CMD) | Windows (PowerShell) | macOS(only) |
| --- | --- | --- | --- | --- |
| 파일 / 폴더 목록 | `ls` | `dir` | `Get-ChildItem` or `ls` or `dir` |  |
| 폴더 이동 | `cd` | `cd` | `Set-Location` or `cd` |  |
| 파일 복사 | `cp` | `copy` | `Copy-Item` or `cp` or `copy` |  |
| 파일 이동 / 이름 변경 | `mv` | `move`or`rename` | `Move-Item` or `mv` |  |
| 파일 삭제 | `rm` | `del` or `erase` | `Remove-Item` or `rm` |  |
| 폴더 생성 | `mkdir` | `mkdir` or `md` | `New-Item -ItemType Directory` or `mkdir` |  |
| 화면 지우기 | `clear` | `cls` | `clear-Host` or `clear` or `cls` |  |
| 파일 내용 출력 | `cat` | `type` | `Get-Content` or `cat` or `type` |  |
| 파일 경로 | `/`  | `\`  | `\` |  |
| 명령어 옵션 | `-`  | `/`  | `/` |  |
| 파이프(`|`) | 텍스트 스트림 전달 | 객체 전달 | 객체 전달 |  |
| 클립보드 복사 | `xclip -selection clipboard` |  | `Set-Clipboard` | `pbcopy` |
| 클립보드 붙여넣기 | `xclip -selection clipboard -o` |  | `Get-Clipboard`  | `pbpaste` |
| 검색 | `grep`  | `findstr` | `Select-String` or `sls` |  |
| 치환 | `sed`  |  | `Get-Content -replace`  |  |
| 열처리 | `awk`  |  | `Import-Csv` … |  |
| 파일 검색 | `find`  |  | `Get-ChildItem -Recurse` or `gci -r` |  |
| 프로세스 모니터링 | `htop`  | `tasklist` | `Get-Process` |  |
| 열린 파일 / 포트 확인 | `lsof`  |  | `Get-Process -IncludeUserName | where {$_.Id -eq PID}`or `netstat -aon` |  |
| 터미널 다중화 | `tmux` |  |  |  |
| 현재 작업 중 폴더의 전체 경로 출력 | `pwd`  | `cd` or `echo %cd%`  | `Get-Location`  |  |

<aside>
💡

Unix에서는 주로 sh, bash 등을 사용.
Linux에서는 주로 bash 사용.
macOS에서는 Unix(BSD) 기반의 zsh를 주로 사용.
Windows에서는 DOS/NT 기반의 CMD, PowerShell을 주로 사용.

</aside>