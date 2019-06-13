package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type gpsinfo struct {
	Code int `json:"code"`
	Data Data `json:"data"`
}
//
//"data":{"area":"","region":"河南","city":"郑州","county":"XX","isp":"电信","country_id":"CN","area_id":"",
// "region_id":"410000","city_id":"410100","county_id":"xx","isp_id":"100017"}
type Data struct {
	Ip string `json:"ip"`
	Country string `json:"country"`
	Area string `json:"area"`
	Region string `json:"region"`
	City string `json:"city"`
	County string `json:"county"`
	Isp string `json:"isp"`
	Country_id string `json:"country_id"`
	Area_id string `json:"area_id"`
	Region_id string `json:"region_id"`
	City_id string `json:"city_id"`
	County_id string `json:"county_id"`
	Isp_id  string `json:"isp_id"`

}
//定位ip的归属地
//采用淘宝的ip接口

const turl  = "http://ip.taobao.com/service/getIpInfo.php?ip="
func IpLocation(ip string) (string,error){
	if resp,err := http.Get(turl+ip);err!=nil{
		fmt.Println(err)
		return "",err
	}else {
		defer resp.Body.Close()
		if body,err := ioutil.ReadAll(resp.Body);err!=nil{
			fmt.Println(err)
			return "",err
		}else{
			var gpsdata gpsinfo
			if err := json.Unmarshal(body,&gpsdata);err!=nil{
				fmt.Println(err)
				return "",err
			}else{
				reigncity := fmt.Sprintf("%s %s",gpsdata.Data.Region,gpsdata.Data.City)
				return reigncity,nil
				//fmt.Println(gpsdata)
				//fmt.Println(gpsdata.Code)
				//fmt.Println(gpsdata.Data)
			}
		}

	}

}
