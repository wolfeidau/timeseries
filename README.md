# Introduction

I deal with a lot of time series data from sensors in the systems I maintain, 
this comes in a few different flavours:

* system performance metrics like cpu and ram usage
* envirnoment sensors like temperature and humidity
* actuations which records a counter indicating how many times an button
  suchlike is triggered.

I also use time series data for monitoring our internal systems and
capturing metrics about how our system is performing, so this is
something I have a keen interest in.

# Overview

This project is designed to be plugged in a micro service in a system,
it's single focus is despatching data revieved via AMQP to one or 
more timeseries store, the first of which is [tempodb](http://tempodb.com).

# TODO

Add some additional backends for the data, these will probably be:

* [Graphite](http://graphite.wikidot.com/)
* [Riak](http://basho.com/riak/)
* [Librato](http://librato.com/)

# Riak Plugin

My aim with this plugin is to use Riak to store the time series data and
possibly Redis to cache and peform aggregate functions.

So we typically have < 10 time series' per customer so the idea was to
have the following structure.

## Storage 

### TimeseriesMeta Bucket

TimeseriesMeta bucket which keeps a record of the timeseries' stored in
the system with some optional meta data.

The key will be the time series ID and the document will contain who
owns the it, along with any configurable attributes such as retension
period.

### User Timeseries Bucket

The user timeseries bucket will named based on the pattern
Timeseries\_[UID], this will contain all the timeseries' for a given
user. This will store a JSON document for each day under the key which
looks like [Sensor ID]-[Date] as seen below.

Key:
````
1_1012BB022412_0101_0_30-2013_10_01
````
Data:
````json
[{"val": 12, "ts": 1380666119164},{"val": 12, "ts":1380666119374}]
````

# Contributing

In lieu of a formal style guide, take care to maintain the existing
coding style. Add unit tests for any new or changed functionality.

# Licence

Copyright (c) 2013 Mark Wolfe
Licensed under the MIT license.
