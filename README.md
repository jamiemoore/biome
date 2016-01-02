# Biome
Manage collections of computers systems (environments).  Associate software, teams, roles and computers to collections of computers (environments).

## What it is
* Manage and view multiple environments
* Assosiate a team to an environment and book it out for a period of time
* Collect the version of the software from the system
* Collect system properites such as cpu count/memory size/disk space.  
* Compare environments 
* Very simple cmdb

##What it is NOT
* Monitoring - install nagios/prometheus
* Metrics Collection - install collectd
* Metrics Display - install grafana

##Architecture

* server - collects information from the client.  Runs the web app
* client - runs from cron collects system/software information and sends it to the server.

##Glossary 

* computer/server/node/system - Can have multiple roles?
* role - one to many relationship to computers
* environment - one to many relationship to roles
* software - many to one relationship to a role
* hardware - refers to both virtual and physical hardware 

Computer must have the following defined -
  * role
  * environment


## Depedancies

```bash
godep save -r ./...
```

##Notes

* Should I use collectd rather than the client?
Collectd could already be in use and this systems should be completely out of band.

