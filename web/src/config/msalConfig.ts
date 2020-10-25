export class MsalConfig {
  clientID = '3a287c7c-6b20-4118-9c4f-1f8b8bb622e7'
  redirectUrl: URL = new URL('http://localhost:8080')
  authority = 'https://calvinverse.b2clogin.com/calvinverse.onmicrosoft.com/B2C_1_user_signin_signup'
  scopes: string[] = ['openid', 'profile']
  graphEndpoint: URL = new URL('https://graph.microsoft.com/v1.0/me')
}
