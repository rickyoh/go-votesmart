package votesmart

import (
	"testing"
	"log"
)


func TestVoteSmartQuery(t *testing.T) {

	args := []map[string]string{}
	arg := map[string]string{}
	arg["key"] = "lastName"
	arg["value"] = lastName
	args = append(args, arg)

	method := "Candidates.getByLastname"

	// @todo set up test
	// data := env.Votesmart.Query(method, args)
}
