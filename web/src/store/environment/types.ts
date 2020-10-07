export interface EnvironmentState {
  error: boolean

  id: string
  name: string
  description: string

  createdOn: Date
  destroyBy: Date

  status: string

  resources: string[] // ID of resources

  tags: string[] // ID of tags

  version: string

  // status
}
