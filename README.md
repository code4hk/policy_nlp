# policy_nlp
basic textual analysis for policy address

有咩比施政報告更無意義，就係施政報告既文字分析
成日話施政報告無新意，其實有幾似？

1. jieba to tokenize with dictionary based n-gram analysis
2. compare cosine similarity using top 5000 terms (i.e. all included here)
3. also print top 20 terms

`go run similar.go`

```
2014年 關鍵字 Top 20
我們/ 發展/ 香港/ 政府/ 服務/ 土地/ 提供/ 計劃/ 房屋/ 社會/ 政策/ 委員會/ 研究/ 增加/ 需要/ 工作/ 人士/ 經濟/ 問題/ 加強
2014年 關鍵字 Top 20
政府/ 發展/ 香港/ 計劃/ 服務/ 提供/ 支援/ 增加/ 土地/ 加強/ 資助/ 工作/ 以及/ 經濟/ 學生/ 研究/ 房屋/ 去年/ 基金/ 繼續
2015年 關鍵字 Top 20
政府/ 發展/ 香港/ 服務/ 計劃/ 提供/ 我們/ 土地/ 研究/ 經濟/ 繼續/ 增加/ 委員會/ 加強/ 工作/ 去年/ 以及/ 單位/ 資助/ 房屋
```

```
vector1:2782 vector2:2905 interesect:1596
0.0020103644793905205  /  0.0021957462915481183
Cosine Similarity 2014 vs 2015: 0.9155722986434404

vector1:2974 vector2:2905 interesect:1588
0.0018985431827035877  /  0.0021389635491341607
Cosine Similarity 2013 vs 2015: 0.8875995962961156

```
