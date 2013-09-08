/*
Shrike is a simple memory based key value store goshire service. Shrike uses the STREST protocal and by default
is configured to enable http, json and binary listeners.

Shrike is primarily an in memory store but by setting a data_dir value in your config file, Shrike
will dump all data on interupt or sigterm to a backup file which it will load back in next time you
start it.

Goshire comes with a bundled client that supports the binary and json protocols.

For more info on how to use the goshire framework checkout https://github.com/trendrr/goshire.
If there are any issues please submit an ticket.
*/
package main
