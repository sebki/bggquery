# Boardgamegeek XML API 2 Client
Go library for connecting to the Boardgamegeek XML API 2, as discribed here: https://boardgamegeek.com/wiki/page/BGG_XML_API2 . Work in Progress

## Endpoints added so far:
- [x] Thing Items
- [x] Family Items
- [x] Plays Items
- [x] User Items
- [x] Collection Items
- [x] Forum List Items
- [x] Forums Items
- [x] Threads Items
- [x] Guilds Items
- [x] Hotness Items
- [x] Search

## Structure
All interfaces, global datatypes and consts are within bggquery.go, where the Query() function is also located. All API endpoints have there own ...Query struct, each with own methods to set query options. The Query struct is eventually passed to the query function, where the response is written to the respective ...Items Struct, which is located in the \_data.go files. 

## Credits
All XML-Structs in this library are made with https://www.onlinetool.io/xmltogo/
