## Player v/s Player Stats

Head to head matchups between a particular user and all others that she has competed against

<pre id="example">
    SELECT 
        stats.uuid AS player, 
        hist.uuid AS opponent, 
        SUM(CASE WHEN hist.result = "won" THEN 1 ELSE 0 END) AS wins, 
        SUM(CASE WHEN hist.result = "lost" THEN 1 ELSE 0 END) AS losses
            FROM jungleville_stats AS stats 
                KEY "zid-jungle-stats-0004" 
                    UNNEST stats.pvp-hist AS hist
                        GROUP BY stats.uuid, hist.uuid
</pre>
