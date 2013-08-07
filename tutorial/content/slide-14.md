## Putting it All Together

The real power of the language comes when we combine all these features together.

The example query on the right combines many of the features of the language into a single query.

Here we match contacts having a name matching the pattern 'da%' or having all of their children over the age of 10.  For each contact satisfying these requirements, we display their name, the number of children, and the full list of children.

This would match 2 contacts ('dave' and 'ian').  We then asked that they be ordered by name in ascending order.  Then we ask to receive only 1 result, skipping over the first result.  This leaves us with the information about 'ian'.

<pre id="example">
SELECT name, LENGTH(children) AS num_children, children 
    FROM contacts 
        WHERE name LIKE 'da%' OR ALL child.age > 10 OVER contacts.children AS child
            ORDER BY name
                LIMIT 1
                    OFFSET 1
</pre>