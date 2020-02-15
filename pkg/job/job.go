// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/service"
	"github.com/rs/zerolog/log"
	"time"
)

type ClusterArgs struct {
	Name      string
	Namespace string
	Version   string
	Data      []byte
}

func ClusterInstall(clusterArgs *ClusterArgs) *db.Job {

	job := db.NewJob("ClusterInstall", "")

	// save job to db
	_, err := db.Save(job)
	if err != nil {
		log.Err(err).Interface("job", job).Msg("save job error")
	}

	log.Info().Str("jobId", job.Uuid).Msg("create a new job of ClusterInstall")

	go func() {

		//create a cluster use parameter
		cluster := db.NewCluster(clusterArgs.Name, clusterArgs.Namespace, clusterArgs.Version,
			db.ComputingBackend{}, db.Party{})
		job.ClusterId = cluster.Uuid

		err := install(cluster, clusterArgs.Data)
		if err != nil {
			job.Result = err.Error()
			job.Status = db.Failed_j
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("install cluster error")
		} else {
			job.Result = "install success"
			job.Status = db.Success_j
			log.Info().Str("ClusterId", cluster.Uuid).Msg("install cluster success")
		}
		job.EndTime = time.Now().String()
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}

		// todo job start status stop timeout
		//for {
		//	//
		//}
		//todo save cluster to db
		log.Info().Str("jobUuid", job.Uuid).Msg("job run success")
	}()

	return job
}

func ClusterUpdate(cluster *db.Cluster) *db.Job {

	job := db.NewJob("ClusterInstall", "")
	// save job to db
	_, err := db.Save(job)
	if err != nil {
		log.Err(err).Interface("job", job).Msg("save job error")
	}
	log.Info().Str("jobId", job.Uuid).Msg("create a new job of ClusterUpdate")
	go func() {
		err := upgrade(cluster)
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("upgrade cluster error")
		} else {
			job.Status = db.Success_j
			log.Info().Str("ClusterId", cluster.Uuid).Msg("upgrade cluster success")
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}
	}()

	return job
}

func ClusterDelete(clusterId string) *db.Job {

	job := db.NewJob("ClusterDelete", "")
	// save job to db
	_, err := db.Save(job)
	if err != nil {
		log.Err(err).Interface("job", job).Msg("save job error")
	}
	log.Info().Str("jobId", job.Uuid).Msg("create a new job of ClusterDelete")
	go func() {
		cluster := new(db.Cluster)
		result, err := db.FindByUUID(cluster, clusterId)
		if err != nil {
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("Cluster does not exist")
			job.Result = err.Error()
			job.Status = db.Failed_j
			err = db.UpdateByUUID(job, job.Uuid)
			if err != nil {
				log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
			}
			return
		}

		err = uninstall(result.(*db.Cluster))
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("delete cluster error")
		} else {
			job.Status = db.Success_j
			log.Info().Str("ClusterId", cluster.Uuid).Msg("delete cluster success")
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}
	}()

	return job
}

func install(fc *db.Cluster, values []byte) error {
	v := new(service.Value)
	v.Val = values
	v.T = "json"

	_, err := service.Install(fc.NameSpaces, fc.Name, fc.Version, v)
	if err != nil {
		return err
	}

	return nil
}

func upgrade(fc *db.Cluster) error {

	err := service.Upgrade(fc.NameSpaces, fc.Name, fc.Version)
	return err

}
func uninstall(fc *db.Cluster) error {

	_, err := service.Delete(fc.NameSpaces, fc.Name)

	return err
}
