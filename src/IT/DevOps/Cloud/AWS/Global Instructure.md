# Global Instructure

<aside>
💡

AWS가 제공하는 클라우드 서비스를 가동하기 위한 물리적 설비.

</aside>

**Region**

- AWS가 서비스를 제공하고 있는 거점(국가, 지역)
- 각 Region은 서로 물리적으로 떨어져 있음
- 여러 AZ로 이루어져 있음
- **AWS Local Zone**
    - 특정 Region의 확장

**AZ**

- Availbility Zone
- 여러 개의 데이터 센터로 구성됨
- **Multi AZ**
    - 내장애성을 높이기 위해 여러 AZ를 이용해 시스템을 구축하는 것

**Edge Location**

- 물리적 데이터 센터
- 저레이턴시로 고속으로 컨텐츠를 송신 가능하게 함
- AWS 네트워크에 접속되어 리퀘스트가 발생할 때마다 가장 가까운 Edge Location에서 처리됨