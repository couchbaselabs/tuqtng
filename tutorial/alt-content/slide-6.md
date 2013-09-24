## Highest Average

Let's order these results with an ORDER BY clause and restrict to one result with a LIMIT clause.

<pre id="example">
SELECT 
    MIN(personal_details.age) AS min_age, 
	MAX(personal_details.age) AS max_age, 
	AVG(personal_details.age) AS avg_age
  FROM profiles
  	WHERE doc_type = "user_profile"
  	  GROUP BY personal_details.state
  	  	ORDER BY avg_age DESC
  	  		LIMIT 1
</pre>