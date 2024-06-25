/**
@Author: wei-g
@Date:   2021/6/23 8:08 下午
@Description:
*/

package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// HmacSha256  使用 HMAC 标准算法 计算 sha256
func HmacSha256(key, data string) string {
	hash := hmac.New(sha256.New, []byte(key)) //创建对应的sha256哈希加密算法
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}
