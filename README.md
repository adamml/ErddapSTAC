# ErddapSTAC

A simple programme to create a [STAC](https://github.com/radiantearth/stac-spec)
catalog for the datasets served by instances of NOAA's 
[Erddap](https://github.com/ERDDAP/erddap) sserver software.

## Motivation

Erddap servers provide a consistent interface to environmental data, and
previous solutions have been created to allow federated search across instances
of Erddap (e.g. ). However, STAC provides a simple way of bridging the gap
between these Erddap focussed approaches to federated search and integration
with more genereic spatio-temporal dataset search.

### What is Erddap?

ERDDAP is a scientific data server that gives users a simple, consistent way to 
download subsets of gridded and tabular scientific datasets in common file 
formats and make graphs and maps.

### What is STAC?

STAC is the SpatioTemporal Asset Catalog specification which aims to make
geospatial assets openly searchable and crawlable.

## Data Modeling

In the current implementation the following design approach has been taken:

- There is one top level STAC Catalog which represents all crawled instances of Erddap
- Each encountered Erddap dataset is represented as a STAC Collection in that top level STAC Catalog
- Each STAC Collection is also linked to a STAC Item which represents the NetCDF file offereing of the dataset from an Erddap server

## Getting set up

1. Ensure you have a golang environment set up on your target machine
1. Clone this git repository
1. There are no further external dependencies in the software

## TODO
1. Create a Docker container to allow swift deployment of this software anywhere