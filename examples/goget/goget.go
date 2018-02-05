package main

import (
	"github.com/reguero/go-snmplib"
)

func main() {
	snmplib.DoGetTestV3("137.138.31.79", ".1.3.6.1.4.1.96.255.1", "loadbalancing", "MD5", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9", "DES", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9")
	snmplib.DoGetTestV3("137.138.31.79", ".1.3.6.1.4.1.96.255.1", "loadbalancing", "MD5", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9", "NOPRIV", "")
	snmplib.DoGetTestV3("188.184.92.88", ".1.3.6.1.4.1.96.255.1", "loadbalancing", "MD5", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9", "DES", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9")
	snmplib.DoGetTestV3("188.184.92.88", ".1.3.6.1.4.1.96.255.1", "loadbalancing", "MD5", "OidLvSu8amNMRz258Udy74tO60p47n0RA4RzaT3j2hhnJkEQg9", "NOPRIV", "")
}
