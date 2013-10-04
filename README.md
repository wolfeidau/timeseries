# Introduction

I work on at (Ninjablocks)[http://ninjablocks.com] which provides a developer platform for internet of things (IoT). We deal with a lot of time series data from sensors in our system, this comes in a few different flavours:

* system performance metrics like cpu and ram usage
* physical sensors like temperature and humidity
* actuations which records a counter indicating how many times an button suchlike is triggered.

# Architecture

My aim with this system is to use Riak to store the time series data and possibly Redis to cache and peform aggregate functions.

So we typically have < 10 time series' per customer so the idea was to have the following structure.

## Storage 

TimeseriesMeta bucket which keeps a record of the timeseries' stored in the system with some optional meta data.

The key will be the time series ID and the document will contain any configurable attributes such as retension period.

Timeseries\_[UID] bucket which contains all this users timeseries' for a given user. This will store key as [ID]-TIMESTAMP and some JSON.
````
{val: [number]} 
````

Where id could be something like 1\_1012BB022412\_0101\_0\_30 and number is say 12.

## Garbage Collection

At some point we will need to start cleaning out old data within riak, this

