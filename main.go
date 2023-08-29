package main

import (
	"clientgo/api"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	kubecost_app        = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=product&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_yewuxian   = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=department&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_department = "http://kubecost.prod.yunlizhi.net/model/allocation/trends?aggregate=team&window=yesterday&accumulate=false&shareIdle=true&idle=true&idleByNode=false&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative"
	kubecost_namespace  = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=namespace&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	mongodbUrl          = "mongodb://10.1.136.162:27017"
	mongodbUsername     = "app_kubecost"
	mongodbPassword     = "OoYCJfXnQYO8rbfXn1"
)

func getApp(client *mongo.Client) {
	//获取所有app的标签费用
	app_response, err := http.Get(kubecost_app)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer app_response.Body.Close()

	bytes, err := ioutil.ReadAll(app_response.Body)
	if err != nil {
		panic(err.Error())
	}

	var app api.App
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	app.Date = date
	err = json.Unmarshal(bytes, &app)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取app集合句柄
	appCollection := client.Database("kubecost").Collection("app")
	_, err = appCollection.InsertOne(context.TODO(), app)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")

}

func getDepartment(client *mongo.Client) {
	//获取所有zuzhijigou的标签费用
	department_response, err := http.Get(kubecost_department)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer department_response.Body.Close()

	bytes, err := ioutil.ReadAll(department_response.Body)
	if err != nil {
		panic(err.Error())
	}

	var department api.Department
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	department.Date = date
	err = json.Unmarshal(bytes, &department)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取app集合句柄
	departmentCollection := client.Database("kubecost").Collection("department")
	_, err = departmentCollection.InsertOne(context.TODO(), department)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")
}

func getYewuxian(client *mongo.Client) {
	//获取所有业务线：willcloud-01的标签费用
	yewuxianResponse, err := http.Get(kubecost_yewuxian)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer yewuxianResponse.Body.Close()

	bytes, err := ioutil.ReadAll(yewuxianResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var yewuxian api.Yewuxian
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	yewuxian.Date = date
	err = json.Unmarshal(bytes, &yewuxian)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取app集合句柄
	yewuxianCollection := client.Database("kubecost").Collection("yewuxian")
	_, err = yewuxianCollection.InsertOne(context.TODO(), yewuxian)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")
}

func getNamespace(client *mongo.Client) {
	//获取所有app的标签费用
	namespaceResponse, err := http.Get(kubecost_namespace)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer namespaceResponse.Body.Close()

	bytes, err := ioutil.ReadAll(namespaceResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var namespace api.Namespace
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	namespace.Date = date
	err = json.Unmarshal(bytes, &namespace)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取namespace集合句柄
	appCollection := client.Database("kubecost").Collection("namespace")
	_, err = appCollection.InsertOne(context.TODO(), namespace)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")

}

func main() {
	clientOptions := options.Client().ApplyURI(mongodbUrl).
		SetAuth(options.Credential{
			Username: mongodbUsername,
			Password: mongodbPassword,
		})

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// 断开连接
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	getApp(client)
	getDepartment(client)
	getYewuxian(client)
	getNamespace(client)

}
