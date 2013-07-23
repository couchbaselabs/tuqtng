## Complex Data - SELECT versus FROM

In the example on the right, the FROM clause fetches all the children records of the contacts. This becomes the input data to the query. Next, the WHERE clause trims the data to only those where the contact name is 'Ian'. Finally, the SELECT clause fetches the first child's name. 

In the query, the child's name is dervied by an expression and has no logical field name. Hence, it would be assigned an auto generated field name, '$1'. To avoid this, we use the AS operator to give an explicit name.

In general, you will find situations where you could either select a smaller set using FROM clause or do a deeper expression in the SELECT clause. You should to do expressions on the SELECT clause as it allows more flexibility when you have the choice.