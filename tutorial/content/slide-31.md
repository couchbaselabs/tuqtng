## Currencies being used 

How many different currencies are being used and the number of users using each of the currencies 

<pre id="example">
    SELECT 
        COUNT(product_details.currency) num_users, 
        product_details.currency 
            FROM  users_with_orders 
                WHERE product_details.currency IS NOT NULL 
                    GROUP BY product_details.currency
</pre>
