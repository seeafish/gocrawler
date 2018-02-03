# Simple webcrawler
This piece of code crawls the main page at https://example.com and finds all the internal links.

It does this by looking for any `<a href="/>` tags it can find, it then sends a request to that page and once again looks for further links.

## How to use
Clone or download the source, then run:

`docker build -t crawler:latest`

**Note: During the build, the unit tests will also run. If the tests fail, the build will not complete.**

Once the docker image is built, simply run:

`docker run --rm cralwer`

And wait for the output.

## Todo
* The app only looks one level down from the initial set of links. Should change the `findLinks()` function to be recursive and traverse a tree.
* Remove duplicate links to prevent them from being traversed more than once.
* Make code more efficient by using pointers rather than request a full HTML structure each time.
* Better unit test coverage.