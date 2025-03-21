# Authentication, Authorization

<aside>
💡 인증과 인가

</aside>

인증(Authentication)과 인가(Authorization)는 정보보호와 관련된 개념으로 자주 사용됨. 

하지만 둘은 서로 다른 개념이며, 이해하지 못하면 보안에 취약할 수 있음.

**인증(Authentication)**

사용자가 자신이 주장하는 것이 맞는지 확인하는 과정. 예를 들면, 로그인 과정에서 사용자가 제공한 아이디와 비밀번호를 서버가 확인하고, 그 결과 사용자가 자신이 주장하는 것이 맞는지 아닌지를 판단.

**인가(Authorization)**

인증된 사용자가 특정 자원에 접근할 권한이 있는지 확인하는 과정.

예를 들면, 사용자가 특정 파일에 접근하려고 할 때, 그 사용자가 그 파일을 읽을 수 있는 권한이 있는지를 확인합니다.

따라서, 인증과 인가는 보안을 강화하기 위해 둘 다 필요. 

인증 없이 인가만 있다면, 누구나 자원에 접근할 수 있기 때문. 

반대로, 인가 없이 인증만 있다면, 인증된 사용자라 하더라도 미승인된 자원에 접근할 수 있기 때문.

[%E1%84%8B%E1%85%B5%E1%86%AB%E1%84%8C%E1%85%B3%E1%86%BC%E1%84%8B%E1%85%B4%20%E1%84%8C%E1%85%A9%E1%86%BC%E1%84%85%E1%85%B2%208b88eecefa4b4ac0a95ea007fff4b2f5](%E1%84%8B%E1%85%B5%E1%86%AB%E1%84%8C%E1%85%B3%E1%86%BC%E1%84%8B%E1%85%B4%20%E1%84%8C%E1%85%A9%E1%86%BC%E1%84%85%E1%85%B2%208b88eecefa4b4ac0a95ea007fff4b2f5)

[OAuth%202%200%205d4df9fa04ee4fa086c71926db847c27](OAuth%202%200%205d4df9fa04ee4fa086c71926db847c27)

[OAuth%202%200%208aa02feaa5364cfb9e465b0c0a161e9e](OAuth%202%200%208aa02feaa5364cfb9e465b0c0a161e9e)

[OAuth%201e3c7b85ddcd4a7eb2ab4e0759459a31](OAuth%201e3c7b85ddcd4a7eb2ab4e0759459a31)

[Basic%20Authorization%2016ff37315c44805dbfc2e86b48fe0aa0](Basic%20Authorization%2016ff37315c44805dbfc2e86b48fe0aa0)