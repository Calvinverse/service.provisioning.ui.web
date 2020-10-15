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
  statusUrl = ''
  tags: TagDefinition[] = new Array<TagDefinition>()
  version = ''

  constructor (
    id = '',
    name = '',
    description = '',
    createdOn: Date = new Date(),
    destroyBy: Date = new Date(),
    resources: string[] = new Array<string>(),
    statusUrl = '',
    tags: TagDefinition[] = new Array<TagDefinition>(),
    version = ''
  ) {
    this.id = id
    this.name = name
    this.description = description
    this.createdOn = createdOn
    this.destroyBy = destroyBy
    this.resources = resources
    this.tags = tags
    this.statusUrl = statusUrl
    this.version = version
  }

  status (): string {
    return this.statusUrl === 'https://example.com/status/ok' ? 'ok' : 'failure'
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
      new Environment(
        'a',
        'Environment A',
        'Description for Environment A',
        new Date(),
        new Date(),
        new Array<string>(),
        'https://example.com/status/ok',
        new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'source',
            value: 'abcdef123454566'
          } as TagDefinition,
          {
            name: 'version',
            value: '1.0.0'
          } as TagDefinition,
          {
            name: 'location',
            value: 'Australia East'
          } as TagDefinition,
          {
            name: 'cloud',
            value: 'Azure'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'production'
          } as TagDefinition
        ),
        '1.0.0'),
      new Environment(
        'b',
        'Environment B',
        'Description for Environment B',
        new Date(),
        new Date(),
        new Array<string>(),
        'https://example.com/status/failure',
        new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'test'
          } as TagDefinition
        ),
        '3.2.5'),
      new Environment(
        'c',
        'Environment C',
        'Description for Environment C',
        new Date(),
        new Date(),
        new Array<string>(),
        'https://example.com/status/ok',
        new Array<TagDefinition>(
          {
            name: 'createdBy',
            value: 'terraform'
          } as TagDefinition,
          {
            name: 'environment',
            value: 'dev'
          } as TagDefinition
        ),
        '6.7.524')
    )
  }

  async get (environmentID: string): Promise<Environment> {
    return new Environment()
  }
}
