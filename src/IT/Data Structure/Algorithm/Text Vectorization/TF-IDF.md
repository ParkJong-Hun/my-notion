# TF-IDF

<aside>
💡

Term Frequency-Inverse Document Frequency.
단어가 문서 내에서 얼마나 중요한지를 나타내는 가중치.

</aside>

**TF**

- Term Frequency.
- 단어 빈도.
- 특정 문서 내에서 단어가 얼마나 자주 등장하는지 나타냄.
- TF(t, d) = 단어 t가 문서 d에 등장한 횟수 / 문서 d의 전체 단어 수.

**IDF**

- Inverse Document Frequency.
- 역문서 빈도.
- 전체 문서에서 특정 단어가 얼마나 중요한지 평가.
- IDF(t) = log (전체 문서 수 / 단어 t가 등장한 문서 수)