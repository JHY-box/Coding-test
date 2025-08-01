-- 코드를 입력하세요
SELECT HISTORY_ID, CAR_ID, DATE_FORMAT(START_DATE,'%Y-%m-%d') AS START_DATE, 
DATE_FORMAT(END_DATE,'%Y-%m-%d') AS END_DATE,
 IF(DATEDIFF(END_DATE,START_DATE)+1 >= 30, '장기 대여', '단기 대여') AS RENT_TYPE

FROM CAR_RENTAL_COMPANY_RENTAL_HISTORY

WHERE DATE_FORMAT(START_DATE, '%Y%m')= '202209'

ORDER BY HISTORY_ID DESC;
                                                          
                                                          -- DATE_ADD(date,INTERVALnunit)
                                                          
                   --    +1 일을 해줘야함                -- date_add(start_date,interval 30 day)
                                                         
                                                         -- DATE_FORMAT(START_DATE, '%Y%m')= 202209