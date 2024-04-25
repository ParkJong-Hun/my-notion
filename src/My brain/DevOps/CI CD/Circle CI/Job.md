# Job

<aside>
💡 CircleCI의 Job은 Step의 모음.
Step은 전부 한 단위로서 새 Container나 가상 머신 내에서 실행됨.
Job은 Workflow를 사용해 오케스트레이션됨.

</aside>

### Step

실행 가능한 커맨드의 모음으로 Job 내에서 실행됨.
코드를 체크아웃하기 위해 steps의 아래에 `checkout:`키가 필요.