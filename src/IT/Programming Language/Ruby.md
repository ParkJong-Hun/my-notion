# Ruby

<aside>
💡 스크립트 언어의 일종.
일본의 마츠모토 유키히로가 발표.

</aside>

**do | {이름} |**

**end**

- 블록을 정의하고 그 안에서 인수(예를 들어 for 문의 각 스텝의변수)를 전달.

**yield**

- do | {이름} |으로 사용될 무언가를 전달할 때 사용.

**#{변수이름}**

- 문자열 보간

**{키}⇒{값}**

- 해시 키-값을 작성할 때 사용.

**begin**

**{처리}**

**rescue {에러} ⇒ {에러변수이름}**

**{에러가 발생할 때 처리}**

**ensure**

**{항상 실행되는 것}**

**end**

- try catch 같은 것.

**attr_accessor**

- 객체의 속성 정의
- getter, setter 자동 생성.

**attr_reader**

- 객체의 속성 정의
- getter만 자동 생성.

**attr_writer**

- 객체의 속성 정의
- setter만 자동 생성.

**module**

- 패키지 같은 단위로 관리.