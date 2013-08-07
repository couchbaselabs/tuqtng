## Limiting the Number of Results with LIMIT

Sometimes queries return a large number of results and it can be helpful to process them in smaller batches.

In the example on the right we ask that it return no more than 2 results.

<pre id="example">
SELECT name 
    FROM tutorial 
        ORDER BY name 
            LIMIT 2
</pre>