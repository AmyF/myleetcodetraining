# Write your MySQL query statement below
select c.name Customers from Customers c where c.id not in (select o.CustomerId from Orders o)