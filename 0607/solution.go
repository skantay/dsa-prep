package leetcode

select sp.name
from SalesPerson as sp
where sp.sales_id not in (
    select distinct o.sales_id
    from Orders as o
    inner join Company as c
    on o.com_id = c.com_id and c.name = 'RED'
)
