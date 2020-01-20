package db

import (
	"testing"
)


func TestAddCluster(t *testing.T) {
	h := NewHelmChart("TestChart", "Test", "Test")
	c := NewCluster("Cluster1", "Test", "1.2", "9999", *h, *NewBaseObject())
	result, error := c.Save("cluster", c)
	if error == nil {
		t.Log(result)
	}
}

func TestFindCluster(t *testing.T) {
	b := NewBaseObject()
	c := &Cluster{}
	results, _ := b.Find("cluster", *c)
	t.Log(results)
}