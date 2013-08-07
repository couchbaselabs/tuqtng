## Matching Elements in Nested Arrays with ANY

Sometimes the conditions you want to filter need to be applied to arrays nested inside the document.  In the tutorial sample dataset people contain an array of children, and each child has a name and an age.

In the example on the right we want to find any person that has a child over the age of 10.

The expression after the OVER clause describes which array we want to work with.  The AS label allows us to assign a name to the element in the array we are considering.  The expression after the ANY keyword is the condition that must be satisified by at least one element in the array.

Try changing ANY to ALL.

<pre id="example">
SELECT name 
    FROM tutorial 
        WHERE ANY child.age > 10 OVER tutorial.children AS child
</pre>