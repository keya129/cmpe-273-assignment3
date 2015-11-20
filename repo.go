package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
  "log"
	"io/ioutil"
	"net/http"
	"encoding/json"
  "strings"
  //"strconv"
  //"github.com/gorilla/mux"
  //"encoding/binary"
  //  "bytes"

  "math/rand"
   uber "github.com/r-medina/go-uber"

)
var clientReq *uber.Client= uber.NewClient("As8XPWG3ReL46A388bynORIEB7cbmQSKVFiZfx1U")

type msg struct {
  Pid        bson.ObjectId  `bson:"_id"`
  Id        int       `bson:"id"`
  Name      string    `bson:"name"`
  Address   string    `bson:"address"`
  City      string    `bson:"city"`
  State     string    `bson:"state"`
  Zip       string        `bson:"zip"`
  Coordinate     Coordinate  `bson:"coordinate"`
}
type tripmsg struct {
  Pid        bson.ObjectId  `bson:"_id"`
  Id          int       `bson:"id"`
  Status      string    `bson:"status"`
  Starting_from_location_id   string    `bson:"starting_from_location_id"`
  Next_destination_location_id string   `bson:"next_destination_location_id"`
  Best_route_location_ids  []string    `bson:"best_route_location_ids"`
  Total_uber_costs  int    `bson:"total_uber_costs"`
  Total_uber_duration float64 `bson:"total_uber_duration"`
  Total_distance  float64 `bson:"total_distance"`
  Uber_wait_time_eta int `bson:"eta"`
  Count int `bson:"count"`
}

func RepoAddLocation(l Location) msg{
  //uri := os.Getenv("MONGOHQ_URL")
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  collection := sess.DB("trip-planner").C("locations")
  pid:=bson.NewObjectId()
  var randno int
  randno=rand.Intn(1000)+rand.Intn(10000)
  doc := msg{Pid: pid,Id:randno,Name: l.Name,Address:l.Address,City:l.City,State:l.State,Zip:l.Zip,Coordinate:l.Coordinate}

  err = collection.Insert(doc)
  if err != nil {
    fmt.Printf("Can't insert document: %v\n", err)
    os.Exit(1)
  }
return doc
}
func RepoAddTrip(l TripInter) tripmsg{
  //uri := os.Getenv("MONGOHQ_URL")
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  collection := sess.DB("trip-planner").C("trips")
  pid:=bson.NewObjectId()
  doc := tripmsg{Pid: pid,Id:l.Id,Status: l.Status,Starting_from_location_id:l.Starting_from_location_id,Next_destination_location_id:"0",Best_route_location_ids:l.Best_route_location_ids,Total_uber_costs:l.Total_uber_costs,Total_uber_duration:l.Total_uber_duration,Total_distance:l.Total_distance,Uber_wait_time_eta:0,Count:0}

  err = collection.Insert(doc)
  if err != nil {
    fmt.Printf("Can't insert document: %v\n", err)
    os.Exit(1)
  }
return doc
}
func RepoShowLocation(l int) Location{
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  var updatedmsg Location
  err = sess.DB("trip-planner").C("locations").Find(bson.M{"id": l}).One(&updatedmsg)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }

return updatedmsg
}
func RepoShowTrip(l int) TripInter{
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})

  var updatedmsg Trip
  err = sess.DB("trip-planner").C("trips").Find(bson.M{"id": l}).One(&updatedmsg)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }
var updatedmsg2 TripInter
updatedmsg2.Id=updatedmsg.Id
updatedmsg2.Status=updatedmsg.Status
updatedmsg2.Starting_from_location_id=updatedmsg.Starting_from_location_id
updatedmsg2.Best_route_location_ids=updatedmsg.Best_route_location_ids
updatedmsg2.Total_uber_costs=updatedmsg.Total_uber_costs
updatedmsg2.Total_uber_duration=updatedmsg.Total_uber_duration
updatedmsg2.Total_distance=updatedmsg.Total_distance


return updatedmsg2
}
func RepoRemoveLocation(l int){
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})


  err = sess.DB("trip-planner").C("locations").Remove(bson.M{"id": l})
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }

}
func RepoUpdateLocation(l int,k Location) Location{
  uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})
  var updatedmsg Location
  var nupdatedmsg Location

  err = sess.DB("trip-planner").C("locations").Find(bson.M{"id": l}).One(&updatedmsg)
