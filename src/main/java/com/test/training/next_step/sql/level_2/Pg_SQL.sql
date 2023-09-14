//@formatter:off
//ORACLE

-- level 2

-- 동물 수 구하기
SELECT NAME, COUNT(NAME) as COUNT
FROM ANIMAL_INS
GROUP BY NAME
HAVING COUNT(NAME) >= 2
ORDER BY NAME;


-- NULL 처리하기
SELECT ANIMAL_TYPE,
       CASE
           WHEN NAME IS NULL THEN 'No name'
           WHEN NAME IS NOT NULL THEN NAME
       END as NAME,
       SEX_UPON_INTAKE
FROM ANIMAL_INS
ORDER BY ANIMAL_ID


-- 재구매가 일어난 상품과 회원 리스트 구하기
SELECT USER_ID, PRODUCT_ID
FROM ONLINE_SALE
GROUP BY USER_ID, PRODUCT_ID
HAVING COUNT(SALES_AMOUNT) >= 2
ORDER BY USER_ID ASC, PRODUCT_ID DESC;


-- 이름에 el이 들어가는 동물 찾기
SELECT ANIMAL_ID,NAME
FROM ANIMAL_INS
WHERE LOWER(NAME) LIKE '%el%' AND ANIMAL_TYPE LIKE 'Dog'
ORDER BY NAME;

//@formatter:off