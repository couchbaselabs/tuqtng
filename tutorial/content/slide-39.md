## Assemble a complete user profile blob

When a player loads his gameworld the client application needs to load the data from all the buckets. 
This can be accomplished with by running a single N1QL query. The query on the right assembles a 
the blobs from all three buckets for a user with key "zid-pork-0001"

<pre id="example">
SELECT * 
    FROM porkville AS game-data 
        JOIN  porkville_stats AS stats
            KEY "zid-pork-stats-0001" 
        NEST  porkville_inbox AS inbox 
            KEY "zid-pork-inbox-0001" 
                WHERE game-data.uuid="zid-pork-0001"
</pre>
