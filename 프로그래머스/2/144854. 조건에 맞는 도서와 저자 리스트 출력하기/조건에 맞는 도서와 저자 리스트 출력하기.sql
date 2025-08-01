-- 코드를 입력하세요
SELECT B.BOOK_ID as book_id,
       A.AUTHOR_NAME as AUTHOR_NAME,
       DATE_FORMAT(B.PUBLISHED_DATE,'%Y-%m-%d') AS PUBLISHED_DATE
from BOOK B
LEFT JOIN AUTHOR A ON B.AUTHOR_ID = A.AUTHOR_ID
WHERE B.CATEGORY = '경제'
ORDER BY B.PUBLISHED_DATE ASC;