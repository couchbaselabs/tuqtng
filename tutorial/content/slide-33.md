## When do people order?

What's the most popular day to place an order?

<pre id="example">
SELECT 
  SUBSTR(shipped_order.order_datetime,0,3) AS day,
  COUNT(*) AS count
    FROM profiles AS profile
        OVER shipped_order IN profile.shipped_order_history
            GROUP BY SUBSTR(shipped_order.order_datetime,0,3)
                ORDER BY count DESC
</pre>