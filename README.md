# vocabulary_autorecord

A command line tool that can help you remember vocabularies

## Build the Tool

To build the tool, you need to have Go installed on your machine. Once you have Go installed, input the following command in your terminal:

```
go build -o autorecord.exe autorecord.go
```

Or you can use the version directly downloaded from the release page.

## Usage

Move to the position of this tool, execute commands to use this application. 

This tool has the following commands:

```
./autorecord.exe help
```
Get the help information of this tool.

```
./autorecord.exe readFile --path Vocabfile.xlsx --sheetName Sheet1
```
Read the vocabulary from the Excel file and save it to the memory.

```
./autorecord.exe configSetting
```
Change the configuration of this tool.

```
./autorecord.exe addSingleVocabulary -v VOCABULARY -d DEFINITION
```
Add single vocabulary to the vocab list (vocabulary: VOCABULARY, definition: DEFINITION).

## Other Information

To get the sheet name in the Excel file, run (Ctrl + Shift + Enter) the following command in one block of your excel file: 
```
=MID(CELL("filename",A1),FIND("]",CELL("filename",A1))+1,255)
```
