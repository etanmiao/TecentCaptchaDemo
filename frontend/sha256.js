export function JsonSort256(jsonDataPost=null) {
    // 第一步
    // HTTP 请求方法（GET、POST ）本示例中为 GET；
    let HTTPRequestMethod = 'GET';
    // URI 参数，API 3.0 固定为正斜杠（/）
    let CanonicalURI = '/';
    // CanonicalQueryString：发起 HTTP 请求 URL 中的查询字符串，对于 POST 请求，固定为空字符串，对于 GET 请求，则为 URL 中问号（?）后面的字符串内容，本示例取值为：Limit=10&Offset=0。注意：CanonicalQueryString 需要经过 URL 编码。
    let CanonicalQueryString = 'Limit=10&Offset=0';
    // 参与签名的头部信息，至少包含 host 和 content-type 两个头部，也可加入自定义的头部参与签名以提高自身请求的唯一性和安全性。拼接规则：1）头部 key 和 value 统一转成小写，并去掉首尾空格，按照 key:value\n 格式拼接；2）多个头部，按照头部 key（小写）的字典排序进行拼接。此例中为：content-type:application/x-www-form-urlencoded\nhost:cvm.tencentcloudapi.com\n
    let CanonicalHeaders = 'content-type:application/x-www-form-urlencoded\nhost:cvm.tencentcloudapi.com\n';
    // 参与签名的头部信息，说明此次请求有哪些头部参与了签名，和 CanonicalHeaders 包含的头部内容是一一对应的。content-type 和 host 为必选头部。拼接规则：1）头部 key 统一转成小写；2）多个头部 key（小写）按照字典排序进行拼接，并且以分号（;）分隔。此例中为：content-type;host
    let SignedHeaders = 'content-type;host';
    // 请求正文的哈希值，计算方法为 Lowercase(HexEncode(Hash.SHA256(RequestPayload)))，对 HTTP 请求整个正文 payload 做 SHA256 哈希，然后十六进制编码，最后编码串转换成小写字母。注意：对于 GET 请求，RequestPayload 固定为空字符串，对于 POST 请求，RequestPayload 即为 HTTP 请求正文 payload。
    let HashedRequestPayload = sha256(encodeURI('')).toLowerCase();
    // 拼接规范请求串
    let CanonicalRequest =
        HTTPRequestMethod + '\n' +
        CanonicalURI + '\n' +
        CanonicalQueryString + '\n' +
        CanonicalHeaders + '\n' +
        SignedHeaders + '\n' +
        HashedRequestPayload;
    console.log('完成第一步',CanonicalRequest);

    // 第二步
    // 签名算法，目前固定为 TC3-HMAC-SHA256；
    let Algorithm = 'TC3-HMAC-SHA256';
    // 请求时间戳，即请求头部的 X-TC-Timestamp 取值，如上示例请求为 1539084154；
    let RequestTimestamp = '1539084154';
    // 凭证范围，格式为 Date/service/tc3_request，包含日期、所请求的服务和终止字符串（tc3_request）。Date 为 UTC 标准时间的日期，取值需要和公共参数 X-TC-Timestamp 换算的 UTC 标准时间日期一致；service 为产品名，必须与调用的产品域名一致，例如 cvm。如上示例请求，取值为 2018-10-09/cvm/tc3_request；
    let CredentialScope = '2018-10-09/cvm/tc3_request';
    // 前述步骤拼接所得规范请求串的哈希值，计算方法为 Lowercase(HexEncode(Hash.SHA256(CanonicalRequest)))。
    let HashedCanonicalRequest = sha256(CanonicalRequest).toLowerCase();
    let StringToSign =
        Algorithm + '\n' +
        RequestTimestamp + '\n' +
        CredentialScope + '\n' +
        HashedCanonicalRequest;
    console.log('完成第二步',StringToSign);

    // 第三步
    // 原始的 SecretKey；
    let SecretKey = "Gu5t9xGARNpq86cd98joQYCN3EXAMPLE";
    // Date：即 Credential 中的 Date 字段信息，如上示例，为2018-10-09；
    let SecretDate = HmacSHA256('2018-10-09',"TC3" + SecretKey);
    // Service：即 Credential 中的 Service 字段信息，如上示例，为 cvm；
    let SecretService = HmacSHA256('cvm',SecretDate);
    // SecretSigning：即以上计算得到的派生签名密钥；
    let SecretSigning = HmacSHA256("tc3_request",SecretService);
    // StringToSign：即步骤2计算得到的待签名字符串；
    let Signature = Hex.stringify(HmacSHA256(StringToSign,SecretSigning));
    console.log('完成第三步',Signature);

    // 第四步
    // 签名方法，固定为 TC3-HMAC-SHA256；
    // let Algorithm = 'TC3-HMAC-SHA256';
    // SecretId：密钥对中的 SecretId；
    let SecretId = 'AKIDz8krbsJ5yKBZQpn74WFkmLPx3EXAMPLE';
    // let Authorization =
    //   Algorithm + ' ' +
    //   'Credential=' + SecretId + '/' + CredentialScope + ', ' +
    //   'SignedHeaders=' + SignedHeaders + ', '
    //   'Signature=' + Signature;
    let Authorization =
        Algorithm + ' ' +
        'Credential=' + SecretId + '/' + CredentialScope + ', ' +
        'SignedHeaders=' + SignedHeaders + ', '+
        'Signature=' + Signature
    console.log('完成第四步',Authorization);
}