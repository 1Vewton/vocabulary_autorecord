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
Change the configuration of this tool.(Include settings for the llm)

```
./autorecord.exe addSingleVocabulary -v VOCABULARY -d DEFINITION
```
Add single vocabulary to the vocab list (vocabulary: VOCABULARY, definition: DEFINITION).

```
./autorecord.exe vocabularyManager
```
Manage the vocabulary list. Currently this can only delete vocabulary from the list. 

```
./autorecord.exe normalExercise
```
Start an exercise about vocabulary in normal mode. (Think about the definition of a vocabulary)

```
./autorecord.exe llmExercise
```
Start an exercise about vocabulary in llm mode. (Type the definition of vocabulary and the llm will check the correctness of the definition)

## Other Information

To get the sheet name in the Excel file, run (Ctrl + Shift + Enter) the following command in one block of your excel file: 
```
=MID(CELL("filename",A1),FIND("]",CELL("filename",A1))+1,255)
```

# Theories

This program is based on the following theories:

## BKT

This program utilizes BKT (Bayesian Knowledge Tracing) to assess whether the student memorizes the vocabularies. BKT assumes that the state of acquiring of knowledge is binary. It will get a possibility between 0.0 and 1.0 of whether the student acquired this knowledge. The closer the possibility is to 1.0, the higher the possibility of acquiring this knowledge. 

Here are some basic parameters for this method, you can change the certain value in `.env` to change the value for these parameters: 

| Parameter Name | Meaning                                                                                                                                                                                                                                                                                          | Field Name in `.env` |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------- |
| $p(L_0)$       | The possibility of student acquired this knowledge before studying.                                                                                                                                                                                                                              | `PL0`                |
| $p(T)$         | The possibility of not acquiring this knowledge after certain exercise                                                                                                                                                                                                                           | `PT`                 |
| $p(G)$         | The possibility of not acquiring this knowledge but gues it right. (In this program's case, the user might guess the meaning of the vocabulary right through inferring the meaning from the similarity of this word to another word or just believes that the definition in the mind is correct) | `PG`                 |
| $p(S)$         | The possibility of already acquired the knowledge but get the wrong answer (In this program's case, the user might temporarily cannot remember the meaning of this word).                                                                                                                        | `PS`                 |

The default value for these parameters are based on reaserch by *Yi et al.*[^1] based on ASSISTments Platform. 

The followings are the formulas. $p(L_{t})$ means the possibility of acquiring the knowledge on the $t^{th}$ practice. 

We will first predict the possibility of user answering correct in the t-th practice: 

$$p(\text{correct})={p(L_t) \cdot (1-p(S)) + (1-p(L_t)) \cdot p(G)}$$

If the user answers correctly: 

$$p(L_t | \text{correct})={p(L_t) \cdot (1-p(S)) \over p(\text{correct})}$$

If the user answer incorrectly: 

$$p(L_t | \text{incorrect})={p(L_t) \cdot p(S) \over 1-p(\text{correct})}$$

Than we will get $p(L_{t+1})$: 

$$p(L_{t+1})=p(L_t|\text{answer})+(1-p(L_t|\text{answer})) \cdot p(T)$$

## Time Decay

In here, we assume that the studied possibility would decay to 1/2 of originall studied possibility after 7 days, there for, we add a equation for each update to calculate the sutdied possibility after $\Delta D$ days: 

$$p(L_{after})=p(L) \cdot e^{-\lambda \cdot \Delta D}$$

Where $\lambda=0.09902$, when $\Delta D = 7$, the possibility would be 1/2 of original possibility. This equation is aiming to simluate the exponential decay of memory. 

You can alter `Lambda` value in `.env` file to change $\lambda$ value. 

# Reference

[^1] Qiu, Y., Qi, Y., Lu, H., Pardos, Z. A., Heffernan, N. T., & Worcester Polytechnic Institute. (2021). Does time matter? Modeling the effect of time in Bayesian knowledge tracing. In Worcester Polytechnic Institute [Journal-article]. https://web.cs.wpi.edu/~nth/pubs_and_grants/papers/2011/EDM%202011/Qiu%20Does%20Time%20Matter.pdf