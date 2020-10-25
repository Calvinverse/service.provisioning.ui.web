import {
  VuexModule,
  Module,
  Mutation,
  Action
} from 'vuex-module-decorators'
import { AuthenticationService } from '../../services/AuthenticationService'
import { Environment, EnvironmentService } from '../../services/EnvironmentService'

const environmentService = new EnvironmentService(AuthenticationService.getInstance())

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

  get environments (): Environment[] {
    return this.items
  }

  get hasAny (): boolean {
    return this.items.length > 0
  }

  // Actions
  @Action
  public clear () {
    console.log('DEBUG: environment store - [action]clear')
    this.context.commit('updateAllEnvironments', [])
    console.log('DEBUG: environment store - [action]clear - complete')
  }

  @Action
  public async get (environmentID: string) {
    console.log('DEBUG: environment store - [action]get')

    // If not in the cache then get it
    const environment = await environmentService.get(environmentID) // What if error
    this.context.commit('updateEnvironment', environment)
    console.log('DEBUG: environment store - [action]get - complete')
  }

  @Action
  public async getAll () {
    console.log('DEBUG: environment store - [action]getAll')
    const environments = await environmentService.getAll() // What if error
    this.context.commit('updateAllEnvironments', environments)
    console.log('DEBUG: environment store - [action]getAll - complete')
  }

  @Action
  public async delete (environmentID: string) {
    console.log('DEBUG: environment store - [action]delete')
    await environmentService.delete(environmentID) // What if error
    this.context.commit('removeEnvironment', environmentID)
    console.log('DEBUG: environment store - [action]delete - complete')
  }

  // Mutations
  @Mutation
  updateAllEnvironments (environments: Environment[]) {
    this.items = environments
  }

  @Mutation
  updateEnvironment (environment: Environment) {
    const index = this.items.findIndex(x => x.id === environment.id)
    if (index > -1) {
      this.items.splice(index, 1, environment)
    } else {
      this.items.push(environment)
    }
  }

  @Mutation
  removeEnvironment (environmentID: string) {
    const index = this.items.findIndex(x => x.id === environmentID)
    if (index > -1) {
      this.items.splice(index, 1)
    }
  }
}
