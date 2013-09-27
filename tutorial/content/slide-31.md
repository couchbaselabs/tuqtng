## Users' Searches

How many searches have the users run?

<pre id="example">
SELECT 
  LENGTH(search_history) AS num_searches, 
  COUNT(*) AS num_profiles 
	FROM profiles 
		WHERE doc_type = "user_profile"
			GROUP BY LENGTH(search_history) 
				ORDER BY num_searches
</pre>