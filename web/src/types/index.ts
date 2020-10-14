
export interface EnvironmentDefinition {
  id: string;
  name: string;
  description: string;

  createdOn: Date;
  destroyBy: Date;

  resources: string[];

  tags: TagDefinition[];

  version: string;

  status (): string;
  create(templateID: string): string[];
  destroy(resourceID: string): boolean;
}

export interface ResourceDefinition {
  id: string;

  name: string;

  // List of ID numbers of resources that the current resource depends on
  dependsOn: string[];

  // ID of the template that was used to create the current resource
  templateId: string;

  tags: TagDefinition[];

  status (): string;
}

export interface TagDefinition {
  name: string;
  value: string;
}

export class Constants {
  static get authenticationChannel (): string {
    return 'authentication'
  }

  static get userLoginTopic (): string {
    return 'authentication-user-login'
  }

  static get userLogoutTopic (): string {
    return 'authentication-user-logout'
  }
}
