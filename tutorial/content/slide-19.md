## Combining Document Joins with Group By

Twitter tweet topics

<pre id="example">
SELECT topic, COUNT(*) 
	FROM tweets 
		OVER tweets.topics AS topic 
			GROUP BY topic
				HAVING COUNT(*) >= 2
</pre>