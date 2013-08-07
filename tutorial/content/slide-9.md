## Combining Multiple Conditions with AND

The AND operator allows us to match documents satisfying two or more conditions.

In the example on the right we only return contacts having at least one child and have a name matching the pattern 'da%'.

Try changing AND to OR.

<pre id="example">
SELECT name 
    FROM contacts 
        WHERE LENGTH(children) > 0 AND name LIKE 'da%'
</pre>