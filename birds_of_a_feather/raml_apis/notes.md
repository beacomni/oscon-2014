### blueprints, swagger, RAML with Dave Nielsen https://twitter.com/davenielsen 
- Key of api acceptance make it easy to for people to use
    - history: payment gateways --> soap --> REST
    - uddi became a joke
    - but soap had wsdl
    - Strikeiron was an early vision for api marketplace
  - advises against putting version id in a header
    - put it in url
- REST lacks any wsdl-like specification beforehand that you can sample

- mulesoft : osprey (node implementation of a raml mocking service framework)
    - https://github.com/mulesoft/osprey

- apiary --uses markdown but not standardized enough
    - has a mocking service
- swagger --based on yaml, not fleshed out
- raml -- the alternative to those
    - bought Programmable Web
------------------
##### usually:
 - apis are built by hand
 - the whole thing is built before feedback gathered
 - not getting feedback means you'll need a newer api sooner
 - then you need to support multiple apis at once
 - plus the distinct code and datastored behind those
 - whitepaper forthcoming
--------------
 - basically a mocking service allows you to quickly iterate through 
 api consumer feedback on its structure

 - swagger 2.0 is a direct response to RAML

 - in the future, you'd be given an endpoint (not standardized) of 
  where the document describing the api lives (and maybe what it's format is)
---------
 - Dave is a big fan of nodejs and express
  (Dave was the first official api evangelist i think)

 - React stubbing framework, and facebook Flux

 - apis are much more like 'software' than 'websites'.  
since they're being consumed by machines rather
than human beings, they lend themselves to waterfall 
during their lifetime, so you do
the mocking up front with something like RAML