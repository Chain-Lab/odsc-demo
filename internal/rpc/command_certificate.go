package rpc

func CertificateProxy(subCommand string, params *Params, resp *CommandRespond) {
	switch subCommand {
	case "publish":
		subCertificatePublish(params, resp)
	case "verify":
		subCertificateVerify(params, resp)
	case "modify":
		subCertificateModify(params, resp)
	}
}

// subCertificatePublish 命令 "app certificate publish" 的执行入口
func subCertificatePublish(params *Params, resp *CommandRespond) {
	data := (*params)["data"]
}

// subCertificateVerify 命令 "app certificate verify" 的执行入口
func subCertificateVerify(params *Params, resp *CommandRespond) {

}

func subCertificateModify(params *Params, resp *CommandRespond) {

}
