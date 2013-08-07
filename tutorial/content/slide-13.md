## Skipping over results with OFFSET

When your application has already processed one batch and you want the next batch you can use the OFFSET clause to skip over the results you've already seen.

In the example on the right we ask that it skip over the first 2 results and send the next 2.

<pre id="example">
SELECT name 
    FROM tutorial 
        ORDER BY name 
            LIMIT 2 
                OFFSET 2
</pre>