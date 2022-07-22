/*
Given a valid (IPv4) IP address, return a defanged version of that IP address.
A defanged IP address replaces every period "." with "[.]".
*/

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
