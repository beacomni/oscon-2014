###Pete Hunt: how instagram.com works
 - slides will probably be at http://www.slideshare.net/floydophone
 - spa rendered entirely in js, using React.
 - built by a very small team
 - high performance data visualization tooling (downloading efficiently to client)
 - why even build a web instance? app on demand, upgrades automatic
 -----------
####SPA
 - data fetching and routing are different problems  
first: traditional, like rails. server for every request
2: server-rendered + ajax
trendy: spa. very strong in terms of user experience

##### increasing spa performance
 - minimize # of http requests
 - minimize # bytes downloaded
 - concat all js into 1 file is wrong
  - at instagram parsing and evaluate 9.5mb (download size is gzipped so smaller)
  - only download the js you need for each "page" / entry point
   - instagram.com has 17 entry points
  - subsequent navigations should not download the same js again
  - instagram uses commonjs for module system  
  - bundle using browserify, requirejs
  - optimize the bundles 
  - asynchrounously load bundles on demand using split points
##### other static resources
  - files like .css can also be considered to be a part of the dependency graph
   - tie an entry point to its styling
   - use require for css, just like you would for js
   - same for images or any static asset like .less
   - instagram eliminated grunt
##### hitches with css
    
   - any or all css modules may be loaded at any time, in any order
   - hard to know if a rule is no longer used

   - css opinions:
    - namespaced, unambiguous class names
    - no cascading, single class name selectors only
    - no overriding (multiple class names applied to a dom node should do unrelated styling)
    - see yahoo team's atomic
    - yandex BEM
    - nicole sullivan object oriented css, nicholas gallagher suit css
   - requirejs likes to load large files
   - grunt, gulp don't know about the dependency graph
   - webpack!! used in production
   - webpack how-to in petehunt github

##### testing .js
   - no tests written in prototyping mode, since its going to be thrown away anyway
    - when getting toward the end
     - huxley for screenshotting compare
     - using 100% React now