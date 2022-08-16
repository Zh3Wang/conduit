package encrypt

import "testing"

func TestHash(t *testing.T) {
	s := "qwe254511"
	ns := Hash(s)
	t.Logf("hash code: %s \n", ns)

	err := Verify(ns, s)
	if err != nil {
		t.Errorf("verify err: %s \n", err)
		return
	}
	t.Logf("verify success \n")
}
