## Navigating in SELECT and FROM

In the last slide, we navigated to a child object using an expression in the SELECT clause. You can do the same thing in the FROM clause.

Navigating to a sub object in the FROM clause limits the input to the SELECT portion of the query to that particular sub object. In the example on the right, the FROM clause fetches the first 'children' object of the 'contacts' document. This becomes the input data to the SELECT clause.

In general, when the expression involves only the '.' and the '[]' operators, you can navigate in either the SELECT or the FROM clause. However, more sophisticated expressions are allowed only in the SELECT clause.

Another difference is that you could omit the 'AS cname' modifier in the SELECT clause. If you did, an auto generated name, 'name' would be used for the attribute. However, when an expression appears in the FROM clause, you must always give it an alias.

Try removing 'AS child' modifier.

<pre id="example">
SELECT child.name AS cname
	FROM contacts.children[0] AS child
</pre>
