Stackdriver Logging のログエントリに `metadata` だけを書く例  
`entry(metadata?, data?)` で引数が1つなら `data` として書く実装を迂回すれば良い  
https://github.com/googleapis/nodejs-logging/blob/263e046603fb8dc105653b860f4936add4c45f71/src/log.ts#L419
