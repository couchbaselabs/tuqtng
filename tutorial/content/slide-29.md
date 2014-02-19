## Case Study I. E-commerce Application 

In this section we are going to demonstrate how N1QL can be used to solve a few typical e-commerce use-cases. We will use the following two buckets 

```
1. users_with_orders. This bucket contains a list of user profiles along with order Ids of the orders that they have placed. 
2. orders_with_users. This bucket contains the description of the orders
```

For starters lets look at what are the languages in use

<pre id="example">
SELECT 
  profile_details.prefs.ui_language AS language, 
  COUNT(*) AS count 
	FROM users_with_orders
		WHERE doc_type = "user_profile"
			GROUP BY profile_details.prefs.ui_language
</pre>
