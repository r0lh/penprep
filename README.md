# penprep
Tool to create a pentest environment with useful sub directory structure and note files for technical security audits

## directory structure
```
.
└── pentest
    ├── external // for external pentest
    │   ├── downloads // downloaded files
    │   ├── exploits // successfully used exploits
    │   ├── loot // files and artefacts from pentest objects
    │   │   └── credentials.txt // credentials founded
    │   │   └── screenshots // saved screenshots, e.g. aquatone
    │   ├── notes.txt // notes for external pentest
    │   ├── pub // directory to present for download
    │   ├── scans // scan files
    │   │   └── nmap // nmap output data
    │   │   └── other // other scan data, e.g. gobuster
    │   ├── tmp // temporary garbage
    │   └── tools // downloaded tools for the pentest 
    └── internal // for internal pentests
        ├── downloads
        ├── exploits
        ├── loot
        │   ├── credentials.txt
        │   ├── hashdumps // for saved hashdumps 
        │   ├── tcpdumps // for saved tcpdumps 
        │   ├── bloodhound // for bloodhound data
        │   ├── tcpdumps // for saved tcpdumps 
        │   └── screenshots // saved screenshots, e.g. aquatone
        ├── notes.txt // notes for internal pentest
        ├── pub
        ├── scans // scan files
        │   └── nmap // nmap output data
        │   └── other // other scan data, e.g. gobuster
        ├── tmp
        └── tools
```

## Prerequisites
Install Go on your system. Follow the official [Go Docs](https://golang.org/doc/install) for your platform.

## install
`git clone https://github.com/r0lh/penprep`

`cd penprep && make build`

`make install`
(installed in your golang-environment - check, if you have this directory correctly in your $PATH and rehashed)

OR

`go install github.com/r0lh/penprep@v1.0.4`
(installed in your golang-environment - check, if you have this directory correctly in your $PATH and rehashed)

## usage

`penprep`

OR 

`~/go/bin/penprep`
