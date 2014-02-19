## Fraudulent Accounts?

We've decided that users with a lot more orders than searches are fraudulent.  Let's find them.

N1QL supports different types of LENGTH functions. 

1. use LENTGH() when the argument is a string
2. use ARRAY_LENGTH() when the argument is an array
3. use OBJECT_LENGTH() when the argument is a map
4. use POLY_LENGTH() when the type of argument is not known

In this query we are computing the number of searches being carried out by a user and the number of orders placed. If the number of orders placed is 8 times the number of searches then that user is probably doing something fishy. 

<pre id="example">
SELECT 
  personal_details.display_name, 
  POLY_LENGTH(shipped_order_history) AS num_orders, 
  ARRAY_LENGTH(search_history) AS num_searches 
	FROM users_with_orders
		WHERE POLY_LENGTH(shipped_order_history) > 8*ARRAY_LENGTH(search_history)
</pre>
