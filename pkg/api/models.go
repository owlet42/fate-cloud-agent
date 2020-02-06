package api

type boostrapParties struct {
	partyId   string
	endpoint  string
	partyType string
}
type installCluster struct {
	Name            string
	Namespace       string
	Version         string
	EggNumber       int
	BoostrapParties boostrapParties
}
