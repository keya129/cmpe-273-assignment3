package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	//"encoding/binary"
	"math/rand"
	"bytes"
	//uber "github.com/r-medina/go-uber"
	"gopkg.in/mgo.v2"
  	"gopkg.in/mgo.v2/bson"
  	"os"
  	//"log"
	

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func LocationShow(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
    todoId := vars["location_id"]
		l, err := strconv.Atoi(todoId)
		if err != nil {
			panic(err)
		}

		var getLoc Location
		getLoc=RepoShowLocation(l)
    //fmt.Fprintln(w, "Todo show:", getLoc)
		res2B, _ :=json.Marshal(getLoc)
		w.Header().Set("Content-Type", "application/json;")
		w.WriteHeader(200)
		fmt.Fprintf(w,"%s",res2B)

}
func TripShow(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoId := vars["trip_id"]
		l, err := strconv.Atoi(todoId)
		if err != nil {
			panic(err)
		}

		var getTrip TripInter
		getTrip=RepoShowTrip(l)
    //fmt.Fprintln(w, "Todo show:", getLoc)
		res2B, _ :=json.Marshal(getTrip)
		w.Header().Set("Content-Type", "application/json;")
		w.WriteHeader(200)
		fmt.Fprintf(w,"%s",res2B)

}

func LocationUpdate(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["location_id"]
	l, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}
	var loc Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loc); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		/*if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}*/
	}
	Nc:=QueryGMaps(loc)

resp := Location{
	Id:loc.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
mess:=	RepoUpdateLocation(l,resp)

	//json.NewEncoder(w).Encode(mess)
	res2B, _ :=json.Marshal(mess)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s",res2B)

}
func LocationRemove(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["location_id"]
	l, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}

	RepoRemoveLocation(l)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(200)

}
func LocationCreate(rw http.ResponseWriter, r *http.Request) {
	//t := RepoCreateTodo(todo)
	var loc Location

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loc); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//rw.WriteHeader(422) // unprocessable entity
		/*if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}*/
	}
	Nc:=QueryGMaps(loc)
	var mess msg

resp := Location{
	Id:mess.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
//js, _ := json.Marshal(resp)
//fmt.Printf("%s", js)
mess=RepoAddLocation(resp)
respnew := Location{
	Id:mess.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
res2B, _ :=json.Marshal(respnew)
rw.Header().Set("Content-Type", "application/json;")
rw.WriteHeader(http.StatusCreated)
fmt.Fprintf(rw,"%s",res2B)

}
func TripFinder(w http.ResponseWriter, r *http.Request){
var reqFinal BestRoutes

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &reqFinal); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
	startId:=reqFinal.Starting_from_location_id
	startC:=getCoordinates(startId)
	//var destCords []Coordinate
	locids:=reqFinal.Location_ids

	destCords := make([]Coordinate, len(locids))

	for k:=0;k<len(locids);k++{
		destCords[k]=getCoordinates(locids[k])
			//fmt.Println(destCords[k])

	}
	lastResp:=getBestRoute(startC,destCords)
	//var newlocids [] string
	newlocids := make([]string,lastResp.Count)

	for k:=0;k<lastResp.Count;k++{
		newlocids[k]=locids[lastResp.Index[k]]
		
	}
	var finalResp TripInter
	finalResp.Id=rand.Intn(1000)+rand.Intn(100)+rand.Intn(10000)
	finalResp.Status="Planning"
	finalResp.Starting_from_location_id=startId
	finalResp.Best_route_location_ids=newlocids
	finalResp.Total_uber_costs=lastResp.SumPrice
	finalResp.Total_uber_duration=lastResp.SumDuration
	finalResp.Total_distance=lastResp.SumDistance
//fmt.Fprintf(w,"%s",finalResp)
RepoAddTrip(finalResp)	
res2B, _ :=json.Marshal(finalResp)
w.Header().Set("Content-Type", "application/json;")
w.WriteHeader(http.StatusCreated)
fmt.Fprintf(w,"%s",res2B)

} 
func TripUpdate(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["trip_id"]
		fmt.Println(todoId)
l, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}

	//cc1:=getCoordinates(todoId)
	
	var GetTrip Trip
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

  //var updatedmsg Trip
  err = sess.DB("trip-planner").C("trips").Find(bson.M{"id": l}).One(&GetTrip)
  if err != nil {
    fmt.Printf("got an error finding a doc %v\n")
    os.Exit(1)
  }

  	fmt.Println("here")
  	  fmt.Println(GetTrip.Starting_from_location_id)

	indx:=GetTrip.Count
	var todoId2 string
	var todoId1 string
