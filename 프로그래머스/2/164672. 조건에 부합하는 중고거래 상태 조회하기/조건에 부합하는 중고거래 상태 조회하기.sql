-- 코드를 입력하세요
SELECT 
board_id, writer_id, title, price, 
replace(replace(replace(status,'SALE', '판매중'),'RESERVED', '예약중'),'DONE', '거래완료') as status

from used_goods_board 

where date_format(created_date, '%Y%m%d')= 20221005

order by board_id desc;