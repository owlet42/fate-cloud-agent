//// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/service"
	"time"
)

type JobType uint32

const (
	INSTALL JobType = 1 + iota
	UNINSTALL
	UPGRADE
	EXEC
)

type status uint32

const (
	Success status = iota
	Fail
	Padding
	Running
)

type Cluster struct {
	Name       string
	Describe   string
	Type       JobType
	CreateTime time.Time
	Status     status
	Id         string // FateCluster id
	MateData   interface{}
	Err        error
	Sub        []interface{}
}

func (c *Cluster) init() string {
	return db.Save("job", c)
}
func (c *Cluster) do(id string) {
	if c.Status != Padding {
		return
	}
	switch c.Type {
	case INSTALL:
		install(c)
	case UPGRADE:
		upgrade(c)
	case UNINSTALL:
		uninstall(c)
	case EXEC:
		exec(c)
	}
	db.Update("job", id, c)
	return
}
func (c Cluster) getSub() []interface{} {
	return c.Sub
}
func (c *Cluster) Run() string { return Run(c) }

// TODO
func exec(j *Cluster) {
	return
}

func install(j *Cluster) {
	f := j.MateData.(db.FateCluster)
	_, err := service.Install(f.NameSpaces, f.Name, f.Chart)
	if err != nil {
		j.Err = err
		j.Status = Fail
		return
	}
	j.Id = db.Save("fate", f)
	j.Status = Success
	return
}
func upgrade(j *Cluster) {
	f := j.MateData.(db.FateCluster)
	err := service.Upgrade(f.NameSpaces, f.Name, f.Chart)
	if err != nil {
		j.Err = err
		j.Status = Fail
		return
	}
	db.Update("fate", j.Id, f)
	j.Status = Success
	return
}
func uninstall(j *Cluster) {
	f := j.MateData.(db.FateCluster)
	_, err := service.Delete(f.NameSpaces, f.Name)
	if err != nil {
		j.Err = err
		j.Status = Fail
		return
	}
	db.Delete("fate", j.Id)
	j.Status = Success
	return
}

func (c *Cluster) Create(t JobType, m interface{}) {
	c.Type = t
	c.MateData = m
}
