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
    │   ├── notes.txt // notes for external pentest
    │   ├── pub // directory to present for download
    │   ├── tmp // temporary garbage
    │   └── tools // downloaded tools for the pentest 
    └── internal // for internal pentests
        ├── downloads
        ├── exploits
        ├── loot
        │   ├── credentials.txt
        │   └── hashdumps // for saved hashdumps 
        ├── notes.txt // notes for internal pentest
        ├── pub
        ├── tmp
        └── tools
```

## install
`go install github.com/r0lh/penprep`

(installed in your golang-environment - check, if you have this directory correctly in your $PATH)

## usage

`$ penprep`


