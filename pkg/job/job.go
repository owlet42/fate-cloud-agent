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

func ClusterInstall(clusterArgs *ClusterArgs) (*db.Job, error) {
	//if ok := service.IsExited(clusterArgs.Name, clusterArgs.Namespace); !ok {
	//	return nil, errors.New("cluster is exited")
	//}
	job := db.NewJob("ClusterInstall", "")

	//  save job to db
	_, err := db.Save(job)
	if err != nil {
		log.Err(err).Interface("job", job).Msg("save job error")
		return nil, err
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
			log.Debug().Str("ClusterId", cluster.Uuid).Msg("install cluster success")
		}

		job.EndTime = time.Now()
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}

		// todo job start status stop timeout
		//for {
		//	//
		//}
		job.Status = db.Success_j

		//todo save cluster to db
		if job.Status == db.Success_j {
			_, err = db.Save(cluster)
			if err != nil {
				log.Err(err).Interface("cluster", cluster).Msg("save cluster error")
			}
			log.Debug().Str("cluster uuid", cluster.Uuid).Msg("create cluster success")
		}

		log.Debug().Interface("job", job).Msg("job run success")
	}()

	return job, nil
}

// todo status no test
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
		if err != nil {
			job.Result = err.Error()
			job.Status = db.Failed_j
			log.Err(err).Str("ClusterId", cluster.Uuid).Msg("upgrade cluster error")
		} else {
			job.Result = "upgrade success"
			job.Status = db.Success_j
			log.Debug().Str("ClusterId", cluster.Uuid).Msg("upgrade cluster success")
		}

		job.EndTime = time.Now()

		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}

		// todo job start status stop timeout
		//for{}
		job.Status = db.Success_j

		if job.Status == db.Success_j {
			err = db.UpdateByUUID(cluster, cluster.Uuid)
			if err != nil {
				log.Err(err).Interface("cluster", cluster).Msg("save cluster error")
			}
			log.Debug().Str("cluster uuid", cluster.Uuid).Msg("Update cluster success")
		}
		log.Debug().Interface("job", job).Msg("job run success")
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
		var cluster *db.Cluster
		job.Status = db.Running_j

		if job.Status == db.Running_j {
			cluster, err = db.ClusterFindByUUID(clusterId)
			if err != nil {
				log.Err(err).Str("ClusterId", clusterId).Msg("Cluster find by uuid error,")
				job.Result = err.Error()
				job.Status = db.Failed_j
			} else {
				log.Debug().Interface("Cluster", cluster).Msg("find cluster success")
			}
		}

		if job.Status == db.Running_j {
			err = uninstall(cluster)
			if err != nil {
				job.Result = err.Error()
				job.Status = db.Failed_j
				log.Err(err).Str("ClusterId", cluster.Uuid).Msg("helm delete cluster error")
			} else {
				job.Result = "uninstall success"
				job.Status = db.Running_j
				log.Debug().Str("ClusterId", cluster.Uuid).Msg("helm delete cluster success")
			}
		}


		// todo job start status stop timeout
		//for{}
		if job.Status==db.Running_j{
			job.Status = db.Success_j
		}
		job.EndTime = time.Now()
		err = db.UpdateByUUID(job, job.Uuid)
		if err != nil {
			log.Err(err).Str("jobId", job.Uuid).Msg("update job By Uuid error")
		}

		if job.Status == db.Success_j {
			err = db.ClusterDeleteByUUID(clusterId)
			if err != nil {
				log.Err(err).Interface("cluster", cluster).Msg("db delete cluster error")
			}
			log.Debug().Str("clusterUuid", clusterId).Msg("db delete cluster success")
		}
		log.Debug().Interface("job", job).Msg("job run success")
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

	err := service.Upgrade(fc.NameSpaces, fc.Name, fc.Version, &service.Value{
		Val: []byte(fc.Values),
		T:   "json",
	})
	return err

}
func uninstall(fc *db.Cluster) error {

	_, err := service.Delete(fc.NameSpaces, fc.Name)

	return err
}
