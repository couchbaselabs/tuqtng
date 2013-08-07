## Pattern Matching Text with LIKE

Inexact string matching can be accomplished using the LIKE operator in the WHERE clause.

The argument on the right hand side of the keyword LIKE is the pattern that the expression must match.
In these patterns `%` is a wildcard which will match 0 or more characters.  Also `_` can be used to match
exactly 1 character.

Try changing LIKE to NOT LIKE.

<pre id="example">
SELECT * 
    FROM contacts 
        WHERE name LIKE 'da%'
</pre>