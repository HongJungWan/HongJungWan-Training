[https://school.programmers.co.kr/learn/courses/30/lessons/301650], 특정 세대의 대장균 찾기

WITH RECURSIVE Generation AS (
    SELECT
        ID,
        PARENT_ID,
        1 AS gen
    FROM ECOLI_DATA
    WHERE PARENT_ID IS NULL

    UNION ALL

    SELECT
        e.ID,
        e.PARENT_ID,
        g.gen + 1
    FROM ECOLI_DATA e
    INNER JOIN Generation g ON e.PARENT_ID = g.ID
    WHERE g.gen < 3  -- 재귀 깊이 제한 추가
)
SELECT ID
FROM Generation
WHERE gen = 3
ORDER BY ID;
