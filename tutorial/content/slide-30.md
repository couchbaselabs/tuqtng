##  Order prices

What is the min, max and avg price for the orders grouped by payment mode

<pre id="example">
    SELECT 
        payment_details.payment_mode, 
        MIN(product_details.sale_price) AS min_price, 
        AVG(product_details.sale_price) AS avg_price, 
        MAX(product_details.sale_price) AS max_price
            FROM orders_with_users 
                WHERE product_details IS NOT NULL
                GROUP BY payment_details.payment_mode
</pre>

