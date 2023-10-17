# penprep
Tool to create a pentest environment with useful sub directory structure and note files for technical security audits

## directory structrue
```
pentest/ 
|- external/          // for external pentest
|-- notes.txt         // notes for external pentest
|-- loot/             // files and artefacts from pentest object
|--- credentials.txt  // credentials found
|-- exploits/         // successfull used exploits
|-- downloads/        // downloaded files
|-- tools/            // downloaded tools for the audit
|-- tmp/              // temporary garbage
|-- pub/              // directory to present for download 
  
|- internal/          // for internal pentests
|-- notes.txt         // notes for internal pentest
|-- loot/             // files and artefacts from pentest object
|--- credentials.txt  // credentials found
|-- exploits/         // successfull used exploits
|-- downloads/        // downloaded files
|-- tools/            // downloaded tools for the audit
|-- tmp/              // temporary garbage
|-- pub/              // directory to present for download 
```

## install
`go install github.com/r0lh/penprep`

(installed in your golang-environment - check, if you have this directory correctly in your $PATH)

## usage

`$ penprep`


