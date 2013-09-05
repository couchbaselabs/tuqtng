## Review

Lets stop and review a powerful query we can do with what we've learned so far.

Here we match people having a yahoo email address or having all of their children over the age of 10.  For each person satisfying these requirements, we display their full name, email address, and the full list of children.

<pre id="example">
SELECT fname || " " || lname AS full_name, email, children 
    FROM tutorial 
        WHERE email LIKE '%@yahoo.com' 
        OR ALL child.age > 10 OVER child IN tutorial.children END
</pre>