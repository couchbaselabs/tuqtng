## Fraudulent Accounts?

We've decided that users with a lot more orders than searches are fraudulent.  Let's find them.

<pre id="example">
SELECT 
  personal_details.display_name, 
  LENGTH(shipped_order_history) AS num_orders, 
  LENGTH(search_history) AS num_searches 
	FROM profiles 
		WHERE LENGTH(shipped_order_history) > 8*LENGTH(search_history)
</pre>