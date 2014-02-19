## Case Study II . Social Game 

In this section we look at some typical queries that are needed for a social game applicaton. 
We look at an imaginary game called porkville uses the following three buckets 

```
1. porkville: this bucket contains the user profile blobs which contains a player game 
related data such as level, experience and various other gameplay information. 
Keys in this bucket are named in the following format: zid-pork-<user_id>

2. porkville_stats: this bucket contains the systems stats such as frame-rate, game loading 
time, PvP stats. Keys in this bucket are named as : zid-pork-stats-<user_id>

3. porkville_inbox: This bucket contains a users inbox. Messages sent to a user are appened 
to the existing array of messages. When a message is consumed by the player those messages 
are removed from the message array. Run the query on the right to see what a message blob 
looks like 
```

<pre id="example">
SELECT * FROM porkville_inbox LIMIT 1
</pre>

Run the following queries to examine the content of a user profile and stats blob

```
SELECT * FROM porkville LIMIT 1
SELECT * FROM porkville_stats LIMIT 1
```
