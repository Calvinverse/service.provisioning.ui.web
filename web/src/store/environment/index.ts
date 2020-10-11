import {
  VuexModule,
  Module,
  Mutation,
  Action
} from 'vuex-module-decorators'
import { AuthenticationService } from '../../services/AuthenticationService'
import { EnvironmentService } from '../../services/EnvironmentService'

const environmentService = new EnvironmentService(AuthenticationService.getInstance())

export class Environment {
  id = ''
  name = ''
  description = ''
  createdOn: Date = new Date()
  destroyBy: Date = new Date()
  status = ''
  resources: string[] = []// ID of resources
  tags: string[] = [] // ID of tags
  version = ''
}

@Module({ namespaced: true })
export class Environments extends VuexModule {
  error = false
  private items: Environment[] = []

  // getters

  get environment () {
    return (environmentID: string) => {
      const result = this.items.find(x => x.id === environmentID)
      return result
    }
  }

  // Actions

  @Action
  public async get (environmentID: string) {
    console.log('DEBUG: Store login called')

    // If not in the cache then get it
    const user = await environmentService.get(environmentID) // What if error
    this.context.commit('updateUser', user)
    this.context.commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  }

  @Action
  public async getAll () {
    console.log('DEBUG: Store login called')
    const user = await environmentService.getAll() // What if error
    this.context.commit('updateUser', user)
    this.context.commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  }

  @Action
  public async delete (environmentID: string) {
    await environmentService.delete(environmentID) // What if error
    this.context.commit('updateUser', {})
    this.context.commit('updateIsAuthenticated', false)
  }

  // Mutations
}
