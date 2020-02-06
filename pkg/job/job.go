// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/service"
	"log"
)

func ClusterInstall(cluster *db.FateCluster) *db.Job {
	job := db.NewJob("ClusterInstall", "")

	// save job to db
	uuid, err := db.Save(job)
	if err != nil {
		log.Println(err)
	}
	job.Uuid = uuid

	go func() {
		err := install(cluster)
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Println("error install cluster:", err)
		} else {
			job.Status = db.Success_j
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Println(err)
		}
	}()

	return job
}

func ClusterUpdate(cluster *db.FateCluster) *db.Job {

	job := db.NewJob("ClusterInstall", "")
	// save job to db
	uuid, err := db.Save(job)
	if err != nil {
		log.Println(err)
	}
	job.Uuid = uuid

	go func() {
		err := upgrade(cluster)
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Println("error upgrade cluster:", err)
		} else {
			job.Status = db.Success_j
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Println(err)
		}
	}()

	return job
}

func ClusterDelete(clusterId string) *db.Job {

	job := db.NewJob("ClusterDelete", "")
	// save job to db
	uuid, err := db.Save(job)
	if err != nil {
		log.Println(err)
	}
	job.Uuid = uuid

	go func() {
		cluster := new(db.FateCluster)
		result, err := db.FindByUUID(cluster, clusterId)
		if err != nil {
			log.Println("error upgrade cluster:", err)
			job.Result = err.Error()
			job.Status = db.Failed_j
			err = db.UpdateByUUID(job, job.Uuid)
			return
		}

		err = uninstall(result.(*db.FateCluster))
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Println("error upgrade cluster:", err)
		} else {
			job.Status = db.Success_j
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Println(err)
		}
	}()

	return job
}

func install(fc *db.FateCluster) error {

	_, err := service.Install(fc.NameSpaces, fc.Name, fc.Version)
	if err != nil {
		return err
	}

	return nil
}

func upgrade(fc *db.FateCluster) error {

	err := service.Upgrade(fc.NameSpaces, fc.Name, fc.Version)
	return err

}
func uninstall(fc *db.FateCluster) error {

	_, err := service.Delete(fc.NameSpaces, fc.Name)

	return err
}
