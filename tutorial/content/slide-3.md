## Summary

Unstructured Query Language 2013 (UNQL 2013) is a query language for Couchbase.

This language attempts to satisfy these [requirements](https://github.com/couchbaselabs/tuqqedin/blob/master/docs/requirements.md).

This document describes the syntax and semantics of the language.

## Statement

An UNQL statement is an instance of:

unql-stmt:

### EXPLAIN

An UNQL statement can be preceded with the keyword "EXPLAIN".  This causes the statement to return information about how the UNQL statement would have operated if the EXPLAIN keyword had been omitted.

The output from EXPLAIN is intended for analysis and troublehsooting only.  The details of the output format are subject to change.

### SELECT Core

Before looking at the full SELECT statement, let us start with a simpler subset:
Hello WOrld
