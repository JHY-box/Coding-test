-- 코드를 입력하세요
SELECT animal_id, name, date_format(datetime, '%Y-%m-%d') as 날짜 from animal_ins order by animal_id asc;