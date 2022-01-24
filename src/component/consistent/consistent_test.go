package consistent

import (
	"fmt"
	"testing"
)

// spy 2022/1/24

func TestConsistent(t *testing.T) {
	hashring := New()
	hashring.NumberOfReplicas = 200

	hashring.Add("server1")
	hashring.Add("server2")
	hashring.Add("server3")
	hashring.Add("server4")
	hashring.Add("server5")
	hashring.Add("server6")

	node, _ := hashring.Get("abc")
	fmt.Println(node) // server5

	hashring.Remove("server5")

	node, _ = hashring.Get("abc")
	fmt.Println(node) // server2

	node, _ = hashring.Get("abc1")
	fmt.Println(node) //server6
}

func TestConsistentRebuild(t *testing.T) {
	hashring := New()
	hashring.NumberOfReplicas = 200

	hashring.Add("server1")
	hashring.Add("server2")
	hashring.Add("server3")
	hashring.Add("server4")
	hashring.Add("server5")
	hashring.Add("server6")

	node, _ := hashring.Get("abc")
	fmt.Println(node) // server5

	nodes := []string{"server2", "server3", "server7"}
	hashring.Rebuild(nodes) // or hashring.Set(nodes)

	node, _ = hashring.Get("abc")
	fmt.Println(node) // server2
}
