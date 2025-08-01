-- 코드를 작성해주세요
SELECT COUNT(*) AS FISH_COUNT,
       B.FISH_NAME AS FISH_NAME
FROM FISH_INFO A
INNER JOIN FISH_NAME_INFO B ON A.FISH_TYPE = B.FISH_TYPE
GROUP BY B.FISH_NAME                          -- GROUP BY B.FISH_TYPE  틀렸던 부분 뭘써도 상관없는데 select 에서 fish_name 떄문에 여기엔 네임을 써야함
ORDER BY FISH_COUNT DESC;