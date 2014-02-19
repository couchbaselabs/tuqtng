## When do people order?

What's the most popular day to place an order?

<pre id="example">
SELECT 
  SUBSTR(shipped_order.order_datetime,0,3) AS day,
  COUNT(*) AS count
    FROM users_with_orders AS profiles
        UNNEST profiles.shipped_order_history AS shipped_order
            GROUP BY SUBSTR(shipped_order.order_datetime,0,3)
                ORDER BY count DESC
</pre>
