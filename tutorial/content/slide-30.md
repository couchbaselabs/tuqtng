## Users' Orders

How many orders do our users have?

<pre id="example">
SELECT 
  LENGTH(shipped_order_history) AS num_orders, 
  COUNT(*) AS num_profiles 
	FROM profiles 
		WHERE doc_type = "user_profile" 
			GROUP BY LENGTH(shipped_order_history)
				ORDER BY num_orders
</pre>