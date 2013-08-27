package couchbase

const templStart = `
function (doc, meta) {
  if (meta.type != "json") return;`

const templFunctions = `
  var stringToUtf8Bytes = function (str) {
    var utf8 = unescape(encodeURIComponent(str));
    var bytes = [];
    for (var i = 0; i < str.length; ++i) {
        bytes.push(str.charCodeAt(i));
    }
    return bytes;
  };

  var indexFormattedValue = function (val) {
    if (val === null) {
      return [$null];
    } else if (typeof val == "boolean") {
      return [$boolean, val];
    } else if (typeof val == "number") {
      return [$number, val];
    } else if (typeof val == "string") {
      return [$string, stringToUtf8Bytes(val)];
    } else if (typeof val == "object") {
      if (val instanceof Array) {
        return [$array, val];
      } else {
        return [$object, val];
      }
    }
  };`

const templExpr = `
  var $var = indexFormattedValue($path);`

const templKey = `
  var key = [$keylist];`

const templEmit = `
  emit(key, null);`

const templEnd = `
}
`
