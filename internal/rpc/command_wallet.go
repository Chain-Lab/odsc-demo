package rpc

func WalletProxy(subCommand string, params *Params, resp *CommandRespond) {
	switch subCommand {
	// 上传已加密的钱包
	case "upload":
		subCertificatePublish(params, resp)
	}
}

func subWalletUpload(params *Params, resp *CommandRespond) {

}
