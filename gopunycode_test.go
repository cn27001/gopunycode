package gopunycode

import "testing"

func TestEncode(t *testing.T) {
    domain, err := ToASCII("домены.рф")
    if err != nil {
        t.Errorf("error: %s", err)
    }
    t.Logf("domain: %s", domain)
}