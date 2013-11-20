# tuqtng implementation notes

* Latest: [implementation notes](https://github.com/couchbaselabs/tuqtng/blob/master/docs/impl.md)
* Modified: 2013-11-18

## Summary

tuqtng is a query execution server process.  At the highest level, the server waits for incoming queries, and upon receiving a query, performs the following steps:

* parse query string into tree structure
* plan execution of query
* choose the best of several plans
* execute the plan
* stream results, errors and warnings back to the client

## Core

There are a few external dependencies and data strucures used throughout the codebase.

### Logging

tuqtng uses the [clog](github.com/couchbaselabs/clog) logging package

This package was chosen because it easily allowed us to turn on/off debug level logging across a whole package.

### Document datastructure

tuqtng uses the [dparval](github.com/couchbaselabs/dparval) package to efficiently work with JSON-like data.

This library was built specifically for tuqtng.  Its major design goal is to allow working with JSON-like data, without parsing unnecessary portsin of a document.  dparval.Value objects can be instantiated either from a raw byte array, or using existing in-memory object structures (JSON compatible).

The library also allows for attaching other related data to the object, while not contaminating the object itself.  This allows the query engine to attach its own internal data to objects during proecssing.

Throughout the codebase when components are passing user data around, they pass dparval.Value objects.

## Package Details

### Main

The main package does very little work.  It parses the command-line parameters, instantiates the corresponding catalog (file/couchbase/mock), network (http), and server objects.  The server is then started.

### Catalog

The catalog package is an abstraction around the capabilities of the underlying database.  All access to index and data goes through the catalog API.

There are 4 implemenations of the catalog API:

* couchbase - uses go-couchbase client library to talk to a couchbase server cluster
* file - uses json files and a specific directory layout
* mock - uses in memory representation
* system - a wrapper catalog which is able to introspect the catalog it wraps, and expose the system catalog as additional buckets in a pool named "system"

The current code will instantiate the type of system catalog that the user requested with the commmand-line arguments.  Next it will instantiate an instance of the system catalog to wrap the one the user requested.  Subseqeuntly, all code will use the system catalog implemenation (calls to non-system buckets are passed through to the underlying catalog implemenation)

### Network

The network package is an abstraction around the interaction between the tuqtng server and client.

A network endpoint is given a QueryChannel, to which is should send all incoming queries.  Queries are represented with the Query object, which contains a QueryRequest, QueryReponse, and a StopChannel.

There is only one network endpoint implemenation, HTTP.  This endpoint is responsible for turning incoming HTTP requests into QueryRequest objects.  Currently, this is just a string representation StringQueryRequest.  It is also responsible for creating a QueryResponse object.  This is done by HttpResponse which is responsible for serializing results, errors and warnings and returning them to the client.

The StopChannel is a channel, which the network endpoint will close, if for any reason it thinks the client is gone and no longer interested in receiving results.  This allows the server to abort execution and stop any expensive processing.

### Server

The server package instantiates two new components, the Compiler and Executor.  These objects are passed a reference to the Catalog, as they both need it to perform their work.

Then the server will read queries off of the QueryChannel.  Each query is then passed to the compiler, resulting in a Plan.  The Plan is then passed to the Executor.  The Executor is given reference to the Query, and directly sends results, warnings and errors through.

### Compiler

The compiler package is an abstraction around a component which turns a query string into a Plan.

There is one implementation of this interface, named StandardCompiler.  The standard compiler consists of three components, the parser, the planner, and the optimizer.  These components are used to perform the following tasks on each incoming query string:

* parse query string into AST
* verify the query AST is semantically valid
* transform the AST into a simpler form (eliminate constants, redundant expressions, etc)
* generate plans for the AST
* choose optimal plan

### Parser

The parser package is an abstraction around the component which turns a query string into an ast.Satement.

There is one implementation of the parser named goyacc, which uses the [nex](http://www-cs-students.stanford.edu/~blynn//nex/) lexer and the [go yacc](http://golang.org/cmd/yacc/) parser generator.

There is a separate build.sh file in this directory which will regenerate the go source files corresponding to the n1ql.nex and n1ql.y grammar files.  The generaged go files are checked into the github repo so that developers not interested in changing the grammar can still install/build the server using the built-in go tooling.

nex uses regular expressions to create tokens from the input.

go-yacc consumes these tokens and as rules are matched the AST is built.

**NOTE**: useful parse errors are not generated at this time.  this is an area open for a lot of improvement.

### AST

The AST is formal structure representing a parsed query/statement.  There are two high-level interfaces of importance:

* ast.Expression - the core interface used for evaluating expressions
* ast.Statement - a loose wrapper around different types of statements

The ast.Statement interface isn't terribly useful at this time.  Most of the code uses a type assertion and works directly with the underlying structures.

**NOTE**: While the structure and naming of the packages make it appear somewhat generic, the AST package is actually the implementation of the semantics of n1ql.  In the future it would probably make sense to reorganize the n1ql specific things together, which currently would include the ast package, and the parser.

### Planner

The planner package is an abstraction around the component which turns an ast.Statement into one or more plan.Plan objects.

There is one implemenation of planner named SimplePlanner.  The SimplePlanner is no longer simple.  It operates as follows.

First, a type assertion switch case calls a different method to handle different types of statements.  Create/drop index statements are basically just wrappers and not that interesting.  We will discuss how plans for SelectStatements are generated.

1.  The datasource referenced in the FROM clause is looked up in the catalog.
2.  All  access paths (indexes) are looked up for this data source.
3.  Each of these access paths are examined, to see if they can be used for this query.  If so, a new plan head is generated.  A plan head is either a SCAN, FETCH, or SCAN followed by FETCH, depending on the query/index match.

At this point, if don't have any query heads, the planner aborts (should never happen unless the bucket was deleted after the query passed semantic validation).

For each plan head, additional operators are added to the plan pipeline.  For example, if the query had a WHERE clause, a Filter operator would be added.  If the query had a LIMIT clause, a Limit operator would be added.

After all the query clauses have been added to the query plan pipeline, the plan is sent on the PlanChannel, and the process is repeated for the next plan head.

After all plan heads have been built into plan.Plans, the PlanChannel is closed.  This informs the optimizer that there are no more plans to consider.

### Optimizer

The optimizer package is an abstraction around the component which considers multiple plans and chooses the best one.

Currently there is one implemenation named SimpleOptimizer.  The SimpleOptimizer does not currently do any quantitative comparison of the plans.  Instead, it simply returns the last plan emitted by the planner.

### Executor

The executor package is an abstraction around the component which takes a single plan.Plan and executes it to produce results, warnings and errors.

There is one implemenation of Executor named InterpretedExecutor.  This implementation uses another package *xpipelinebuilder* to do instantiate xpipeline.Operator objects from their corresponding plan.PlanElement objects.

The Operator elements are connected to one another as described in the plan.  Its helpful to think of this as a pipeline.  Typically a SCAN operator at the head of the pipeline generates documents, which flow through the pipeline, and result objects come out the end of the pipeline.

Each plan.Operator is connected to its upstream operator.  When an operator is asked to run, it asks its source operator for the channel to read objects from, and then invokes source.Run() on a separte go routine.  This in turn starts the source operator.  After starting its source operator, it attempts to read objects from the sources channel.

Each operator also has a StopChannel that can be closed to signal upstream operators that they should stop doing any work.  All operators must monitor their downstreams StopChannel and be prepared to stop as soon as possible in the event it is closed.

For example, the Limit operator will close the StopChannel after it has processed the specified limit number of documents.  This signals all of the upstream operators that they no longer need to keep working, otherwise they would continue to perform a large amount of unnecessary work.

## Interesting Sections of Code

### Functional Dependency Checking

Many times during processing it is useful to know if a particular expression can be satisfied by a list of other expressions or if it depends on something else.

For example the expression *1 + age* cannot be satisfied by an empty expression list, because it DEPENDS on the evaluation of the expression *age*.

The object *ExpressionFunctionalDependencyChecker* allows us to perform this check.  It uses the visitor pattern.  You instantiate it with a list of Expressions you have available.  And then you ask it to visit the Expression in question.  If it cannot be satisfied by the provided expressions it returns an error.

### Expression Simplification

One example usage of functional dependency checking is expression simplification.  We simply ask all non-literal expressions if they can be satisfied by the empty expression list.  If they can, it means they don't depend on any external factors.  If they don't depend on any external factors, then we can evaluate them now with an empty context, and replace them in the AST with the literal value of the result.

For example, the expression *1+1* does not depend on anything.  So we evaluate it in an empty context, get the result of 2, and replace *1+1* in the AST with new literal value *2*.

### Semantic Validation

There are many queries which are syntactically valid, but semantically invalid.  The purpose of semantic validation is to detect these queries and give an informative message to the user.

The *ast.Statement* interface exposes a method VerifySemantics().  The compiler invokes this method after parsing the statement, but before passing it to the planner.

The following steps are performed to verify that a select statement is valid (with explanatory notes where necessary):

1.  A list of explicit aliases introduced in the projection is computed.  (SELECT abc AS xyz ...)  This list is checked for duplicates.  Duplicate aliases are an error.
2.  All expressions in the projection that were NOT explicitly named are assigned a name.  This could either be a meaningful name (the document property age gets the alias "age") or they could be meaningless (the expression 1+1 gets the alias "$1")  If an assigned alias conflicts with an explicitly specified alias, this is an error.  If two assigned aliases are duplicates, this is also an error.
3. The FROM clause is then fixed up.  Sereral things happen at this stage:
    1.  The bucket is identified.  (sometimes its part of a complex expression, like beer-sample.addresses)
    2.  If the bucket was part of a complex expression, the projection sub-expression is fixed up (beer-sample.addresses[0] separates out the bucket portion and leaves the remaining projection as addresses[0])
    3.  Aliases for the bucket is generated (if it wasn't explicitly specified).  Most but not all from expressions can be automatically assigned a name.  Alias assignment is invoked recursively on any nested FROM expression (due to an OVER clause)  **NOTE**: we use the same FROM structure to represent chained OVERs
4.  Now that all buckets and OVERs have proper aliases, all path expressions are rewritten into *formal notation*.  Formal notation means all paths begin with a bucket alias or OVER alias.  For example, *SELECT name FROM beer-sample* becomes *SELECT beer-sample.name FROM beer-sample*.  For queries queries with more than one named source (like bucket alias and an OVER alias) we REQUIRE that they fully-qualify all references.  This means that *SELECT name FROM bucket OVER path* would be an error, because *name* was not properly qualified.  Errors like this are detected at this phase because we have a list of all the valid aliases.  In the single alias case we prefix the alias if it is missing, in the multiple alias case there is no default and an error is generated.
5.  At this point, all expressions need to be validated.  This step invovles invoking an expression validator on all parts of the statement where an Expression can occur.  Expression validators are built using the Visitor pattern.  Validation currently checks the following things:
    1.  Some clauses cannot contain aggregate functions (WHERE, GROUP BY clauses)
    2.  Only the COUNT() function allows * as an argument
    3.  Only COUNT() and ARRAY_AGG() functions allow the DISTINCT qualifier
    4.  Unknown function names are detected
    5.  Function arity is validated
6.  Finally, additional checks must be made if this is an aggregate query:
    1.  Verify that the projection consists of either aggregate functions OR expressions that depend ONLY on the GROUP BY clause.
    2.  Verify that the HAVING clause depends ONLY on items available in the GROUP BY clause.
    3.  Verify that the ORDER BY clause is either an aggregate function OR a reference to explicit projection alias.
    4.  IF, this query ALSO has a DISTINCT qualifier there is one additional check we must enforce.  ALL expressions in the ORDER BY clause must ALSO be in the projection (this ensures deterministic output by projecting the fields used for ordering, thus affecting which rows were kept during DISTINCT removal)

### Query Optimization

Currently all query optimizations are chosen based on the existance and suitability of a single index.  Essentially the list of access paths is walked sequentially and compared with various aspects of the query.  Depending on the type of optimization in question, if everything matches, a query head is produced.  The following optimizations are supported.

#### Portion of the WHERE clause can be converted into one or more range scans on an index

The following steps are performed given an index.

1.  Is this index a RangeIndex?  If no, stop.
2.  Does it have a WHERE clause?  If no, stop.
3.  Does the FROM clause project a sub expression?  If yes, stop.
4.  The first expression that the index was built (could be multi-part key, other parts currently ignored) on is converted to formal notation (in the context of this query).  So if the index was built on the field "abv", and in this query the bucket was given alias "b".  Then for our purposes, the index expression in formal notation is "b.abv".
5.  WHERE clause is converted to Negation Normal Form.
6.  WHERE clause is converted to Conjunctive Normal Form.  WHERE clause should now either be a simple predicate or an AND expression.
7.  If the clause is a simple expression, we check to see if it is sargable.  If the clause is an AND expression we check each consituent piece, to see if any of them are sargable.
8.  If a simple expression was sargeable, then we get a single range to scan returned.  If it was an AND expression and multiple constituents were sargable, and attempt is made to combine them and return the smallest sets of ranges.  (abv > 5 AND abv < 7, should yield a single range between 5 and 7)

#### Projection of MIN(expr) can scan a single row

1.  If we've already deteremined that we can do a range scan AND we determine that the projection only contains MIN() on an expression that only depends on the same expression as the index, then we can scan one row.  (ie, MIN(abv) WHERE abv > 5, can scan single row)
2.  If we're not doing a range scan on an index, but there is NO group by clause (aggregating across whole bucket) AND the projection is ONLY MIN(*), then we can scan one row from all_docs.

#### Index covers projection

If the index contains all the values needed to satisfy the projection, then the FETCH operator is eliminated.  The SCAN operator will reconstitute enough of the document from the index to allow for evaluation of all the projection values.

#### Fetch directly by IDs

If the WHERE clause is in a form recognized to be a direct match against either a single ID or a list of IDs then the SCAN operator is eliminated and replaced with a direct list of IDs to fetch.

### Query Optimization Notes

#### FILTER operator not removed, even when range scanning an index

Because there is no consistency between index scans and document fetches, we could possibly fetch documents which have been updated.  Thus to avoid returning obviously wrong results, we re-evaluate the entire WHERE clause in memory, even if the index told us it would be satisfied.

#### MAX(expr) not optimized

Because of the collation order, we can ensure a mechanism to scan the first non-eliminated row for a MIN() query (we want the first non-null, non-missing value).  We do this by starting a scan at null, non-inclusive.  Unfortunately there is no such corresponding high-value from which we can do the same for MAX().  Its possible a new API method could emulate this behavior.

#### Fast Count optimization disabled

We implemented an optimization to perform fast COUNT() queries.  Unforunately we also later found out that the "total_rows" field returned by the view engine can be wrong during a rebalance operation.  There seemed to be no easy way to know for sure that a rebalance was occurring at the same time we read the "total_rows" field, so we disable this optimization.  (the code is there but commented out)