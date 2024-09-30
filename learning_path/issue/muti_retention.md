# vmselect support different strategy to query multi cluster

## problem

the size of time series data is too large, and the query performance is poor. so I want to split the data into two cluster. one cluster's retention is 7 days and the other one is 6 months. 

vmselect query data from two cluster depend on the query range. I expect that vmselect query recently 7 days time series data from the 7 days retention cluster, and query the other data from the 6 months retention cluster.

## solution

vmselect support different strategy to query multi cluster. 