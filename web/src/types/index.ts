
export interface Environment {
  id: string
  name: string
  description: string

  createdOn: Date
  destroyBy: Date

  resources: string[]

  tags: Tag[]

  version: string

  status (): string
  create(templateID: string): string[]
  destroy(resourceID: string): boolean
}

export interface Resource {
  id: string

  name: string

  // List of ID numbers of resources that the current resource depends on
  dependsOn: string[]

  // ID of the template that was used to create the current resource
  templateId: string

  tags: Tag[]

  status (): string
}

export interface Tag {
  name: string
  value: string
}