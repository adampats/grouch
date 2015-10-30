# grouch
An object store written Go

I'm building it still...

![under construction](http://i.imgur.com/dkRBjaa.gif)

### Why

Because I want a simple, RESTful, object storage server without having to deploy Riak, Ceph, or Swift.  Think of Grouch as a cross between AWS S3 and CouchDB.

### Usage



### API

**/api** - Management API - for configuring the server.  Auth, storage, perms, etc.

**/auth** - Authentication endpoint.  Where to get an auth token.

**/bucket/:bucket** - Where the buckets are stored.  Do GETs and PUTs here.
