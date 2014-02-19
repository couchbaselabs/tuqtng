## Top ten customers

Here we use the JOIN clause to get our top ten customers

In this query we are first join the list of orders that placed by each customer and iterate over each order value to generate the total amount spent and group the results y th customers display name. The limit clause is used to only display the top ten spenders. 

<pre id="example">
    SELECT 
        user.personal_details.display_name, 
        SUM(orders.payment_details.total_charges) as spent 
            FROM users_with_orders user 
                JOIN orders_with_users orders 
                    KEYS ARRAY s.order_id FOR s IN user.shipped_order_history END 
                        GROUP BY user.personal_details.display_name 
                            ORDER BY spent DESC 
                                LIMIT 10
</pre>