nupdatedmsg.Id=updatedmsg.Id
nupdatedmsg.Name=k.Name
nupdatedmsg.City=k.City
nupdatedmsg.State=k.State
nupdatedmsg.Zip=k.Zip
nupdatedmsg.Coordinate=updatedmsg.Coordinate
nupdatedmsg.Address=updatedmsg.Address


  err = sess.DB("trip-planner").C("locations").Update(updatedmsg,nupdatedmsg)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }
return nupdatedmsg
}

func Example1(l int, New1 Trip){

uri:="mongodb://keya:123@ds045064.mongolab.com:45064/trip-planner"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }

  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()

  sess.SetSafe(&mgo.Safe{})
  //var Old Trip
  //err = sess.DB("trip-planner").C("trips").Find(bson.M{"id": l}).One(&Old)

fmt.Println("here")
//.Println(Old)
fmt.Println(New1)
colQuerier := bson.M{"id":l}
  change := bson.M{"$set": bson.M{"eta":New1.Uber_wait_time_eta,"count":New1.Count}}
  err = sess.DB("trip-planner").C("trips").Update(colQuerier, change)
  //err = sess.DB("trip-planner").C("trips").Update(Old,New)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }





}
 
func RepoUberFindPD(c1 Coordinate,c2 Coordinate) PriceDist{
  k:=0
  lat1:=c1.Lat
  lng1:=c1.Lng
  lat2:=c2.Lat
  lng2:=c2.Lng
  data := make([]Price, 100)

  client := uber.NewClient("As8XPWG3ReL46A388bynORIEB7cbmQSKVFiZfx1U")
  prices, err := client.GetPrices(lat1,lng1,lat2,lng2)
if err != nil {
    fmt.Println(err)
        fmt.Println(c1)
                fmt.Println(c2)

        fmt.Println("above this")


} else {
    for _, price := range prices {
            var temp Price
            var temp2=*price
            temp.Estimate=temp2.Estimate
            temp.LowEstimate=temp2.LowEstimate
            temp.HighEstimate=temp2.HighEstimate
            temp.Duration=temp2.Duration
            temp.Distance=temp2.Distance
            //fmt.Println(temp)
            data[k]=temp
            //fmt.Println("k","%d",k)
            k++ 
            //fmt.Println(*price)
    }
}
var pd PriceDist
pd.Price=data[0].LowEstimate
pd.Distance=data[0].Distance
pd.Duration=data[0].Duration
return pd
}
func QueryGMaps(l Location) Coordinate{
  var err error
  sym:=strings.Split(l.Address," ")
  var queryString string
  queryString=sym[0]
  for k:=1;k<len(sym);k++{
queryString=queryString+"+"+sym[k]
}
symmore:=strings.Split(l.City," ")
for k:=1;k<len(symmore);k++{
queryString=queryString+"+"+symmore[k]
}
queryString=queryString+"+"+l.State
url:="http://maps.google.com/maps/api/geocode/json?address="+queryString+"&sensor=false"
fmt.Println(url)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  	log.Fatal(err)
  }
  req.SetBasicAuth("<token>", "x-oauth-basic")

  client := http.Client{}
  res, err := client.Do(req)
  if err != nil {
  	log.Fatal(err)
  }

  log.Println("StatusCode:", res.StatusCode)
  var dat map[string]interface{}
  // read body
  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
  	log.Fatal(err)
  }

  err = json.Unmarshal(body, &dat);if err != nil {
  panic(err)
  }
//  fmt.Println(dat)

  v:=dat["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"]

  //v:=dat.results["geometry"].(map[string]interface{})["location"]
  //fmt.Println(v)
  var nCord Coordinate
  nCord.Lat=v.(map[string]interface{})["lat"].(float64)
  nCord.Lng=v.(map[string]interface{})["lng"].(float64)

