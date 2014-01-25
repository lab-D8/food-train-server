package token

import (
	"testing"
	"github.com/iph/catan/util"
)

func TestSHA512(t *testing.T){
	var testStr = util.GenerateSHA512Hash("testing")
	var testing_hash = "521b9ccefbcd14d179e7a1bb877752870a6d620938b28a66a107eac6e6805b9d0989f45b5730508041aa5e710847d439ea74cd312c9355f1f2dae08d40e41d50"
	if testing_hash != testStr {
		t.Errorf("Something is wrong with the hashing function..")
	}
	
}
