import { MsalConfig } from '@/config/msalConfig'
import { Constants } from '@/types'
import { Channel, EventBus, Topic } from 'estacion/lib'
import * as Msal from 'msal'
import { EventBusService } from './EventBusService'

const config = new MsalConfig()

export interface AuthenticationServiceDefinition {
  login(): void;
  logout(): void;
  getUser(): Msal.Account;
  isAuthenticated(): boolean;
}

// This should be a singleton because internally Msal.UserAgentApplication keeps state
// which isn't linked to vuex or anything
export class AuthenticationService implements AuthenticationServiceDefinition {
  private static instance: AuthenticationService

  static getInstance (): AuthenticationService {
    if (!AuthenticationService.instance) {
      AuthenticationService.instance = new AuthenticationService(config.clientID, config.authority, config.scopes)
    }

    return AuthenticationService.instance
  }

  private app: Msal.UserAgentApplication
  private scopes: string[]

  private bus: EventBus

  private authenticationChannel: Channel
  private loginTopic: Topic
  private logoutTopic: Topic

  private constructor (clientId: string, authority: string, scopes: string[]) {
    this.app = new Msal.UserAgentApplication(
      {
        auth: {
          clientId: clientId,
          authority: authority,
          validateAuthority: false,
          navigateToLoginRequestUrl: false
        },
        cache: {
          cacheLocation: 'localStorage',
          storeAuthStateInCookie: true
        }
      }
    )

    this.scopes = scopes

    this.bus = EventBusService.getInstance().getBus()
    this.authenticationChannel = this.bus.channel(Constants.authenticationChannel)
    this.loginTopic = this.authenticationChannel.topic(Constants.userLoginTopic)
    this.logoutTopic = this.authenticationChannel.topic(Constants.userLogoutTopic)
  }

  async login (): Promise<boolean> {
    const loginRequest = {
      scopes: this.scopes,
      prompt: 'select_account'
    }

    const accessTokenRequest = {
      scopes: this.scopes
    }

    try {
      const loginResponse = await this.app.loginPopup(loginRequest)
      console.log(`Login was a success ${loginResponse}`)
    } catch (error) {
      console.log(`Login error ${error}`)
      return false
    }

    try {
      const tokenResponse = await this.app.acquireTokenSilent(accessTokenRequest)
      console.log(`Token response acquired silently - ${tokenResponse}`)
    } catch (error) {
      console.log(`Failed to acquire the token silently, using a pop up -- ${error}`)

      try {
        const tokenResponse = await this.app.acquireTokenPopup(accessTokenRequest)
        console.log(`Token response acquired with a pop up - ${tokenResponse}`)
      } catch (errorPopup) {
        console.log(`Error acquiring the popup: ${errorPopup}`)
        return false
      }
    }

    this.loginTopic.emit({})

    return true
  }

  logout (): void {
    this.app.logout()
    this.logoutTopic.emit({})
  }

  getUser (): Msal.Account {
    return this.app.getAccount()
  }

  isAuthenticated (): boolean {
    return this.app !== undefined
  }
}
