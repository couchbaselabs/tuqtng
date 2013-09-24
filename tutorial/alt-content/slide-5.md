## Aggregates By State

Now let's compute the same statistics grouped by state.

<pre id="example">
SELECT 
	MIN(personal_details.age) AS min_age, 
	MAX(personal_details.age) AS max_age, 
	AVG(personal_details.age) AS avg_age
  FROM profiles
  	WHERE doc_type = "user_profile"
  	  GROUP BY personal_details.state
</pre>