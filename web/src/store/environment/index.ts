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
import { EnvironmentService } from '../../services/EnvironmentService'

const environmentService = new EnvironmentService(config.clientID, config.authority, config.scopes)

@Module({ namespaced: true, name: 'environment' })
class Environment extends VuexModule {
  error: boolean = false
  id: string = ''
  name: string = ''
  description: string = ''
  createdOn: Date = new Date()
  destroyBy: Date = new Date()
  status: string = ''
  resources: string[]  = []// ID of resources
  tags: string[] = [] // ID of tags
  version: string = ''

  @Action
  public async get(environmentID: string) {
    console.log('DEBUG: Store login called')
    const user = await environmentService.get(environmentID) // What if error
    this.context.commit('updateUser', user)
    this.context.commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  }

  @Action
  public async getAll() {
    console.log('DEBUG: Store login called')
    const user = await environmentService.getAll() // What if error
    this.context.commit('updateUser', user)
    this.context.commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  }

  @Action
  public async delete() {
    await environmentService.delete() // What if error
    this.context.commit('updateUser', {})
    this.context.commit('updateIsAuthenticated', false)
  }
}
