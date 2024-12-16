import type { FormError } from './Form'

export enum Role {
  USER = 'user',
  ADMIN = 'admin',
  MODERATOR = 'moderator',
}

export interface NewRole {
  login: string
  new_role: Role
}

export interface NewRoleErrors {
  login?: FormError[]
  new_role?: FormError[]
}
