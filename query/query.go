//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package query contains artifacts shared across other packages in query
processing.

*/
package query

// Error will eventually include code, message key, and internal error
// (cause) and message
type Error error

// Item is a pipeline data item, i.e. any data produced or consumed in
// query processing
type Item interface {
}

// ItemChannel is a channel of Items
type ItemChannel chan *Item
