deepCopy
========
[![GoDoc](https://godoc.org/github.com/mohae/deepcopy?status.svg)](https://godoc.org/github.com/mohae/deepcopy)[![Build Status](https://travis-ci.org/mohae/deepcopy.png)](https://travis-ci.org/mohae/deepcopy)

DeepCopy makes deep copies of things: unexported field values are not copied.

## Features

 - Copy from field to field
 - Copy from slice to slice
 - Copy from struct to slice
 - Copy from map to map
 - Deep Copy

## Usage
    cpy := deepcopy.Copy(orig)
