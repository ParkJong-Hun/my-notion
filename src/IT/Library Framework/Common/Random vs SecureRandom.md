# Random vs SecureRandom

<aside>
💡

Java의 난수 생성을 위한 두 클래스.

</aside>

### SecureRandom

- [java](java)
- LCG를 사용.
- Random의 하위 클래스.
- 복잡한 알고리즘으로 예측이 어려움.
- 최대 128비트.

### Random

- java.util
- CSPRNG를 사용.
- 시스템 시간을 사용한 알고리즘.
- 최대 48비트.