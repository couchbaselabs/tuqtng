## <b>Case Study II . Social Game</b>

In this section we look at some typical queries that are needed for a social game applicaton. 
Let's imagine that we are building a game called jungleville that uses the following three buckets: 

<ul>
<li>
<b>jungleville:</b> this bucket contains the user profile data and per user game related data such as level, experience and various other gameplay information. Keys in this bucket are named in the following format: zid-jungle-{user_id}
</li>
<li>
<b>jungleville_stats:</b> this bucket contains the systems stats such as frame-rate, game loading 
time, PvP stats. Keys in this bucket are named as : zid-jungle-stats-{user_id}
</li>
<li>
<b>jungleville_inbox:</b> This bucket contains a users inbox. Messages sent to a user are appened 
to the existing array of messages. When a message is consumed by the player those messages 
are removed from the message array. Run the query on the right to see what a message blob 
looks like
</li>
</ul>

<pre id="example">
SELECT * FROM jungleville_inbox LIMIT 1
</pre>

Try also running the following queries to examine the content of a user profile and stats blob
<span style="color: red">
SELECT * FROM jungleville LIMIT 1
<br/>
SELECT * FROM jungleville_stats LIMIT 1
</span>
