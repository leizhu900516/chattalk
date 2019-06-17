package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	unKnowLocation = "未知"
	turl  = "http://ip.taobao.com/service/getIpInfo.php?ip="
)
type gpsinfo struct {
	Code 		int  `json:"code"`
	Data 		Data `json:"data"`
}
// example
// "data":{"area":"","region":"河南","city":"郑州","county":"XX","isp":"电信","country_id":"CN","area_id":"", "region_id":"410000","city_id":"410100","county_id":"xx","isp_id":"100017"}
type Data struct {
	Ip 			string `json:"ip"`
	Country 	string `json:"country"`
	Area 		string `json:"area"`
	Region 		string `json:"region"`
	City 		string `json:"city"`
	County 		string `json:"county"`
	Isp 		string `json:"isp"`
	Country_id 	string `json:"country_id"`
	Area_id 	string `json:"area_id"`
	Region_id 	string `json:"region_id"`
	City_id 	string `json:"city_id"`
	County_id 	string `json:"county_id"`
	Isp_id  	string `json:"isp_id"`

}
//定位ip的归属地
//采用淘宝的ip接口

func IpLocation(ip string) (string,error){
	fmt.Println("ip=",ip)
	if resp,err := http.Get(turl+ip);err!=nil{
		fmt.Println("1",err)
		return unKnowLocation,err
	}else {
		defer resp.Body.Close()
		if body,err := ioutil.ReadAll(resp.Body);err!=nil{
			fmt.Println("2",err)
			return unKnowLocation,err
		}else{
			var gpsdata gpsinfo
			if err := json.Unmarshal([]byte(body),&gpsdata);err!=nil{
				fmt.Println("3",err)
				return unKnowLocation,err
			}else{
				reigncity := fmt.Sprintf("%s %s",gpsdata.Data.Region,gpsdata.Data.City)
				return reigncity,nil
			}
		}

	}

}
