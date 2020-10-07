import * as Msal from 'msal'
import {
  Environment,
  Resource,
  Tag
} from '../types/index'

export class EnvironmentImpl implements Environment {
  id: string = ''
  name: string = ''
  description: string = ''
  createdOn: Date = new Date()
  destroyBy: Date = new Date()
  resources: string[] = new Array<string>()
  tags: Tag[] = new Array<Tag>()
  version: string = ''
  status(): string {
    throw new Error('Method not implemented.')
  }
  create(templateID: string): string[] {
    throw new Error('Method not implemented.')
  }
  destroy(resourceID: string): boolean {
    throw new Error('Method not implemented.')
  }

}

export class EnvironmentService {
  private account: Msal.Account

  constructor(account: Msal.Account) {
    this.account = account
  }

  // Create a new environment and return the ID of the environment
  async create () : Promise<string> {
    return ''
  }

  async delete (environmentID: string) : Promise<boolean> {
    return false
  }

  async getAll () : Promise<Environment[]> {
    // Get all the environments
    return new Array<Environment>()
  }

  async get (environmentID: string) : Promise<Environment> {
    return new EnvironmentImpl()
  }
}
