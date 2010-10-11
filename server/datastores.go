package main

/*
Convenience file.

The listing below forces the imported packages to run their init() functions, 
which makes them register themselves as a valid Datastore.

Any new datastore implementations should be added to this list in order to
have them take effect.
*/

import _ "mudkip/mongo"
import _ "mudkip/mysql"
import _ "mudkip/sqlite"
