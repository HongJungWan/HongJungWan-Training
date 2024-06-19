[https://school.programmers.co.kr/learn/courses/30/lessons/276035], FrontEnd 개발자 찾기

SELECT
    ID,
    EMAIL,
    FIRST_NAME,
    LAST_NAME
FROM DEVELOPERS D, SKILLCODES S
WHERE 1 = 1
  AND (D.SKILL_CODE & S.CODE) > 0
  AND CATEGORY = 'Front End'
GROUP BY ID, EMAIL, FIRST_NAME, LAST_NAME
ORDER BY ID;
