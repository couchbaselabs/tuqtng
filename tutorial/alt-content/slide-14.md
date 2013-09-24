## Search Categories for Music Lovers?

What other categories do the music lovers search for?

<pre id="example">
SELECT 
  ARRAY search.category OVER search IN profiles.search_history 
  WHEN search.category != "Music" END AS search_categories
    FROM profiles
		WHERE 
			ANY search.category = "Music" OVER search IN profiles.search_history END
</pre>