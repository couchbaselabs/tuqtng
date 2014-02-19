## Search Categories for Music Lovers?

What other categories do the music lovers search for?

<pre id="example">
SELECT DISTINCT
    ARRAY search.category FOR search IN search_history 
        WHEN search.category != "Music" END AS search_categories
            FROM users_with_orders
                WHERE ANY search IN search_history SATISFIES search.category = "Music" END
</pre>
