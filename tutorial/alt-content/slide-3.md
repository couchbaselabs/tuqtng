## How many profiles?

Use COUNT(*) to see home many user profiles we have.

<pre id="example">
SELECT COUNT(*)
  FROM profiles
  	WHERE doc_type = "user_profile"
</pre>
