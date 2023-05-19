import { CommonTableList, CommonTreeSelect } from '@/types/type'

export interface MfaType extends CommonTableList {
    // ID
    id: string
    // 账号ID
    uid: string
    // 提供方
    issuer: string
    // 用户登录账号
    account: string
    // 用户显示昵称
    name: string
    // 邮箱地址
    desc?: string
    // 所属状态是否有效  1是有效 0是失效
    status: 0 | 1
  }

  export interface MfaInfoType extends CommonTableList {
    // ID
    id: string
    // 账号ID
    uid: string
    // 提供方
    issuer: string
    // 位数
    digits: number
    // 时效
    period: number
    // Qr码
    qrcode: string
    // 密钥
    secret: string
    // 用户登录账号
    account: string
    // 用户显示昵称
    name: string
    // 邮箱地址
    desc?: string
    // 所属状态是否有效  1是有效 0是失效
    status: 0 | 1
  }  

 export interface OtpType extends CommonTableList {
   // ID
   id: string
   // 关系ID
   rid: string
   // 提供者
   issuer: string
   // 账号
   account: string
   // 计时
   period: number
   // OTP Token
   token: string
 }