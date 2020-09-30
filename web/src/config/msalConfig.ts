import http from '../http-commons'

class MsalConfig {
  clientID: string = '15fa3344-bbaf-4cae-8e69-2279803b4ba4'
  redirectUrl: URL = new URL(window.location.origin)
  authority: URL = new URL('https://login.microsoftonline.com/common')
  graphScopes: string[] = ['user.read']
  graphEndpoint: URL = new URL('https://graph.microsoft.com/v1.0/me')
}

export default new MsalConfig()
