package main

import (
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/strings/slices"
	"log"
	"os"
)

var (

	// outside cluster client
	config, _    = clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientset, _ = kubernetes.NewForConfig(config)
)

func main() {

	// slice with allowed namespaces
	ns := []string{"ose-test-namespace-1", "ose-test-namespace-11", "ose-test-ns", "ose-groups"}

	// get all service accounts
	ListSa, _ := clientset.CoreV1().ServiceAccounts("").List(context.Background(), v1.ListOptions{})

	// get all service accounts and their namespaces
	sl1 := []map[string]string{} // slice for maps namespace and service account
	for _, x := range ListSa.Items {
		if slices.Contains(ns, x.Namespace) {

			M1 := map[string]string{ // add namespace name and service account name to map
				x.Namespace: x.Name}
			sl1 = append(sl1, M1)        // add map to slice
			M1 = make(map[string]string) // clear map M1
		}
	}
	log.Println(sl1)
	// здесь у нас slice каждому ns свой ns
	// [map[ose-groups:default] map[ose-test-namespace-1:default] map[ose-test-namespace-1:test-sa] map[ose-test-namespace-11:default] map[ose-test-ns:default] map[ose-test-ns:ose-sa]]
	M1 := make(map[string][]string) // init empty map
	for _, x := range sl1 {         //iterate over slice which one contain map
		for k, v := range x {
			if _, ok := M1[k]; !ok { // при первой итерации у нас нету ключа
				M1[k] = make([]string, 0) // <<< значит мы создаем срез " make([]string, 0)" в нашей мапе M1 и в key добавляем наше имя namespace
			}
			M1[k] = append(M1[k], v) // здесь мы в нашей мапе в срез добавляем наш service account
		}
	}
	/*
			При второй итерации берем уже след namespace name и если имя ns нас небыло в нашем случае это будет ose-test-namespace-1
			видим что его так же нет в нашей мап, повторяем процедуру, след ns так ose-test-namespace-1 в данном случае у нас уже есть ключ со значемнием
			ose-test-namespace-1 и значит мы не создаем новый срез а сразу делаем append с этим ключем ose-test-namespace-1 в slice и получаем
		    ose-test-namespace-1:[default test-sa] <<<  namespace и его 2 service account и т.д. в итоге получаем то, что нам было необходимо

	*/
	log.Println(M1)
	//  map[ose-groups:[default] ose-test-namespace-1:[default test-sa] ose-test-namespace-11:[default] ose-test-ns:[default ose-sa]]
	sl1 = nil //  set slice to nil to prevent overload
}
