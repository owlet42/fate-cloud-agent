// job
package job

import (
	"fate-cloud-agent/pkg/db"
	"fate-cloud-agent/pkg/service"
	"github.com/rs/zerolog/log"
)

func ClusterInstall(cluster *db.FateCluster, values string) *db.Job {
	job := db.NewJob("ClusterInstall", "")

	// save job to db
	_, err := db.Save(job)
	if err != nil {
		log.Err(err).Interface("job", job).Msg("save job error")
	}

	job.ClusterId = cluster.Uuid

	log.Info().Str("jobId", job.Uuid).Msg("create a new job of ClusterInstall")

	go func() {

		err := install(cluster, values)
		job.Result = err.Error()
		if err != nil {
			job.Status = db.Failed_j
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("install cluster error")
		} else {
			job.Status = db.Success_j
			log.Info().Str("ClusterId", cluster.Uuid).Msg("install cluster success")
		}
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}

		// todo job start status stop timeout
		//for {
		//	//
		//}

	}()

	return job
}

func ClusterUpdate(cluster *db.FateCluster) *db.Job {

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
		cluster := new(db.FateCluster)
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

		err = uninstall(result.(*db.FateCluster))
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

func install(fc *db.FateCluster, values string) error {
	v := new(service.Value)
	v.Val = values
	v.T = "json"

	_, err := service.Install(fc.NameSpaces, fc.Name, fc.Version, v)
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
