-- 코드를 작성해주세요
select year(YM) as year, round(avg(pm_val1),2) as 'pm10', round(avg(pm_val2),2) as 'pm2.5'
from AIR_POLLUTION
where location2 = "수원" 
group by year
order by year asc;