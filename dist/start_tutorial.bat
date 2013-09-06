@echo off

pushd "%~dp0"
echo.
echo In your web browser open http://localhost:8093/tutorial/
echo.

start "TuqTng" "http://localhost:8093/tutorial/"

tuqtng -couchbase dir:data

popd
