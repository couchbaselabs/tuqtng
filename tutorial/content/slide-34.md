## Search for music? 

How many of our users have searched for music?

<pre id="example">	
SELECT COUNT(*)
    FROM profiles
		WHERE ANY search.category = "Music" OVER search IN profiles.search_history END
</pre>