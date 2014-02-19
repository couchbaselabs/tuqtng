## Search for music? 

How many of our users have searched for music?

<pre id="example">	
SELECT 
    COUNT(*) music_lovers
        FROM users_with_orders
            WHERE ANY search IN search_history SATISFIES search.category = "Music" END

</pre>
