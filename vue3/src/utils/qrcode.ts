import QrcodeDecoder from 'qrcode-decoder'
import QrReader from 'qrcode-reader'

export function ParseQrPictureFile(file, callback) {
    let result

    const reader = new FileReader()
    reader.readAsDataURL(file.originFileObj);
    reader.addEventListener('load', () => {
        const qd = new QrcodeDecoder()
        const qr = new QrReader()
        qr.callback = function(error, result) {
            if(error) {
              console.log(error)
              return;
            }
            console.log(result)
        }
        qr.decode(reader.result)

        qd.decodeFromImage(reader.result).then((res) => {
            if (res && !isEmpty(res.data)) {
                result = parseAuthUri(reader.result, res.data)
            } else {
                result = {"data":{},"error":"Invalid QRCode file"}
            }

            callback(result)
        })
    })
}

function parseAuthUri(qrCode, data) {
    const head = 'otpauth://totp/'
    if (data.indexOf(head) === 0) {
        let account = data.substring(data.indexOf(head) + head.length, data.indexOf('?'))
        let issuer = ''
        if (account.indexOf(':') >= 0) {
            issuer = account.substring(0, account.indexOf(':'))
            account = account.substring(account.indexOf(':') + 1)
        }
        const params = data.substring(data.indexOf('?') + 1)
        const paramsArray = params.split('&')
        let secret = ''
        let period = 30
        let digits = 6
        const qrcode = qrCode
        paramsArray.forEach(element => {
            if (element.indexOf('=') < 0) {
                return
            }
            const kv = element.split('=')
            if (kv.length < 2) {
                return
            }

            const key = kv[0]
            const value = kv[1]
            switch (key) {
                case 'secret':
                    secret = value
                    break
                case 'issuer':
                    issuer = value
                    break
                case 'period':
                    period = value
                    break
                case 'digits':
                    digits = value
                    break
            }
        })

        return {"data": {"account": account, "issuer": issuer, "secret": secret, "period": period, "digits": digits, "qrcode": qrcode}, "error": ''}
    }

    return {"data": {}, "error": 'Invalid QRCode file'};
}

function isEmpty(obj) {
    if (typeof obj === 'undefined' || obj == null || obj === '') {
        return true
    } else {
        return false
    }
}