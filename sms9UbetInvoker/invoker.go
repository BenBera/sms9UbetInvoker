package sms9UbetInvoker

import (
	"github.com/leo202465/sms9UbetInvoker/cryptor"
)

// 创建一个请求器
func NewInvoker(_shopId, _aesKey, _md5key, _baseUrl string) Invoker {
	return Invoker{
		aesKey:  _aesKey,
		md5Key:  _md5key,
		baseUrl: _baseUrl,
		shopId:  _shopId,
	}
}

// 生成加密内容
func (this Invoker) Sign(_param map[string]any) string {
	// 截取AESKey的前16个字符
	ivStr := this.aesKey[:16]
	// 将参数序列化为JSON字符串
	jsonStr := cryptor.JsonEncode(_param)
	// 拼接MD5Key到Json的后方，然后整体进行MD5加密，获得MD5签名
	md5Sign := cryptor.MD5([]byte(jsonStr + "Key=" + this.md5Key))
	// 将JSON字符串和MD5签名使用美元符号拼接，然后整体进行AES-CBC加密，加密结果用base64加密
	encodeStr := cryptor.AesCBCEncode(jsonStr+"$"+md5Sign, this.aesKey, ivStr)
	return encodeStr
}
