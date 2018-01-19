![experimental](https://svg-badge.appspot.com/badge/stability/experimental?f44)

[![GoDoc](https://godoc.org/github.com/julianfrank/gtmesh?status.svg)](https://godoc.org/github.com/julianfrank/gtmesh)
[![Coverage Status](https://coveralls.io/repos/github/julianfrank/gtmesh/badge.svg?branch=master)](https://coveralls.io/github/julianfrank/gtmesh?branch=master) 
[![Build Status](https://travis-ci.org/julianfrank/gtmesh.svg?branch=master)](https://travis-ci.org/julianfrank/gtmesh)
[![Go Report Card](https://goreportcard.com/badge/github.com/julianfrank/gtmesh)](https://goreportcard.com/report/github.com/julianfrank/gtmesh)
[![Sourcegraph](https://sourcegraph.com/github.com/julianfrank/gtmesh/-/badge.svg)](https://sourcegraph.com/github.com/julianfrank/gtmesh?badge)
[![Gitter](https://img.shields.io/badge/gitter-join-brightgreen.svg)](https://gitter.im/jfopensource/gtmesh)

# GT-Mesh
## A Peer-To-Peer Mesh Implementation based on [rsms/gotalk](https://github.com/rsms/gotalk)

## Motivation & History
I found Rasmus's Gotalk to be very useful for RPC and with more experimentation founf the two server RPC capability to be great but wanted to expand to a bigger mesh or >2 Servers. I initially started out forking out the gotalk library but found that this 'Meshing' Capability was more complex than a simple patch. Further not sure of how it would impact the gotalk's core design and hence moved the code to gtmesh and here we go...

## Usage
[TBD]

## To-Do
- [ ] Better Tests
- [ ] Convergence Algorithm Fine Tuning
- [ ] Standard Metrics
- [ ] Metrics Reporting
- [ ] Web Socket Endpoint Functionality
- [ ] REST EndPoint Functionality
