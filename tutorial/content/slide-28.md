## What languages?

What different languages do our users use?

<pre id="example">
 SELECT DISTINCT profile_details.prefs.ui_language 
 	FROM profiles
 		WHERE doc_type = "user_profile"
</pre>