import * as localforage from 'localforage'

const localForage = localforage.createInstance({
  name: '安全凭证管理后台',
})

export { localForage }
