-- 코드를 입력하세요
SELECT animal_id, name from animal_ins where animal_type = 'dog' and instr(name, 'el') > 0 order by name;