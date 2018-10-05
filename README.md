# Go based Screeps Engine
This is an implementation of the Screeps engine in Go. 

For now, its focused on implement the `processor` module, which handles all the
intents for rooms each tick.

Theoretically, the entire server could be implemented in Go as a single self 
contained binary. However, implementing the runner and backend are currently
beyond the scope of this project.

The project is currently hardcoded to look for mongo and redis on default ports.
It uses the `screeps` database in mongo. In the future this should be handled
by either a dedicated config file, ENV vars or by parsing .screepsrc. Config
loading needs to be generic enough to support all three.

I have included a test mod for the server that will disable the node processor.

As of writing this, only the `move` and `tick` intent for creeps are
implemented, neither is a complete implementation.