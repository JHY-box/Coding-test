-- 코드를 작성해주세요
select count(fish_type) as fish_count, max(length) as max_length, fish_type
from fish_info
group by fish_type
having avg(ifnull(length,10)) >= 33
order by fish_type