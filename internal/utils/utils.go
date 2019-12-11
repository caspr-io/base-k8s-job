package utils

import (
	"os"
)

func StdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return stat.Mode()&os.ModeCharDevice == 0
}

// func createURL(appInstance *db.Instance) string {
// 	urlTemplate := appInstance.ApplicationObj.URLPattern
// 	// TODO: Pass variables from tenant-service to application-service
// 	replacer := strings.NewReplacer(
// 		"<tenant>", appInstance.Subscription)
// 	// replacer := strings.NewReplacer(
// 	// 	"<vendor>", appInstance.ApplicationDefinition.VendorName,
// 	// 	"<application>", appInstance.ApplicationDefinition.ApplicationName,
// 	// 	"<tenant>", appInstance.TenantName)

// 	return replacer.Replace(urlTemplate)
// }
