## How many use each language?

Now, how many users use each language?

<pre id="example">
SELECT 
  profile_details.prefs.ui_language AS language, 
  COUNT(*) AS count 
	FROM profiles 
		WHERE doc_type = "user_profile"
			GROUP BY profile_details.prefs.ui_language
</pre>