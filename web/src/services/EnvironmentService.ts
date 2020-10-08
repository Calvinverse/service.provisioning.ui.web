import {
  EnvironmentDefinition,
  TagDefinition
} from '../types/index'
import { AuthenticationServiceDefinition } from './AuthenticationService'

export class Environment implements EnvironmentDefinition {
  id = ''
  name = ''
  description = ''
  createdOn: Date = new Date()
  destroyBy: Date = new Date()
  resources: string[] = new Array<string>()
  tags: TagDefinition[] = new Array<TagDefinition>()
  version = ''

  status (): string {
    throw new Error('Method not implemented.')
  }

  create (templateID: string): string[] {
    throw new Error('Method not implemented.')
  }

  destroy (resourceID: string): boolean {
    throw new Error('Method not implemented.')
  }
}

export class EnvironmentService {
  private authenticationService: AuthenticationServiceDefinition

  constructor (authenticationService: AuthenticationServiceDefinition) {
    this.authenticationService = authenticationService
  }

  // Create a new environment and return the ID of the environment
  async create (): Promise<string> {
    if (this.authenticationService.isAuthenticated()) {
      return ''
    } else {
      return ''
    }
  }

  async delete (environmentID: string): Promise<boolean> {
    return false
  }

  async getAll (): Promise<Environment[]> {
    // Get all the environments
    return new Array<Environment>()
  }

  async get (environmentID: string): Promise<Environment> {
    return new Environment()
  }
}
