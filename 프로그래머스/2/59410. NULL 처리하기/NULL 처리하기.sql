-- 코드를 입력하세요
SELECT animal_type, ifnull(NAME, 'No name') AS NAME, SEX_UPON_INTAKE from animal_ins order by animal_id asc;