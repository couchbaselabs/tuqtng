## Friend Lists

Join a player's blob with their friends blob. Such types of joins are needed to compute local leaderboards

<pre id="example">
SELECT 
    jungleville.level, 
    friends 
        FROM jungleville 
            JOIN  jungleville friends 
                KEYS jungleville.friends 
        WHERE jungleville.uuid="zid-jungle-0002"
</pre>
