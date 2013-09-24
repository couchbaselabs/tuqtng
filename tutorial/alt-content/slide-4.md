## Calculate Aggregates

Calculate some interesting aggregate functions on the user's personal data.

<pre id="example">
SELECT 
	MIN(personal_details.age) AS min_age, 
	MAX(personal_details.age) AS max_age, 
	AVG(personal_details.age) AS avg_age
  FROM profiles
  	WHERE doc_type = "user_profile"
</pre>
