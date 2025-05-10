# IAM

<aside>
💡

Identity and Access Management.

AWS 리소스에 대한 접근을 안전하게 제어하는 서비스.

</aside>

**IAM User**

- 개별 AWS 계정. 특정 권한을 가진 사람이나 앱을 나타냄.
- 장기 보안 인증 정보를 사용.

**IAM Group**

- IAM User의 집합.
- 그룹에 권한을 할당해 그룹 모든 사용자가 해당 권한 상속 가능.

**IAM Role**

- 특정 권한을 가진 객체.
- User나 AWS 서비스가 일시적으로 맡을 수 있음.
- 단기 보안 인증 정보를 제공해 보안 강화.

**IAM Policy**

- JSON 형식 문서로 여러 권한을 할당하는 데 사용.
- AWS 관리형 정책과 고객 관리형 정책으로 구분.