if(indx==len(GetTrip.Best_route_location_ids)){
  	todoId2=GetTrip.Starting_from_location_id
	todoId1=GetTrip.Best_route_location_ids[indx-1]
  }else{
  	todoId2=GetTrip.Best_route_location_ids[indx]
  	todoId1=GetTrip.Starting_from_location_id

  }
	cc2:=getCoordinates(todoId2)
	fmt.Println("here2")
	cc1:=getCoordinates(todoId1)

	fmt.Println(cc1.Lat)
	fmt.Println(cc2.Lat)
	fmt.Println(cc1.Lng)
	fmt.Println(cc2.Lng)
	//c1:=string(cc2.Lat)
	c1:=strconv.FormatFloat(cc2.Lat, 'f', 6, 64)
	c2:=strconv.FormatFloat(cc2.Lng, 'f', 6, 64)	
	c3:=strconv.FormatFloat(cc1.Lat, 'f', 6, 64)
	c4:=strconv.FormatFloat(cc1.Lng, 'f', 6, 64)	
	fmt.Println(c1)
		fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	var myurl Buffer
	myurl.Product_id="04a497f5-380d-47f2-bf1b-ad4cfdcb51f2"
	myurl.Start_latitude=cc1.Lat
	myurl.Start_longitude=cc1.Lng
	myurl.End_latitude=cc2.Lat
	myurl.End_longitude=cc2.Lng
	res2B, _ :=json.Marshal(myurl)
	url := "https://sandbox-api.uber.com/v1/requests"
	fmt.Println("URL:>", url)
	var jsonStr = []byte(res2B)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer FNet8JSp0KMMbAMfICAyh1nebRoMAq")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
	panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	var dat EtaResp
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
    err = json.Unmarshal(body, &dat);if err != nil {
  	panic(err)
  }
  tempindx:=GetTrip.Count
  fmt.Println(dat.Eta)
  if(indx!=len(GetTrip.Best_route_location_ids)){
  	indx++
  }else{
  	indx=0
  }
  GetTrip.Count=indx
  GetTrip.Uber_wait_time_eta=dat.Eta
  fmt.Println("l")
  fmt.Println(l)
  Example1(l,GetTrip)
  var FinalResp TripFinal
  FinalResp.Id=GetTrip.Id
  if(tempindx==len(GetTrip.Best_route_location_ids)){
  	FinalResp.Status="Finished"
  }else{
  	FinalResp.Status="Planning"
  }
  FinalResp.Id=GetTrip.Id
  if(tempindx!=len(GetTrip.Best_route_location_ids)){
  FinalResp.Starting_from_location_id=GetTrip.Starting_from_location_id
  }else{
  FinalResp.Starting_from_location_id=GetTrip.Best_route_location_ids[tempindx-1]	
  }
    if(tempindx!=len(GetTrip.Best_route_location_ids)){

  FinalResp.Next_destination_location_id=GetTrip.Best_route_location_ids[tempindx]
  }else{
  FinalResp.Next_destination_location_id=GetTrip.Starting_from_location_id
	
  }
  FinalResp.Best_route_location_ids=GetTrip.Best_route_location_ids
  FinalResp.Total_uber_costs=GetTrip.Total_uber_costs
  FinalResp.Total_uber_duration=GetTrip.Total_uber_duration
  FinalResp.Total_distance=GetTrip.Total_distance
  FinalResp.Uber_wait_time_eta=GetTrip.Uber_wait_time_eta
  //FinalResp.Count=tempindx

 res2BBB, _ :=json.Marshal(FinalResp)
w.Header().Set("Content-Type", "application/json;")
w.WriteHeader(http.StatusCreated)
fmt.Fprintf(w,"%s",res2BBB)
 
}
