import {
  VuexModule,
  Module,
  Mutation,
  Action
} from 'vuex-module-decorators'
import { Account } from 'msal'
import {
  Options,
  url
} from 'gravatar'
import { AuthenticationService } from '../../services/AuthenticationService'
import { MsalConfig } from '../../config/msalConfig'

const config = new MsalConfig()
const authenticationService = new AuthenticationService(config.clientID, config.authority, config.scopes)

@Module({ namespaced: true, name: 'profile' })
class Profile extends VuexModule {
  user?: Account;
  error: boolean = false;
  isAuthenticated: boolean = false;

  get fullName (): string {
    const user = this.user
    const firstName = (user && user.idToken && user.idToken.given_name) || ''
    const lastName = (user && user.idToken && user.idToken.family_name) || ''
    return `${firstName} ${lastName}`
  }

  get gravatarImage (): string {
    const user = this.user
    const email = (user && user.idToken && user.idToken.emails[0]) || ''
    return url(email)
  }

  @Action
  public async login () {
    console.log('DEBUG: Store login called')
    const user = await authenticationService.login() // What if error
    this.context.commit('updateUser', user)
    this.context.commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  }

  @Action
  public logout () {
    authenticationService.logout() // What if error
    this.context.commit('updateUser', {})
    this.context.commit('updateIsAuthenticated', false)
  }

  @Mutation
  updateIsAuthenticated (value: boolean) {
    this.isAuthenticated = value
  }

  @Mutation
  updateUser (value: Account) {
    this.user = value
  }
}

export default Profile