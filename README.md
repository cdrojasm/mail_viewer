# Email viewer application :mailbox_with_no_mail:

> An application that uses specific content stractor, various indexation strategies and single and concurrent procesing. 

## Features
- Full reactive responsive interface for search in mobile and desktop interfaces. 
- Lightweight search engine. :checkered_flag:
- Concurrent and single approach to extraction and indexation of content. 

## Description 
__This repo contains an small application that develops content extraction and indexation using golang. An user interface to deep down into indexed content and all supported in zincSearch search engine.__

## Profiling and indexation performance

With the aim of find out the performance of the indexer implementation a profiling was carried out. Profiling was developed
used a sample of the dataset with 20K documents selected in a random way. 

### First implementation indexer function

#### Description
Befores start with the task, a function that walk trhough directory of enron mail files and build a list of files available to index is runned. 
Algorithm iterate over this list and execute two process, the first process comprends a content extraction using tags that are always present in the documents. Document content is marshall into a json. After that, marshall json is send to zincSearch. 

Document content extraction is realized at same time that extracted content from plain text is removed with the aim to boost the search of tags along the current document.

#### Results

The first implementation of the indexation function take near of 5 minutes to complete the task. 
Each function was analyzed trhough profiling pprof and then it was checked each part of code that requires enhance. Above images present results of loadStringDataFile, processingDoc and indexingDataToSearchZinc functions. 

### Enhancements plan
With the aim to remove signiticant bottle necks, follow changes in indexer script functions are proposed.
- Enhance file reading functions, it was possible to implement a line by line reading funciton that could be improve this section.
- Delete all fmt.Println, this output to console spend a lot of time in the three functions, single calls could spend near of 50% of time spend.
- Generate precompiled regular expressions to extract content from document files. Trim functions could spend a large amount of resources against regex. 
- Avoid compile regex for each iteration cycle. 
- generate necesary changes to implement go routines during the call of content extractiona and indexing data. 
- integrate walk along enron mail directory. <- this could affects the performance. 