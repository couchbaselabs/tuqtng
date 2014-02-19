## Friend Lists

Join a player's blob with their friends blob. Such types of joins are needed to compute local leaderboards

<pre id="example">
SELECT 
    porkville.level, 
    friends 
        FROM porkville 
            JOIN  porkville friends 
                KEYS porkville.friends 
        WHERE porkville.uuid="zid-pork-0002"
</pre>
