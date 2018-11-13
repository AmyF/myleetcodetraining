# Write your MySQL query statement below
select w1.id id from weather w1 join weather w2 on datediff(w1.RecordDate,w2.RecordDate) = 1 and w1.temperature > w2.temperature