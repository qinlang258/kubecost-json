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
	kubecost_project = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=product&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	//这个business指的是业务线
	kubecost_business = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=department&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	//这个deployment是指的 bigdata与techcenter
	kubecost_department  = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=team&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_namespace   = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=namespace&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_deployment  = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=deployment&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_daemonset   = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=daemonset&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	kubecost_statefulset = "http://kubecost.prod.yunlizhi.net/model/allocation/view?aggregate=statefulset&window=yesterday&shareIdle=true&idle=true&idleByNode=false&includeSharedCostBreakdown=true&shareTenancyCosts=true&shareNamespaces=monitoring%2Ckube-system%2Cmissing-container-metrics%2Ccattle-fleet-system%2Ccattle-impersonation-system%2Ccattle-system&shareCost=NaN&shareSplit=weighted&chartType=costovertime&costUnit=cumulative&step="
	mongodbUrl           = "mongodb://10.1.136.162:27017"
	mongodbUsername      = "app_kubecost"
	mongodbPassword      = "OoYCJfXnQYO8rbfXn1"
)

func getProject(client *mongo.Client) {
	//获取所有app的标签费用
	app_response, err := http.Get(kubecost_project)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer app_response.Body.Close()

	bytes, err := ioutil.ReadAll(app_response.Body)
	if err != nil {
		panic(err.Error())
	}

	var project api.Project
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	project.Date = date
	err = json.Unmarshal(bytes, &project)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取app集合句柄
	appCollection := client.Database("test").Collection("project")
	_, err = appCollection.InsertOne(context.TODO(), project)

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
	departmentCollection := client.Database("test").Collection("department")
	_, err = departmentCollection.InsertOne(context.TODO(), department)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")
}

func getBusiness(client *mongo.Client) {
	//获取所有业务线：willcloud-01的标签费用
	businessResponse, err := http.Get(kubecost_business)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer businessResponse.Body.Close()

	bytes, err := ioutil.ReadAll(businessResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var business api.Business
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	business.Date = date
	err = json.Unmarshal(bytes, &business)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取app集合句柄
	yewuxianCollection := client.Database("test").Collection("business")
	_, err = yewuxianCollection.InsertOne(context.TODO(), business)

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
	appCollection := client.Database("test").Collection("namespace")
	_, err = appCollection.InsertOne(context.TODO(), namespace)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}
	fmt.Println("Document inserted successfully!")
}

func getDeployment(client *mongo.Client) {
	//获取所有app的标签费用
	namespaceResponse, err := http.Get(kubecost_deployment)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer namespaceResponse.Body.Close()

	bytes, err := ioutil.ReadAll(namespaceResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var deployment api.Deployment
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	deployment.Date = date
	err = json.Unmarshal(bytes, &deployment)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取namespace集合句柄
	appCollection := client.Database("test").Collection("deployment")
	_, err = appCollection.InsertOne(context.TODO(), deployment)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")

}

func getDaemonSet(client *mongo.Client) {
	//获取所有app的标签费用
	daemonSetResponse, err := http.Get(kubecost_daemonset)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer daemonSetResponse.Body.Close()

	bytes, err := ioutil.ReadAll(daemonSetResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var daemonset api.DaemonSet
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	daemonset.Date = date
	err = json.Unmarshal(bytes, &daemonset)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取namespace集合句柄
	appCollection := client.Database("test").Collection("daemonset")
	_, err = appCollection.InsertOne(context.TODO(), daemonset)

	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	fmt.Println("Document inserted successfully!")
}

func getStatefulSet(client *mongo.Client) {
	//获取所有app的标签费用
	statsfulSetResponse, err := http.Get(kubecost_statefulset)
	if err != nil {
		fmt.Println("Error while making the request:", err)
		return
	}
	defer statsfulSetResponse.Body.Close()

	bytes, err := ioutil.ReadAll(statsfulSetResponse.Body)
	if err != nil {
		panic(err.Error())
	}

	var statefulSet api.StatefulSet
	currentTime := time.Now()
	date := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Println(date)

	statefulSet.Date = date
	err = json.Unmarshal(bytes, &statefulSet)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		return
	}

	// 获取namespace集合句柄
	appCollection := client.Database("test").Collection("statefulset")
	_, err = appCollection.InsertOne(context.TODO(), statefulSet)

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

	getProject(client)
	getDepartment(client)
	getBusiness(client)
	getNamespace(client)
	getDaemonSet(client)
	getDeployment(client)
	getStatefulSet(client)
}
