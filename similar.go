package main

import (
  "fmt"
  "github.com/wangbin/jiebago"
  "github.com/wangbin/jiebago/analyse"
  "io/ioutil"
  "math"
  "strings"
  )

  func check(e error) {
    if e != nil {
      panic(e)
    }
  }

  var topKSimilar = 5000

  //refer to https://github.com/gyuho/goling/blob/master/similar/similar.go#L16
  func getCosine(vect1 map[string]float64, vect2 map[string]float64) float64{
    // intersection contains common characters
    intersection := []string{}
      // traverse keys and add what is common to both.
      // (keys, in Hash/Map, are like indices in array)
      for key := range vect1 {
        if _, exist := vect2[key]; exist {
          intersection = append(intersection, key)
        }
      }

      fmt.Printf("vector1:%v vector2:%v interesect:%v \n",len(vect1),len(vect2),len(intersection))

      // If all the vector elements are equal, cos will be 1.
      // Equal texts return the value 1.
      // We need to traverse the intersection(common) characters of texts.
      // In doing so, we can expect two same texts to return 1, cos 0°

      // to calculate A·B
      sum := 0.0
      for _, elem := range intersection {
        sum += float64(vect1[elem]) * float64(vect2[elem])
      }
      numerator := sum

      // to calculate |A|*|B|
      sum1 := 0.0
      for key := range vect1 {
        sum1 += math.Pow(float64(vect1[key]), 2)
      }
      sum2 := 0.0
      for key := range vect2 {
        sum2 += math.Pow(float64(vect2[key]), 2)
      }
      denominator := math.Sqrt(sum1) * math.Sqrt(sum2)

      // smoothing because we can't divide by 0
      if numerator == 0.0 || denominator == 0.0 {
        return 0.0001
      }

      fmt.Println(numerator," / ",denominator)
      return float64(numerator) / denominator
  }

  func getTagsAsVector(data []byte) map[string]float64 {
    content := string(data)
    topTfIdfs := analyse.ExtractTagsWithWeight(content, topKSimilar)

    vect := make(map[string]float64)
    for _, ti := range topTfIdfs {
      fmt.Println(ti)
      // fmt.Println(reflect.TypeOf(ti))
      fmt.Println(ti.Freq())
      vect[ti.Word()]=ti.Freq()
    }
    return vect

  }
  //TODO top K with count
  func printTopK(tags []string, data map[string]float64,topKWord int){
    // i :=1
    // for _,t := range tags{
    //   fmt.Printf("  %s :%v \n", t,data[t])
    //   // i++
    // }

    fmt.Println(strings.Join(tags, "/ "))
  }

  func printAll(vector map[string]float64,year int){
    fmt.Println("%d年 關鍵字/Tf-idf",year)
    for w,f := range vector{
      fmt.Println(w+": ",f)
    }
  }

  func main() {
    jiebago.SetDictionary("dict.txt.big")

    topKWord :=20
    data2013, err := ioutil.ReadFile("2013.txt")
    check(err)
    data2014, err := ioutil.ReadFile("2014.txt")
    check(err)
    data2015, err := ioutil.ReadFile("2015.txt")
    check(err)

    // fmt.Print(string(data2013))
    // fmt.Print(string(data2014))
    // fmt.Print(string(data2015))
    // analyse.SetIdf("/Path/to/idf/file")
    // fmt.Println(strings.Join(analyse.ExtractTags(sentence, 20), "/ "))
    // fmt.Printf("【全模式】: %s\n\n", strings.Join(jiebago.Cut(string(data2014), true, true), "/ "))
    // fmt.Printf("【精确模式】: %s\n\n", strings.Join(jiebago.Cut(string(data2014), false, true), "/ "))
    v1 := getTagsAsVector(data2013)
    v2 := getTagsAsVector(data2014)
    v3 := getTagsAsVector(data2015)

    printAll(v1,2013)
    printAll(v2,2014)
    printAll(v3,2015)

//lazy so just re-use the analyse fx
    fmt.Println("2014年 關鍵字 Top",topKWord)
    printTopK(analyse.ExtractTags(string(data2013), topKWord),v1,topKWord)
    fmt.Println("2014年 關鍵字 Top",topKWord)
    printTopK(analyse.ExtractTags(string(data2014), topKWord),v1,topKWord)
    fmt.Println("2015年 關鍵字 Top",topKWord)
    printTopK(analyse.ExtractTags(string(data2015), topKWord),v2,topKWord)


    similarity1415 :=  getCosine(v2,v3)
    fmt.Println("Cosine Similarity 2014 vs 2015:",similarity1415)

    similarity1315 :=  getCosine(v1,v3)
    fmt.Println("Cosine Similarity 2013 vs 2015:",similarity1315)
    // Filter stop words again from results
  }
