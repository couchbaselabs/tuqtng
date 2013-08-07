## Ordering results with ORDER BY

Queries can optionally include an ORDER BY clause describing how the results should be sorted.

In the example on the right we ask that the people be listed in alphabetical order.

Try adding DESC after name.

<pre id="example">
SELECT name 
    FROM tutorial 
        ORDER BY name
</pre>