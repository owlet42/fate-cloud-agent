package api

type userStatus string

const (
	DEPRECATE userStatus = "deprecate"
	AVAILABLE userStatus = "available"
)

type ServiceStatus string //	[ creating deleting updating available unavailable ]
const (
	CREATING    ServiceStatus = "creating"
	DELETING    ServiceStatus = "deleting"
	UPDATING    ServiceStatus = "updating"
	AVAILABLEs  ServiceStatus = "available"
	UNAVAILABLE ServiceStatus = "unavailable"
)

type Egg struct {
	Id      string        `json:"id"`
	Address string        `json:"address"`
	Status  ServiceStatus `json:"status"`
}

type Eggroll struct {
	Roll struct {
		Id      string        `json:"id"`
		Address string        `json:"address"`
		Status  ServiceStatus `json:"status"`
	} `json:"roll"`
	EggNumbers int   `json:"egg_numbers"`
	Eggs       []Egg `json:"eggs"`
}

type Spark struct{}

type BoostrapParties struct {
	PartyId   string `json:"party_id"`
	Endpoint  string `json:"endpoint"`
	PartyType string `json:"party_type"`
}
type ComputingBackend struct {
	BackendType string  `json:"backend_type"`
	BackendInfo Eggroll `json:"backend_info"`
}

type Party struct {
	PartyId   string `json:"party_id"`
	Endpoint  string `json:"endpoint"`
	PartyType string `json:"party_type"`
}

type JobStatus string

type cluster struct {
	UUID             string           `json:"cluster_id"`
	Name             string           `json:"name"`
	Namespace        string           `json:"namespace"`
	Version          string           `json:"version"`
	Status           ServiceStatus    `json:"status"`
	Backend          ComputingBackend `json:"backend"`
	BootstrapParties BoostrapParties  `json:"bootstrap_parties"`
}

type user struct {
	UUID       string     `json:"uuid"`
	UserName   string     `json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	UserStatus userStatus `json:"userStatus"`
}

type job struct {
	UUID      string    `json:"job_id"`
	Status    JobStatus `json:"status"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Method    string    `json:"method"`
	Creator   string    `json:"creator"`
	SubJobs   []string  `json:"sub_jobs"`
}
