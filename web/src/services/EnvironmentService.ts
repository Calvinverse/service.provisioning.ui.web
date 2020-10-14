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
    return new Array<Environment>(
      {
        id: 'a',
        name: 'Environment A',
        description: 'Description for Environment A',
        createdOn: new Date(),
        destroyBy: new Date(),
        resources: new Array<string>(),
        tags: new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'production'
          } as TagDefinition
        ),
        version: '1.0.0'
      } as Environment,
      {
        id: 'b',
        name: 'Environment B',
        description: 'Description for Environment B',
        createdOn: new Date(),
        destroyBy: new Date(),
        resources: new Array<string>(),
        tags: new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'test'
          } as TagDefinition
        ),
        version: '3.2.5'
      } as Environment,
      {
        id: 'c',
        name: 'Environment C',
        description: 'Description for Environment C',
        createdOn: new Date(),
        destroyBy: new Date(),
        resources: new Array<string>(),
        tags: new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'dev'
          } as TagDefinition
        ),
        version: '6.7.524'
      } as Environment
    )
  }

  async get (environmentID: string): Promise<Environment> {
    return new Environment()
  }
}
