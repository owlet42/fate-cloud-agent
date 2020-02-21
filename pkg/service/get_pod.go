package service

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetPod() {
	//EnvCs.Lock()
	//err := os.Setenv("HELM_NAMESPACE", "")
	//if err!=nil{
	//	panic(err)
	//}
	//settings := cli.New()
	//EnvCs.Unlock()
	//
	//cfg := new(action.Configuration)
	////out := os.Stdout
	////namespace := ""
	//if err := cfg.Init(settings.RESTClientGetter(), "", os.Getenv("HELM_DRIVER"), debug); err != nil {
	//	fmt.Println(err)
	//}
	//if err := cfg.KubeClient.IsReachable(); err != nil {
	//	fmt.Println(err)
	//}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
	}
	config.String()

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}

	//
	//	configFlags := kube.GetConfig("", "", "")
	//config, _ := configFlags.ToRESTConfig()
	////config, _ := settings.RESTClientGetter().ToRESTConfig()
	//
	//clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err)
	}

	pods, err := clientset.CoreV1().Pods("fate-9999").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	fmt.Printf("%-15s %-35s %-15s\n", "Namespace", "Name", "Status")
	for _, v := range pods.Items {
		fmt.Printf("%-15s %-35s %-15s\n", v.Namespace, v.Name, v.Status.Phase)
	}
}
