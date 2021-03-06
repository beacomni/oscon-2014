###Alejandro Cabrera (see github) : The Case for Haskell / Ask More of Your Languages
 - slides at: https://blog.cppcabrera.com/posts/oscon-talk-2014.html
 - types are a valuable tool for enforcing and for communicating what we can do
    - help prove our types correct
    - complementary to testing
    - communicate intent and abstraction
 - user defined types: typedef that compiler treats as distinct from the original
    - rough approximation is to wrap simple types in c structs
 - **parametric polymorphism**, like generics
 - **product types:** compile time tuples
 - **sum types:** compile time enforced enum, disjoint union
 - **recursive type:** express infinite data structures
 - **higher kinds:** types that exist a level above types.  'what a 
   valid type should look like'.
    - types descripe shape of values
    - kinds describe shape of types
 - **effect tracking** -- allows for compiler to detect when
   - code would change state,
   - side effects ineracting with 'world' would happen

 - **pattern matching:** write functions against shape of the data
     - less verbose than switch
---------
#####more on types
 - type system: iteratively prove your programs correct curry-howard style
    - Stephanie Weirich
 - types communicate to others unambiguously what you mean
 - a programming language is the frontend for you type system
#####future:
- type system advancments:
- dependent types
-  lineary types
-  smarter gradual typing
  
#####smarter tools:
- hoogle hyoo
- lamdu, light table
- rest, ivory
