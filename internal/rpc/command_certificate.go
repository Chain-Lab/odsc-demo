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

}

// subCertificateVerify 命令 "app certificate verify" 的执行入口
func subCertificateVerify(params *Params, resp *CommandRespond) {

}

// subCertificateModify 命令 "app certificate modify" 的执行入口
func subCertificateModify(params *Params, resp *CommandRespond) {

}
