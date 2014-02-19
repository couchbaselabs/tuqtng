## Monthly revenues 

How do our monthly revenues stack up 

Excercises

1. Change the query to only display the top 3 months
2. Change the query to display only the worst 3 months
3. Change the query to display revnues ordered by year instead 

<pre id="example">
    SELECT 
        count(*) AS num_orders, 
        SUBSTR(order_details.order_datetime, 4, 3) AS month, 
        SUM(payment_details.total_charges) as revenue 
            FROM orders_with_users 
                WHERE doc_type="order" 
                    GROUP BY SUBSTR(order_details.order_datetime, 4, 3) 
                        ORDER BY revenue DESC
</pre>