return nCord
}
func getCoordinates(l string) Coordinate{
url:="http://localhost:8081/location/"+l
fmt.Println(url)
var resCoord Location
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }
  req.SetBasicAuth("<token>", "x-oauth-basic")

  client := http.Client{}
  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(body, &resCoord);if err != nil {
  panic(err)
  }
  sendCord:=resCoord.Coordinate
  return sendCord
}
func getBestRoute(c1 Coordinate, c []Coordinate) TripResp{
startmatrix := make([]PriceDist,len(c))
finaldistmatrix := make([]float64,100)
finaldurationmatrix := make([]float64,100)
finalpricematrix := make([]int,100)
finalindexmatrix := make([]int,100)

for k:=0;k<len(c);k++{
      startmatrix[k]=RepoUberFindPD(c1,c[k])
      fmt.Print(startmatrix[k])

}
fmt.Println("\n")

matrix := make([][]PriceDist, len(c)) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range matrix {
  matrix[i] = make([]PriceDist, len(c))
}
for k:=0;k<len(c);k++{
  //fmt.Println(c[k])
  for j:=0;j<len(c);j++{
      matrix[k][j]=RepoUberFindPD(c[k],c[j])
      fmt.Print(matrix[k][j])
  }
fmt.Println("\n")
}
min:=0
minDist:=startmatrix[0].Distance
for k:=0;k<len(startmatrix);k++{
if startmatrix[k].Distance<minDist {
minDist=startmatrix[k].Distance
min=k
}
}
//finalindexmatrix[0]=100000

finalindexmatrix[0]=min
finaldistmatrix[0]=startmatrix[min].Distance
finalpricematrix[0]=startmatrix[min].Price
finaldurationmatrix[0]=startmatrix[min].Duration

//fmt.Println(finalindexmatrix[1])
m:=0
k:=0
//minPrice:=0
//minDuration:=0.0

//someflag:=0
for counter:=0;m<len(c)+1;counter++{
    //fmt.Println("min")

  //fmt.Println(min)
      //fmt.Println(m)
    //fmt.Println(m)

minDist=matrix[min][0].Distance
if(finalindexmatrix[m]==0){
     minDist=matrix[min][1].Distance
 
}

for k=0;k<len(c);k++{
      //fmt.Println("minDist")


//fmt.Println(minDist)
//fmt.Println(finalindexmatrix[m])
flag:=0
for j:=0;j<=m;j++{
    //fmt.Println("bhai")

  //fmt.Print(finalindexmatrix[j])
//finalindexmatrix[j]
if(k==finalindexmatrix[j]){
  flag=1
  }
}
//fmt.Println("\n")

if(flag==0){
         // fmt.Println("flag0")

if (matrix[min][k].Distance<=minDist) {
        //fmt.Println("under if")

      //fmt.Println(matrix[min][k].Distance)

minDist=matrix[min][k].Distance
min=k

//minPrice=matrix[min][k].Price
//minDuration=matrix[min][k].Duration


    //fmt.Println(minDist)

}
}
}
m++
//fmt.Println(minDist)
//fmt.Println(m)
finalindexmatrix[m]=min
//finaldistmatrix[m]=minDist
//finalpricematrix[m]=minPrice
//finaldurationmatrix[m]=minDuration

}
var min2 float64
var min3 int
var min4 float64
var sumDist float64
var sumPrice int
var sumDuration float64
for k:=0;k<m-1;k++{
  fmt.Print(finalindexmatrix[k])
  if(finalindexmatrix[k]==0){
    min2=matrix[finalindexmatrix[k]][1].Distance
}else{
    min2=matrix[finalindexmatrix[k]][0].Distance
}
  for j:=0;j<len(c);j++{
    flag:=0
for jk:=0;jk<=k;jk++{
    //fmt.Print("index ")

  //fmt.Println(finalindexmatrix[jk])
  //fmt.Print("j ")

  //fmt.Print(j)

//finalindexmatrix[j]
if(j==finalindexmatrix[jk]){
  flag=1
  }
    //fmt.Print("flag ")

//fmt.Println(flag)

}

  if (min2>matrix[finalindexmatrix[k]][j].Distance && j!=finalindexmatrix[k] && flag==0){
  min2=matrix[finalindexmatrix[k]][j].Distance
  //min3=matrix[finalindexmatrix[k]][j].Price
  //min4=matrix[finalindexmatrix[k]][j].Duration

  } 
  } 
    fmt.Print("min2 ")

  fmt.Println(min2)
  sumDist+=min2
  sumPrice+=min3
  sumDuration+=min4

  finaldistmatrix[k]=min2
  finalpricematrix[k]=min3
  finaldurationmatrix[k]=min4
}
lastRide:=RepoUberFindPD(c[finalindexmatrix[m-1]],c1)
  sumDist+=lastRide.Distance
  sumDuration+=lastRide.Duration
  sumPrice+=lastRide.Price
  //fmt.Println(finaldistmatrix[k])
  //fmt.Println(finaldurationmatrix[k])
var lastResp TripResp
lastResp.Index=finalindexmatrix
lastResp.SumDistance=sumDist
lastResp.SumDuration=sumDuration
lastResp.SumPrice=sumPrice
lastResp.Count=m-1
return lastResp

}
